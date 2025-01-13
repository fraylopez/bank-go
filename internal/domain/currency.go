package domain

type Currencies string

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
