package val

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	go func() {
		for i := 0; i < 10000; i++ {
			Deposit(1)
		}
		done <- struct{}{}
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			Deposit(1)
		}
		done <- struct{}{}
	}()

	<-done
	<-done
	fmt.Println(Balance())
}

func TestWithdraw(t *testing.T) {
	for i := 0; i < 10000; i++ {
		Deposit(1)
	}

	fmt.Println(Balance())

	done := make(chan struct{})

	go func() {
		for i := 0; i < 10000/2; i++ {
			Withdraw(1)
		}
		done <- struct{}{}
	}()
	go func() {
		for i := 0; i < 10000/2; i++ {
			Withdraw(1)
		}
		done <- struct{}{}
	}()

	<-done
	<-done
	fmt.Println(Balance())
}
