package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransfer(t *testing.T) {
	sender := createRandomAccount(t)
	receiver := createRandomAccount(t)

	fmt.Println(">> before:", sender.Balance, receiver.Balance)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := testStore.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: sender.ID,
				ToAccountID:   receiver.ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}
	// check results
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		//? check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, sender.ID, transfer.FromAccountID)
		require.Equal(t, receiver.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, amount)
		require.NotZero(t, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = testStore.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		//? check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, sender.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = testStore.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, receiver.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = testStore.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// ?check accounts
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, sender.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, receiver.ID, toAccount.ID)

		// ? check accounts balance
		fmt.Println(">>tx ", fromAccount.Balance, toAccount.Balance)
		diff1 := sender.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - receiver.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// ?check the final updated balance
	result_balance_sender, err := testStore.GetAccount(context.Background(), sender.ID)
	require.NoError(t, err)

	result_balance_receiver, err := testStore.GetAccount(context.Background(), receiver.ID)
	require.NoError(t, err)

	require.Equal(t, sender.Balance-int64(n)*amount, result_balance_sender.Balance)
	require.Equal(t, receiver.Balance+int64(n)*amount, result_balance_receiver.Balance)

	fmt.Println(">> before:", result_balance_sender.Balance, result_balance_receiver.Balance)

}
