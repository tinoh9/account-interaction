package bank

import "testing"

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name: "Adam",
			Address: "Westminster, California",
			Phone: "(123) 456 789",
		},
		Number: 1001,
		Balance: 0,
	}
	if account.Name == "" {
		t.Error ("Cannot create anAccount object")
	}
}