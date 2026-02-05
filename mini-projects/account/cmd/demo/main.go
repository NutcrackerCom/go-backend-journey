package main

import (
	"fmt"
	"sync"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/account/internal/bank"
)

func main() {
	//var logg logger.ConsoleLogger = logger.ConsoleLogger{}
	var service *bank.Service = bank.NewService(nil)
	service.Create(0, "A", 100000)
	service.Create(1, "B", 100000)
	wg := &sync.WaitGroup{}
	for range 100000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			service.Transfer(0, 1, 1)
		}()
	}
	wg.Wait()
	first, _ := service.Get(0)
	second, _ := service.Get(1)
	fmt.Println(first, second)

}
