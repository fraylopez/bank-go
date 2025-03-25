package money

func USD(amount float64) Money {
	return Money{
		Amount:   amount,
		Currency: "USD",
	}
}

func EUR(amount float64) Money {
	return Money{
		Amount:   amount,
		Currency: "EUR",
	}
}
