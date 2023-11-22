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
	createdAccount := createRandomAccount(t)
	gotAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotAccount)

	require.Equal(t, createdAccount.ID, gotAccount.ID)
	require.Equal(t, createdAccount.Owner, gotAccount.Owner)
	require.Equal(t, createdAccount.Balance, gotAccount.Balance)
	require.Equal(t, createdAccount.Currency, gotAccount.Currency)
	require.WithinDuration(t, createdAccount.CreatedAt.Time, gotAccount.CreatedAt.Time, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)

	arg := UpdateAccountParams {
		ID: createdAccount.ID,
		Balance: util.RandomMoney(),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, createdAccount.ID, updatedAccount.ID)
	require.Equal(t, createdAccount.Owner, updatedAccount.Owner)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, createdAccount.Currency, updatedAccount.Currency)
	require.WithinDuration(t, createdAccount.CreatedAt.Time, updatedAccount.CreatedAt.Time, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)
	deletedAccount, err := testQueries.DeleteAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, deletedAccount)

	require.Equal(t, createdAccount.ID, deletedAccount.ID)
	require.Equal(t, createdAccount.Owner, deletedAccount.Owner)
	require.Equal(t, createdAccount.Balance, deletedAccount.Balance)
	require.Equal(t, createdAccount.Currency, deletedAccount.Currency)
	require.WithinDuration(t, createdAccount.CreatedAt.Time, deletedAccount.CreatedAt.Time, time.Second)

	got_account, err := testQueries.GetAccount(context.Background(), deletedAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, got_account)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams {
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}