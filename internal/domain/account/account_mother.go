package account

func BuildAccount() *Account {
	holder := "John Doe"
	return NewAccount(holder, "USD")
}

func BuildEURAccount() *Account {
	holder := "John Doe"
	return NewAccount(holder, "EUR")
}

func BuildUSDAccount() *Account {
	holder := "John Doe"
	return NewAccount(holder, "USD")
}
