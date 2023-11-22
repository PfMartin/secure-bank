package db

import (
	"context"
	"testing"
	"time"

	"github.com/PfMartin/secure-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	created_account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: created_account.ID,
		Amount: util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	createdEntry := createRandomEntry(t)
	gotEntry, err := testQueries.GetEntry(context.Background(), createdEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotEntry)

	require.Equal(t, createdEntry.ID, gotEntry.ID)
	require.Equal(t, createdEntry.AccountID, gotEntry.AccountID)
	require.Equal(t, createdEntry.Amount, gotEntry.Amount)
	require.WithinDuration(t, createdEntry.CreatedAt.Time, gotEntry.CreatedAt.Time, time.Second)
}

func TestListEntries(t *testing.T) {
	createdEntry := createRandomEntry(t)

	arg := ListEntriesParams {
		Limit: 1,
		Offset: 0,
		AccountID: createdEntry.AccountID,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)

	require.Len(t, entries, 1)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}