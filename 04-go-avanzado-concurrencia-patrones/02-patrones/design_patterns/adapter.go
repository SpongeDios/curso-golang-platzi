package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct {
}

func (c *CashPayment) Pay()  {
	fmt.Println("Payment using Cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct {
}

func (p BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying using BankAccount %d", bankAccount)
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (b BankPaymentAdapter) Pay() {
	b.BankPayment.Pay(b.bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)
	bank := &BankPayment{}
	bankAdapted := &BankPaymentAdapter{BankPayment: bank, bankAccount: 123456789}
	//ProcessPayment(bank)
	ProcessPayment(bankAdapted)
}