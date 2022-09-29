package pointersAndErrors

type Wallet struct {
	balance int
}

// take a pointer to that wallet so that we can change the original values within it
//(w *Wallet) = struct pointers
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
