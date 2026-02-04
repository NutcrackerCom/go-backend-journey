package main

import (
	"fmt"
	"sync"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/account"
)

func main() {
	//var logger account.ConsoleLogger
	wg := new(sync.WaitGroup)
	var accUserA account.Account = account.Account{Owner: "A", Balance: 100000}
	var accUserB account.Account = account.Account{Owner: "B", Balance: 100}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			accUserA.Transfer(&accUserB, 10)
		}(wg)
	}

	wg.Wait()
	fmt.Println(accUserA, accUserB)
}
