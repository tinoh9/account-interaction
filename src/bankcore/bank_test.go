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
		t.Error("Cannot create an Account object")
	}
}

func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name: "Adam",
			Address: "Westminster, California",
			Phone: "(123) 456 789",
		},
		Number: 101,
		Balance: 0,
	}
	account.Deposit(10)

	if account.Balance != 10 {
		t.Error("Balance is not being updated after the deposit!")
	}
}

func TestInvalidDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name: "Adam",
			Address: "Westminster, California",
			Phone: "(123) 456 789",
		},
		Number: 1001,
		Balance: 0,
	}
	if err := account.Deposit(-10); err == nil {
		t.Error("Only positive number is allowed to deposit!")
	}
}

func TestWithdraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name: "Adam",
			Address: "Westminster, California",
			Phone: "(123) 456 789",
		},
		Number: 1001,
		Balance: 0,
	}
	account.Deposit(10)
	account.Withdraw(10)

	if account.Balance != 0 {
		t.Error("Balance is not being updated after withdraw!")
	}
}

func TestStatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name: "Adam",
			Address: "Westminster, California",
			Phone: "(123) 456 789",
		},
		Number: 1001,
		Balance: 0,
	}
	account.Deposit(500)
	statement := account.Statement()
	if statement != "1001 - Adam - 500" {
		t.Error("Account statement does not have the proper format!")
	}
}