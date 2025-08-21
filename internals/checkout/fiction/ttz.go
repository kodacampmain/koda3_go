package fiction

import (
	"log"
)

type TTZ struct {
	Name     string
	Payments []int
}

func NewTTZ(name string) *TTZ {
	return &TTZ{
		Name:     name,
		Payments: make([]int, 0),
	}
}

func (t *TTZ) Pay(prices []int) {
	var total int
	for _, v := range prices {
		if v <= 0 {
			log.Println("Harga tidak boleh dibawah 0")
			return
		}
		total += v
	}
	t.Payments = append(t.Payments, total)
}
