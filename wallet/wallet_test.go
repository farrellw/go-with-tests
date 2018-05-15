package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	compareBitcoin := func(t *testing.T, w Wallet, want Bitcoin){
		t.Helper()
		got := w.Balance()

		if got != want {
			t.Errorf("Got '%s', wanted '%s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T){
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		compareBitcoin(t, wallet, want)
	})
	t.Run("Withdraw", func(t *testing.T){
		wallet := Wallet{balance: Bitcoin(30)}

		wallet.Withdraw(Bitcoin(18))
		want := Bitcoin(12)

		compareBitcoin(t, wallet, want)
	})
}
