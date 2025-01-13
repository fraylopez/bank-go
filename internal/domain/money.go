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

func MoneyFrom(amount float64, currency string) Money {
	return Money{
		Amount:   amount,
		Currency: currency,
	}
}

func (m *Money) Add(addend Money) (Money, error) {
	if m.Currency != addend.Currency {
		return Money{}, &CurrencyMismatchError{}
	}
	return Money{

		Amount:   m.Amount + addend.Amount,
		Currency: m.Currency,
	}, nil
}

func (m *Money) Subtract(subtrahend Money) (Money, error) {
	if m.Currency != subtrahend.Currency {
		// convert amount to m.Currency
		return Money{}, &CurrencyMismatchError{}
	}
	return Money{
		Amount:   m.Amount - subtrahend.Amount,
		Currency: m.Currency,
	}, nil
}

func (m *Money) Equals(other Money) bool {
	return m.Amount == other.Amount && m.Currency == other.Currency
}
