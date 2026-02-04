package main

import (
	"fmt"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/account/internal/bank"
	"github.com/NutcrackerCom/go-backend-journey/mini-projects/account/internal/logger"
)

func main() {
	var logg logger.ConsoleLogger = logger.ConsoleLogger{}
	var service bank.Service = *bank.NewService(logg)
	service.Create(0, "A", 100)
	service.Create(1, "B", 100)
	service.Transfer(0, 1, 50)
	acc0, _ := service.Get(0)
	acc1, _ := service.Get(0)
	fmt.Println(acc0, acc1)
}
