package online

import "log"

type SindbadPay struct {
	Name string
}

func NewSindbadPay(name string) *SindbadPay {
	return &SindbadPay{
		Name: name,
	}
}

func (s *SindbadPay) Pay(prices []int) {
	var total int
	for _, v := range prices {
		total += v
	}
	log.Printf("Pembayaran sebesar %d berhasil dilakukan melalui %s", total, s.Name)
}
