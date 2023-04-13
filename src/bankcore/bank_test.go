package bank

import "testing"

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name: "Adam",
			Address: "Westminster, California",
			Phone: "(123) 456 789",
		},
		Number: 101,
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

func TestInvalidNegativeDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name: "Adam",
			Address: "Westminster, California",
			Phone: "(123) 456 789",
		},
		Number: 101,
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
		Number: 101,
		Balance: 0,
	}
	account.Deposit(10)
	account.Withdraw(10)

	if account.Balance != 0 {
		t.Error("Balance is not being updated after withdraw!")
	}
}

func TestInvalidNegativeWithdraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name: "Adam",
			Address: "Westminster, California",
			Phone: "(123) 456 789",
		},
		Number: 101,
		Balance: 0,
	}
	if err := account.Withdraw(-10); err == nil {
		t.Error("Only positive number is allowed to withdraw!")
	}
}

func TestInvalidExceedWithdraw(t *testing.T) {
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
	account.Withdraw(20)

	if account.Balance < 0 {
		t.Error("Withdraw amount should be smaller than the deposit amount!")
	}
}

func TestStatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name: "Adam",
			Address: "Westminster, California",
			Phone: "(123) 456 789",
		},
		Number: 101,
		Balance: 0,
	}
	account.Deposit(500)
	statement := account.Statement()
	if statement != "101 - Adam - 500" {
		t.Error("Account statement does not have the proper format!")
	}
}

func TestTransfer(t *testing.T) {
	account1 := Account{
		Customer: Customer{
			Name: "Adam",
			Address: "Westminster, California",
			Phone: "(123) 456 789",
		},
		Number: 101,
		Balance: 0,
	}

	account2 := Account{
		Customer: Customer{
			Name: "Eve",
			Address: "Los Angeles, California",
			Phone: "(000) 000 000",
		},
		Number: 102,
		Balance: 0,
	}
	account1.Deposit(500)
	err := account1.Transfer(200, &account2)

	if account1.Balance != 300 && account2.Balance != 200 {
		t.Error("Transfer between Account 1 and Account 2 is not working!", err)
	}
}

func TestInvalidNegativeTransfer(t *testing.T) {
	account1 := Account{
		Customer: Customer{
			Name: "Adam",
			Address: "Westminster, California",
			Phone: "(123) 456 789",
		},
		Number: 101,
		Balance: 0,
	}

	account2 := Account{
		Customer: Customer{
			Name: "Eve",
			Address: "Los Angeles, California",
			Phone: "(000) 000 000",
		},
		Number: 102,
		Balance: 0,
	}
	account1.Deposit(500)
	if err := account1.Transfer(-200, &account2); err == nil {
		t.Error("Transfer amount should be positive number!", err)
	}
}

func TestInvalidExceedTransfer(t *testing.T) {
	account1 := Account{
		Customer: Customer{
			Name: "Adam",
			Address: "Westminster, California",
			Phone: "(123) 456 789",
		},
		Number: 101,
		Balance: 0,
	}

	account2 := Account{
		Customer: Customer{
			Name: "Eve",
			Address: "Los Angeles, California",
			Phone: "(000) 000 000",
		},
		Number: 102,
		Balance: 0,
	}
	account1.Deposit(500)
	if err := account1.Transfer(600, &account2); err == nil {
		t.Error("Transfer amount of account 1 should be smaller than the deposit amount of account 1!", err)
	}
}
