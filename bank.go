package bank

var balnace int

func Deposit(amount int) {
	balnace = balnace + amount
}

func Balance() int {
	return balnace
}
