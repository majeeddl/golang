package pointers

import "testing"

func TestWallet(t *testing.T) {

	assertError := func(r testing.TB, got error, want error) {

		t.Helper()

		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		// wallet.Deposit(10)

		// got := wallet.Balance()
		// want := 10

		// if got != want {
		// 	t.Errorf("got %d want %d", got, want)
		// }

		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

		// got := wallet.Balance()

		// want := Bitcoin(10)

		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {

		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)

		// if err == nil {
		// 	t.Error("wanted an error but didn't get one")
		// }

	})

}
