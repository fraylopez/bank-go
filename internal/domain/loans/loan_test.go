package loans_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fraylopez/bank-go/internal/domain/account"
	"github.com/fraylopez/bank-go/internal/domain/loans"
)

func TestNoLoansService_CanAssignLoan(t *testing.T) {
	t.Run("should not allow loans", func(t *testing.T) {
		service := loans.NewNoLoansService()

		allowed, _ := service.CanAssignLoan(account.BuildEURAccount())

		assert.False(t, allowed)
	})
}
