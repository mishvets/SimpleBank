package db

import (
	"context"
	"testing"

	"github.com/mishvets/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, accIdFrom, accIdTo int64) Transfer {
	arg := CreateTransferParams{
		FromAccountID: accIdFrom,
		ToAccountID:   accIdTo,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t, createRandomAccount(t).ID, createRandomAccount(t).ID)
}

func TestGetTransfer(t *testing.T) {
	transferExpected := createRandomTransfer(t, createRandomAccount(t).ID, createRandomAccount(t).ID)
	transferGetter, err := testQueries.GetTransfer(context.Background(), transferExpected.ID)
	require.NoError(t, err)
	require.Equal(t, transferExpected, transferGetter) // Note: can be difference in the field "CreatedAt"
}

func TestListTransfers(t *testing.T) {
	accFrom := createRandomAccount(t)
	accTo := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		createRandomTransfer(t, accFrom.ID, accTo.ID)
		createRandomTransfer(t, accTo.ID, accFrom.ID)
	}
	arg := ListTransfersParams{
		FromAccountID: accFrom.ID,
		ToAccountID: accFrom.ID,
		Limit:     5,
		Offset:    5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == accFrom.ID || transfer.ToAccountID == accFrom.ID)
	}
}
