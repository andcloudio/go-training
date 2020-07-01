package bank

var depositsStream = make(chan int)
var balancesStream = make(chan int)

// Deposit amount
func Deposit(amount int) { depositsStream <- amount }

// Balance return
func Balance() int { return <-balancesStream }

func teller() {
	var balance int
	for {
		select {
		case amount := <-depositsStream:
			balance += amount
		case balancesStream <- balance:
		}
	}
}

func init() {
	go teller()
}
