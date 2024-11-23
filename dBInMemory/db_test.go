package dBInMemory

import (
	"testing"
)

// TestSetAndGet verifies that values can be set and retrieved correctly.
func TestSetAndGet(t *testing.T) {
	store := NewDB()

	// Set values
	store.Set(1, "value1")
	store.Set(2, "value2")

	// Verify values
	if store.data[1] != "value1" {
		t.Fatalf("Expected value1 for key 1, got %v", store.data[1])
	}
	if store.data[2] != "value2" {
		t.Fatalf("Expected value2 for key 2, got %v", store.data[2])
	}
}

// TestDeleteByID verifies that values can be deleted by their ID.
func TestDeleteByID(t *testing.T) {
	store := NewDB()

	// Set and delete a value
	store.Set(1, "value1")
	store.DeleteByID(1)

	// Verify deletion
	if _, exists := store.data[1]; exists {
		t.Fatalf("Expected key 1 to be deleted, but it still exists")
	}
}

// TestBeginTransaction verifies that a transaction can begin without errors.
func TestBeginTransaction(t *testing.T) {
	store := NewDB()

	// Begin a transaction
	if err := store.Begin(); err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}
}

// TestCommitTransaction verifies that changes made in a transaction are committed.
func TestCommitTransaction(t *testing.T) {
	store := NewDB()

	// Set initial values and begin a transaction
	store.Set(1, "value1")
	store.Begin()

	// Make changes in the transaction
	store.Set(1, "newValue1")
	store.Set(2, "value2")

	// Commit the transaction
	if err := store.Commit(); err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	// Verify final state after commit
	if store.data[1] != "newValue1" || store.data[2] != "value2" {
		t.Fatalf("Unexpected state after commit: %v", store.data)
	}
}

// TestRollbackTransaction verifies that changes made in a transaction are rolled back.
func TestRollbackTransaction(t *testing.T) {
	store := NewDB()

	// Set initial values and begin a transaction
	store.Set(1, "value1")
	store.Begin()

	// Make changes in the transaction
	store.Set(2, "value2")
	store.DeleteByID(1)

	// Rollback the transaction
	store.Rollback()

	// Verify final state after rollback
	if _, exists := store.data[2]; exists {
		t.Fatalf("Rollback failed, key 2 should not exist: %v", store.data)
	}
	if store.data[1] != "value1" {
		t.Fatalf("Rollback failed, key 1 should still have value 'value1'")
	}
}

// TestMultipleTransactions verifies that multiple transactions work as expected.
func TestMultipleTransactions(t *testing.T) {
	store := NewDB()

	// Set initial values
	store.Set(1, "value1")

	// First transaction: Commit changes
	store.Begin()
	store.Set(2, "value2")
	if err := store.Commit(); err != nil {
		t.Fatalf("Failed to commit first transaction: %v", err)
	}

	if store.data[2] != "value2" {
		t.Fatalf("First transaction commit failed: %v", store.data)
	}

	// Second transaction: Rollback changes
	store.Begin()
	store.Set(3, "value3")
	store.Rollback()

	if _, exists := store.data[3]; exists {
		t.Fatalf("Second transaction rollback failed, key 3 should not exist: %v", store.data)
	}
}
