package account

import "sync"

// Define the Account type here.
type Account struct {
	Amount   int64
	IsClosed bool
	mu       *sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{
		Amount:   amount,
		IsClosed: false,
		mu:       &sync.Mutex{},
	}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.IsClosed {
		return 0, false
	}
	return a.Amount, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.IsClosed || a.Amount+amount < 0 {
		return 0, false
	}
	a.Amount += amount
	return a.Amount, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.IsClosed {
		return 0, false
	}
	amount := a.Amount
	a.Amount = 0
	a.IsClosed = true
	return amount, true
}
