package app

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	handlersmicroservices "github.com/pauliusluksys/golang-Reddit/handlers/microservices"
	userHandler "github.com/pauliusluksys/golang-Reddit/handlers/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
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

type Message struct {
	Greeting string `json:"greeting"`
	Response string `json:"response"`
}

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	wsConn *websocket.Conn
)

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	var err error
	wsConn, err = wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("could not upgrade: %s\n", err.Error())
		return
	}
	defer wsConn.Close()
	var number int = 12
	for {
		var msg Message
		//msg.response = "Good night from the server!"
		err := wsConn.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("error reading JSON: %s\n", err.Error())
			break
		}
		fmt.Printf("Message Received: %s\n", msg.Greeting)
		number++
		SendMessage(msg.Greeting + strconv.Itoa(number))
	}
}

func SendMessage(msg string) {
	fmt.Println("WORKS UP UNTIL THIS POINT WITH MSG: " + msg)
	err := wsConn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		fmt.Printf("error sending message: %s\n", err.Error())
	}
}
func Start() {
	gormDb := GormDbConnections()
	router := mux.NewRouter()
	router.HandleFunc("/api/auth/login", userHandler.UserLogin(gormDb)).Methods("POST")
	router.HandleFunc("/api/auth/signup", userHandler.UserSignup(gormDb)).Methods("POST")

	//router.HandleFunc("/api/auth/create-user", ).Methods("POST")
	//router.HandleFunc("api/something", utils.CheckTokenHandler(v1.GetSomething)).Methods("GET")
	//router.HandleFunc("/socket", WsEndpoint)
	log.Fatal(http.ListenAndServe(":9100", router))

}
func GormDbConnections() *gorm.DB {
	//dbSqlx := domain.SqlxDbConnections()
	//seeds.Execute(dbSqlx, "UserSeed")
	var myEnv map[string]string
	myEnv, err := godotenv.Read()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", myEnv["DB_USER"], myEnv["DB_PASSWORD"], myEnv["DB_ADDR"], myEnv["DB_PORT"], myEnv["DB_NAME"])
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
