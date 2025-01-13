package domain

type Money struct {
	Amount   float64
	Currency string
}

type SupportedCurrencies struct {
	EUR string
	USD string
}

func USD(amount float64) Money {
	return Money{
		Amount:   amount,
		Currency: "USD",
	}
}

func NewMoney(currency string) Money {
	return Money{
		Amount:   0,
		Currency: currency,
	}
}

func (m *Money) Add(adder Money) (Money, error) {
	if m.Currency != adder.Currency {
		return Money{}, &CurrencyMismatchError{}
	}
	return Money{
		Amount:   m.Amount + adder.Amount,
		Currency: m.Currency,
	}, nil
}

func (m *Money) Subtract(amount float64, currency string) (Money, error) {
	if m.Currency != currency {
		// convert amount to m.Currency
		return Money{}, &CurrencyMismatchError{}
	}
	return Money{
		Amount:   m.Amount - amount,
		Currency: m.Currency,
	}, nil
}

func (m *Money) Equals(other Money) bool {
	return m.Amount == other.Amount && m.Currency == other.Currency
}
