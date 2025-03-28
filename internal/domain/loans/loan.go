package loans

import "github.com/fraylopez/bank-go/internal/domain/account"

type LoanService interface {
	CanAssignLoan(account *account.Account) (bool, error)
}

type NoLoansService struct{}

func NewNoLoansService() LoanService {
	return &NoLoansService{}
}

func (*NoLoansService) CanAssignLoan(_ *account.Account) (bool, error) {
	return false, nil
}
