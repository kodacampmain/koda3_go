package checkout

type PaymentMethod interface {
	Pay(prices []int)
}

func Checkout(paymentService PaymentMethod, prices []int) {
	paymentService.Pay(prices)
}
