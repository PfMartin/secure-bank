package db

import (
	"context"
	"testing"
	"time"

	"github.com/PfMartin/secure-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	created_account := createRandomAccount(t)
	got_account, err := testQueries.GetAccount(context.Background(), created_account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, got_account)

	require.Equal(t, created_account.ID, got_account.ID)
	require.Equal(t, created_account.Owner, got_account.Owner)
	require.Equal(t, created_account.Balance, got_account.Balance)
	require.Equal(t, created_account.Currency, got_account.Currency)
	require.WithinDuration(t, created_account.CreatedAt.Time, got_account.CreatedAt.Time, time.Second)
}