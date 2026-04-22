package main

import (
	"log"
	"os"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/todo-list-server/internal/server"
)

func main() {
	flog, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer flog.Close()
	mylog := log.New(flog, `serv `, log.LstdFlags|log.Lshortfile)

	server := server.NewServer(mylog)
	if err := server.Http.ListenAndServe(); err != nil {
		mylog.Fatalf("Error %v", err)
		return
	}
}
