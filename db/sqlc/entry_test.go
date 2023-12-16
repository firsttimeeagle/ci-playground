package db

import (
	"context"
	"testing"
	"time"

	"github.com/brizaldi/go-simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, accountID *int64) Entry {
	var account Account
	var arg CreateEntryParams

	if accountID == nil {
		account = createRandomAccount(t)
		arg = CreateEntryParams{
			AccountID: account.ID,
			Amount:    util.RandomMoney(),
		}
	} else {
		arg = CreateEntryParams{
			AccountID: *accountID,
			Amount:    util.RandomMoney(),
		}
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
	createRandomEntry(t, nil)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t, nil)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntrys(t *testing.T) {
	account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomEntry(t, &account.ID)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}