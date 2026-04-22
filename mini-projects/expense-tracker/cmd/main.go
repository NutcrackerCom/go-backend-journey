package main

import (
	"log"
	"os"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/expense-tracker/internal/server"
)

func main() {
	flog, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer flog.Close()
	myLog := log.New(flog, `serv `, log.LstdFlags|log.Lshortfile)

	srv := server.NewServer(myLog)
	if err := srv.Http.ListenAndServe(); err != nil {
		myLog.Fatalf("Error %v", err)
		return
	}
}
