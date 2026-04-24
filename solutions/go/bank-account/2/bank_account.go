package bankaccount

import "sync"

type Account struct {
	mu     sync.Mutex
	amount int64
	opened bool
}

func Open(amt int64) *Account {
	if amt < 0 {
		return nil
	}
	return &Account{
		amount: amt,
		opened: true,
	}
}

func (a *Account) Balance() (bal int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.opened {
		return 0, false
	}
	return a.amount, true
}

func (a *Account) Deposit(amt int64) (newBal int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.opened {
		return 0, false
	}
	if amt < 0 && a.amount+amt < 0 {
		return 0, false
	}
	a.amount += amt
	return a.amount, true
}

func (a *Account) Withdraw(amt int64) (newBal int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.opened || amt < 0 || amt > a.amount {
		return 0, false
	}
	a.amount -= amt
	return a.amount, true
}

func (a *Account) Close() (pay int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.opened {
		return 0, false
	}
	pay = a.amount
	a.amount = 0
	a.opened = false
	return pay, true
}