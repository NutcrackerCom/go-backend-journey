package main

import (
	"fmt"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/account"
)

func main() {
	var logger *account.Collector
	var accUserA account.Account = account.Account{Owner: "A", Balance: 100, log: logger}
	var accUserB account.Account = account.Account{Owner: "B", Balance: 100}
	accUserA.Transfer(&accUserB, 35)
	fmt.Println(accUserA, accUserB)
}
