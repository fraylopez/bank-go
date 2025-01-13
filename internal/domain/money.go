package domain

type Money struct {
	Amount   float64
	Currency Currencies
}
type Currencies string

const (
	EUROS    Currencies = "EUR"
	USDollar            = "USD"
)

func Currency(currency string) Currencies {
	switch currency {
	case "EUR":
		return EUROS
	case "USD":
		return USDollar
	default:
		panic("unknown currency")
	}
}

func (c Currencies) String() string {
	switch c {
	case EUROS:
		return "EUR"
	case USDollar:
		return "USD"
	default:
		panic("unknown currency")
	}
}

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

func (m *Money) IsLessThan(other Money) bool {
	return m.Amount < other.Amount
}
