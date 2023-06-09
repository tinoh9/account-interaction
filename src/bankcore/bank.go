package bank

import (
	"errors"
	"fmt"
)

// Customer struct
type Customer struct {
	Name 	string
	Address string
	Phone 	string
}

// Account struct
type Account struct {
	Customer
	Number 	int32
	Balance float64
}

// Bank interface
type Bank interface {
	Statement() string
}

// Deposit method
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("You cannot deposit zero or negative amount!")
	}

	a.Balance += amount
	return nil
}

// Withdraw method
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("You cannot withdraw zero or negative amount!")
	}

	if amount > a.Balance {
		return errors.New("You cannot withdraw the amount which is larger than the account balance!")
	}

	a.Balance -= amount
	return nil
}

// Statement method
func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

// New Statement method
func Statement(b Bank) string {
	return b.Statement()
}

// Transfer method
func (a *Account) Transfer(amount float64, des *Account) error {
	if amount <= 0 {
		return errors.New("You cannot transfer zero or negative amount!")
	}

	if amount > a.Balance{
		return errors.New("You cannot withdraw the amount which is larger than the account balance!")
	}

	a.Withdraw(amount)
	des.Deposit(amount)
	return nil
}