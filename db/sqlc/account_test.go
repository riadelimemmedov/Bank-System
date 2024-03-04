package db

import (
	"context"
	"fmt"
	"simplebank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// ?createRandomAccount
func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Name:     util.RandomOwner(),
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testStore.CreateAccount(context.Background(), arg)
	fmt.Println("Created Account is ", account)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Name, account.Name)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

// !TestCreateAccount
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

// !TestGetAccount
func TestGetAccount(t *testing.T) {
	created_account := createRandomAccount(t)
	selected_account, err := testStore.GetAccount(context.Background(), created_account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, selected_account)

	require.Equal(t, created_account.ID, selected_account.ID)
	require.Equal(t, created_account.Owner, selected_account.Owner)
	require.Equal(t, created_account.Balance, selected_account.Balance)
	require.Equal(t, created_account.Currency, selected_account.Currency)
	require.WithinDuration(t, created_account.CreatedAt, selected_account.CreatedAt, time.Second)
}

// ! TestUpdateAccount
func TestUpdateAccount(t *testing.T) {
	created_account := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      created_account.ID,
		Name:    util.RandomOwner(),
		Balance: util.RandomMoney(),
	}

	updated_account, err := testStore.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updated_account)
	require.Equal(t, created_account.ID, updated_account.ID)
	require.NotEqual(t, created_account.Name, updated_account.Name)
	require.Equal(t, arg.Name, updated_account.Name)
	require.Equal(t, created_account.Owner, updated_account.Owner)
	require.NotEqual(t, created_account.Balance, updated_account.Balance)
	require.Equal(t, arg.Balance, updated_account.Balance)
	require.Equal(t, created_account.Currency, updated_account.Currency)
	require.WithinDuration(t, created_account.CreatedAt, updated_account.CreatedAt, time.Second)
}

// !TestDeleteAccount
func TestDeleteAccount(t *testing.T) {
	created_account := createRandomAccount(t)
	err := testStore.DeleteAccount(context.Background(), created_account.ID)
	require.NoError(t, err)
	require.Nil(t, err)

	account, err := testStore.GetAccount(context.Background(), created_account.ID)

	require.Error(t, err)
	require.Empty(t, account)
}

// !TestListAccounts
func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testStore.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

}
