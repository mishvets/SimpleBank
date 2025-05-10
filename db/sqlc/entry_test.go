package db

import (
	"context"
	"testing"

	"github.com/mishvets/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, accId int64) Entry {
	arg := CreateEntryParams{
		AccountID: accId,
		Amount:    util.RandomMoney(),
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
	createRandomEntry(t, createRandomAccount(t).ID)
}

func TestGetEntry(t *testing.T) {
	entryExpected := createRandomEntry(t, createRandomAccount(t).ID)
	entryGetter, err := testQueries.GetEntry(context.Background(), entryExpected.ID)
	require.NoError(t, err)
	require.Equal(t, entryExpected, entryGetter) // Note: can be difference in the field "CreatedAt"
}

func TestListEntries(t *testing.T) {
	acc := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, acc.ID)
	}
	arg := ListEntriesParams{
		AccountID: acc.ID,
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)

	require.Len(t, entries, 5)
	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
