package domain

import (
	"testing"
)

func TestMoney(t *testing.T) {
	t.Run("Prevents unknown Currencies", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic, got nil")
			}
		}()
		_ = NewMoney("XYZ")
	})

	t.Run("Equals", func(t *testing.T) {
		money := USD(10)
		if !money.Equals(USD(10)) {
			t.Errorf("Expected 10, got %v", money)
		}
	})

	t.Run("Add", func(t *testing.T) {
		money := USD(10)
		addend := USD(5)
		result, _ := money.Add(addend)
		if !result.Equals(USD(15)) {
			t.Errorf("Expected 15, got %v", result)
		}
	})

	t.Run("Add with currency mismatch", func(t *testing.T) {
		money := USD(10)
		addend := EUR(5)
		_, err := money.Add(addend)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("Subtract", func(t *testing.T) {
		money := USD(10)
		subtrahend := USD(5)
		result, _ := money.Subtract(subtrahend)
		if !result.Equals(USD(5)) {
			t.Errorf("Expected 5, got %v", result)
		}
	})

	t.Run("Subtract with currency mismatch", func(t *testing.T) {
		money := USD(10)
		subtrahend := EUR(5)
		_, err := money.Subtract(subtrahend)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
