package db

import (
	"context"
	"testing"
	"time"

	"github.com/PfMartin/secure-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) (Transfer, int64, int64) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID: toAccount.ID,
		Amount: util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer, fromAccount.ID, toAccount.ID
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	createdTransfer, _, _ := createRandomTransfer(t)
	gotTransfer, err := testQueries.GetTransfer(context.Background(), createdTransfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotTransfer)

	require.Equal(t, createdTransfer.FromAccountID, gotTransfer.FromAccountID)
	require.Equal(t, createdTransfer.ToAccountID, gotTransfer.ToAccountID)
	require.Equal(t, createdTransfer.Amount, gotTransfer.Amount)
	require.WithinDuration(t, createdTransfer.CreatedAt.Time, gotTransfer.CreatedAt.Time, time.Second)
}

func TestListTransfers(t *testing.T) {
	_, fromAccountId, toAccountId := createRandomTransfer(t)

	args := ListTransfersParams{
		FromAccountID: fromAccountId,
		ToAccountID: toAccountId,
		Limit: 1,
		Offset: 0,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, transfers, 1)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}