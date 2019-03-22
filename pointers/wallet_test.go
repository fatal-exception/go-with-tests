package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("Was expecting an error from Withdraw")
		}
		if got != want {
			t.Errorf("Error string got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
		if err != nil {
			t.Errorf("Was not expecting an error from withdraw")
		}

	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingFunds := Bitcoin(20)
		wallet := Wallet{balance: startingFunds}
		err := wallet.Withdraw(Bitcoin(30))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingFunds)
	})
}
