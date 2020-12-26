package accounts

import (
	"errors"
	"fmt"
)

// Account create struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Cant'w withdraw")

// NewAccount create Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (theAccount *Account) Deposit(amount int) {
	theAccount.balance += amount
}

// Balance of your account
func (theAccount Account) Balance() int {
	return theAccount.balance
}

// Withdraw x amount from your account
func (theAccount *Account) Withdraw(amount int) error {
	if theAccount.balance < amount {
		return errNoMoney
	}
	theAccount.balance -= amount
	return nil
}

// Owner of account
func (theAccount Account) Owner() string {
	return theAccount.owner
}

func (theAccount Account) String() string {
	return fmt.Sprint(theAccount.owner, "'s account.\nHas: ", theAccount.Balance())
}
