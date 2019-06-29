package pointer

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("Got %s Want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Errorf("Wanted an error but didn't one")
		}

		if got != want {
			t.Errorf("Got '%s', want '%s'", got, want)
		}
	}

	t.Run("Deposit ", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)
		wallet.Withdraw(2)

		want := Bitcoin(8)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw Insufficient amount", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(20)
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
