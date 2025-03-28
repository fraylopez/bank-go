package loans

import (
	"github.com/fraylopez/bank-go/internal/domain/money"
)

type LoanService interface {
	CanAssignLoan(accountBalance money.Money, requested money.Money) (bool, error)
}

type NoLoansService struct{}

func NewNoLoansService() LoanService {
	return &NoLoansService{}
}

func (*NoLoansService) CanAssignLoan(_ money.Money, _ money.Money) (bool, error) {
	return false, nil
}
