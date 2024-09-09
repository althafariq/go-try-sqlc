package db

// func TestTransferTx(t *testing.T) {
// 	store := NewStore(testDB)

// 	account1 := createRandomAccount(t)
// 	account2 := createRandomAccount(t)

// 	// run a concurrent transfer transaction
// 	n := 5
// 	amount := int64(10)

// 	for i := 0; i < n; i++ {
// 		go func() {

// 	}
// }

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecTx(t *testing.T) {
	store := NewStore(testDB)

	// Test successful transaction
	err := store.execTx(context.Background(), func(q *Queries) error {
		// Perform some queries here
		_, err := q.CreateAccount(context.Background(), CreateAccountParams{
			Owner:    "test_owner",
			Balance:  100,
			Currency: "USD",
		})
		return err
	})
	require.NoError(t, err)

	// Test transaction rollback
	err = store.execTx(context.Background(), func(q *Queries) error {
		// Perform some queries here
		_, err := q.CreateAccount(context.Background(), CreateAccountParams{
			Owner:    "test_owner",
			Balance:  100,
			Currency: "USD",
		})
		if err != nil {
			return err
		}
		// Force an error to trigger rollback
		return fmt.Errorf("forced error")
	})
	require.Error(t, err)
}
