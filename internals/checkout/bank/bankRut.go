package bank

import "log"

type BankRut struct {
	Name string
}

func NewBankRut(name string) *BankRut {
	return &BankRut{
		Name: name,
	}
}

func (b *BankRut) Pay(prices []int) {
	var total int
	for _, v := range prices {
		total += v
	}
	log.Printf("Pembayaran sebesar %d berhasil dilakukan melalui %s", total, b.Name)
}
