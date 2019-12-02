package val

import (
	"sync"
)

// var sema = make(chan struct{}, 1)
// var balance int

// func Deposit(amount int) {
// 	sema <- struct{}{}
// 	balance += amount
// 	<-sema
// }

// func Balance() int {
// 	sema <- struct{}{}
// 	defer func() { <-sema }()
// 	return balance
// }

var balance int
var mu sync.RWMutex

func Deposit(amount int) {
	mu.Lock()
	deposit(amount)
	mu.Unlock()
}
func deposit(anount int) {
	balance += anount
}

func Balance() int {
	mu.RLock()  // 读锁
	defer mu.RUnlock()
	return balance
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}
