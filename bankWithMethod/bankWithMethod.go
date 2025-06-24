package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма должна быть положительной")
	}
	b.Balance += amount
	return nil
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма должна быть положительной")
	}
	if amount > b.Balance {
		return errors.New("недостаточно средств на счете")
	}
	b.Balance -= amount
	return nil
}

func (b BankAccount) GetBalance() float64 {
	return b.Balance
}

func (b BankAccount) PrintInfo() {
	fmt.Printf("Владелец счёта: %s\n", b.Owner)
	fmt.Printf("Текущий баланс: %.2f руб.\n", b.GetBalance())
}

func main() {
	account := BankAccount{
		Owner:   "Алексей Смирнов",
		Balance: 1000.00,
	}

	account.PrintInfo()
	fmt.Println()

	err := account.Deposit(500)
	if err != nil {
		fmt.Println("Ошибка пополнения:", err)
	} else {
		fmt.Println("Счёт успешно пополнен на 500 руб.")
		account.PrintInfo()
	}

	fmt.Println()

	err = account.Withdraw(200)
	if err != nil {
		fmt.Println("Ошибка снятия:", err)
	} else {
		fmt.Println("Со счёта снято 200 руб.")
		account.PrintInfo()
	}

	fmt.Println()

	err = account.Withdraw(2000)
	if err != nil {
		fmt.Println("Ошибка снятия:", err)
	} else {
		fmt.Println("Со счёта снято 2000 руб.")
		account.PrintInfo()
	}
}
