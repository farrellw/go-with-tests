package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	compareBalance := func(t *testing.T, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()

		if got != want {
			t.Errorf("Got '%s', wanted '%s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		if got == nil {
			t.Fatal("Wanted an error, but did not get one")
		}

		if got != want {
			t.Errorf("Got '%s', wanted '%s'", got, want)
		}
	}

	assertNoError := func(t *testing.T, err error){
		if err != nil{
			t.Errorf("Wanted no error, got '%s", err)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		compareBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		t.Run("Sufficient Funds", func(t *testing.T) {
			wallet := Wallet{balance: Bitcoin(30)}

			err := wallet.Withdraw(Bitcoin(18))
			want := Bitcoin(12)

			compareBalance(t, wallet, want)
			assertNoError(t, err)
		})

		t.Run("Insufficient Funds", func(t *testing.T){
			startingBalance := Bitcoin(5)
			withdrawBalance := Bitcoin(8)
			wallet := Wallet{balance: startingBalance}

			err := wallet.Withdraw(withdrawBalance)
			
			compareBalance(t, wallet, startingBalance)

			assertError(t, err, ErrInsufficientFunds)
		})
	})
}
