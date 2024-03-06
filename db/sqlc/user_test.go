package db

import (
	"context"
	"simplebank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// ?createRandomAccount
func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testStore.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

// !TestCreateAccount
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

// !TestGetAccount
func TestGetUser(t *testing.T) {
	created_user := createRandomUser(t)
	selected_user, err := testStore.GetUser(context.Background(), created_user.Username)

	require.NoError(t, err)
	require.NotEmpty(t, selected_user)

	require.Equal(t, created_user.Username, selected_user.Username)
	require.Equal(t, created_user.HashedPassword, selected_user.HashedPassword)
	require.Equal(t, created_user.FullName, selected_user.FullName)
	require.Equal(t, created_user.Email, selected_user.Email)
	require.WithinDuration(t, created_user.PasswordChangedAt, selected_user.PasswordChangedAt, time.Second)
	require.WithinDuration(t, created_user.CreatedAt, selected_user.CreatedAt, time.Second)
}
