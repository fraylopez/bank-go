package money

import "github.com/fraylopez/bank-go/internal/domain"

type Money struct {
	Amount   float64
	Currency Currencies
}

const (
	EUROS    Currencies = "EUR"
	USDollar            = "USD"
)

func NewMoney(currency string) Money {
	return Money{
		Amount:   0,
		Currency: Currency(currency),
	}
}

func MoneyFrom(amount float64, currency string) Money {
	return Money{
		Amount:   amount,
		Currency: Currency(currency),
	}
}

func (m *Money) Add(addend Money) (Money, error) {
	if m.Currency != addend.Currency {
		return Money{}, &domain.CurrencyMismatchError{}
	}
	return Money{

		Amount:   m.Amount + addend.Amount,
		Currency: m.Currency,
	}, nil
}

func (m *Money) Subtract(subtrahend Money) (Money, error) {
	if m.Currency != subtrahend.Currency {
		// convert amount to m.Currency
		return Money{}, &domain.CurrencyMismatchError{}
	}
	return Money{
		Amount:   m.Amount - subtrahend.Amount,
		Currency: m.Currency,
	}, nil
}

func (m *Money) Equals(other Money) bool {
	return m.Amount == other.Amount && m.Currency == other.Currency
}

func (m *Money) IsLessThan(other Money) bool {
	return m.Amount < other.Amount
}
