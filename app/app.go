package app

import (
	"context"
	handlersmicroservices "github.com/pauliusluksys/golang-Reddit/handlers/microservices"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//func Start() {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatalf("Some error occured. Err: %s", err)
//	}
//	//db := domain.GormDbConnections()
//	//dbSqlx := domain.SqlxDbConnections()
//	//migrations.PostMigration()
//	//seeds.Execute(dbSqlx, "PostSeed")
//	domain.PostGorm{}.TableName()
//
//	r := routes()
//
//	err = r.Run(":8080")
//	if err != nil {
//		panic("Gin routing has failed: " + err.Error())
//	}
//}
func Microservices() {
	//fmt.Println(runtime.NumCPU())
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlersmicroservices.NewHello(l)
	//gh := handlersmicroservices.NewGoodbye(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	//sm.Handle("/goodbye", gh)
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	// wrapping ListenAndServe in gofunc so it's not going to block
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	// make a new channel to notify on os interrupt of server (ctrl + C)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	// This blocks the code until the channel receives some message
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	// Once message is consumed shut everything down
	// Gracefully shuts down all client requests. Makes server more reliable
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)

}
