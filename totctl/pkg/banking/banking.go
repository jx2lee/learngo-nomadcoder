package banking

import (
	"fmt"

	"github.com/jx2lee/learngo-nomadcoder/totctl/pkg/banking/accounts"
)

func GetBankInfo() {

	account := accounts.NewAccount("jx2lee")
	account.Deposit(1200)
	fmt.Println(account.Balance())
	err := account.Withdraw(500)
	if err != nil {
		fmt.Println(err)
	}

	// String function
	fmt.Println(account)

}
