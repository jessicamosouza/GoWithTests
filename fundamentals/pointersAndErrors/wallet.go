package pointersAndErrors

import "fmt"

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

// take a pointer to that wallet so that we can change the original values within it
// (w *Wallet) = struct pointers
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
