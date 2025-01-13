package domain

func BuildAccount() *Account {
	holder := "John Doe"
	return NewAccount(holder, "USD")
}
