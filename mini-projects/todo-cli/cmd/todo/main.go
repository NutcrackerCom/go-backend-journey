package main

import (
	"fmt"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/todo-cli/internal/todo"
)

func main() {

	service := todo.NewService()
	service.Add("t0")
	service.Add("t1")
	service.Add("t2")
	fmt.Println(service.List())
	service.Done(0)
	service.Done(2)
	fmt.Println(service.List())
	service.Delete(0)
	fmt.Println(service.List())
	service.Add("t4")
	fmt.Println(service.List())
	/*var add string
	var list string
	var done string
	var delete string
	flag.StringVar(&add, "add", "", "-add")
	flag.StringVar(&list, "list", "", "-list")
	flag.StringVar(&done, "done", "", "-done")
	flag.StringVar(&delete, "delete", "", "-delete")
	flag.Parse()

	var service todo.Service = todo.NewService()
	for {
		if add != "" {
			service.Add(add)
			fmt.Println("ADDED", service.List())
		} else if list != "" {
			fmt.Println(service.List())
		} else if i, err := strconv.Atoi(done); err == nil {
			service.Done(i)
		} else if i, err := strconv.Atoi(delete); err == nil {
			service.Delete(i)
		} else {
			flag.Usage()
		}
	}*/
}
