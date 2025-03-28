package loans_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fraylopez/bank-go/internal/domain/loans"
	"github.com/fraylopez/bank-go/internal/domain/money"
)

func TestNoLoansService_CanAssignLoan(t *testing.T) {
	t.Run("should not allow loans", func(t *testing.T) {
		service := loans.NewNoLoansService()

		allowed, _ := service.CanAssignLoan(money.EUR(10), money.EUR(20))

		assert.False(t, allowed)
	})
}
