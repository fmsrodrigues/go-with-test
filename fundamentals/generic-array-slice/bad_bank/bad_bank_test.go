package bad_bank_test

import (
	a "fundamentals/generic-array-slice"
	bb "fundamentals/generic-array-slice/bad_bank"

	"testing"
)

func TestBadBank(t *testing.T) {
	var (
		riya  = bb.Account{Name: "Riya", Balance: 100}
		chris = bb.Account{Name: "Chris", Balance: 75}
		adil  = bb.Account{Name: "Adil", Balance: 200}

		transactions = []bb.Transaction{
			bb.NewTransaction(chris, riya, 100),
			bb.NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account bb.Account) float64 {
		return bb.NewBalanceFor(account, transactions).Balance
	}

	a.AssertEqual(t, newBalanceFor(riya), 200)
	a.AssertEqual(t, newBalanceFor(chris), 0)
	a.AssertEqual(t, newBalanceFor(adil), 175)
}
