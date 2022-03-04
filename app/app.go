package app

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/pauliusluksys/golang-Reddit/domain"
	handlersmicroservices "github.com/pauliusluksys/golang-Reddit/handlers/microservices"
	userHandler "github.com/pauliusluksys/golang-Reddit/handlers/user"
	v1 "github.com/pauliusluksys/golang-Reddit/handlers/v1"
	"github.com/pauliusluksys/golang-Reddit/middlewares"
	"github.com/pauliusluksys/golang-Reddit/utils"
	"github.com/pauliusluksys/golang-Reddit/validation"
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

	//logger := utils.NewLogger()
	//configs := utils.NewConfigurations(logger)
	//validator := validation.NewValidation()
	//gormDb, err := domain.GormDbConnections(configs, logger)
	//mailService := servicesMail.NewSGMailService(logger, configs)
	//uh := handlers.NewAuthHandler(logger, configs, validator, gormDb, authService, mailService)
	//migrations.PostCommentMigration()
	//seeds.Execute(domain.SqlxDbConnections(), "PostCommentsSeed")
	logger := utils.NewLogger()
	configs, err := utils.NewConfigurations(logger)
	if err != nil {
		panic(err.Message)
	}
	// validator contains all the methods that are need to validate the user json in request
	validator := validation.NewValidation()

	// create a new connection to the postgres db store
	db, err := data.NewConnection(configs, logger)
	if err != nil {
		logger.Error("unable to connect to db", "error", err)
		panic(err)
	}

	gormDb := domain.GormDbConnections()
	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()
	auth := api.PathPrefix("/auth").Subrouter()
	email := auth.PathPrefix("/email").Subrouter()
	auth.HandleFunc("/login", userHandler.UserLogin(gormDb)).Methods("POST")
	auth.HandleFunc("/signup", userHandler.UserSignup(gormDb)).Methods("POST")
	auth.HandleFunc("/post", middlewares.CheckAuth(v1.PostH)).Methods("GET")
	auth.HandleFunc("/posts", middlewares.CheckAuth(v1.AllPostsH)).Methods("GET")
	auth.HandleFunc("/greet", middlewares.CheckAuth(v1.Greet)).Methods("GET")

	api.HandleFunc("/post", middlewares.CheckAuth(v1.PostH)).Methods("GET")
	api.HandleFunc("/post/comments", v1.PostComments).Methods("GET")
	api.HandleFunc("/post/comments/store", v1.PostCommentsStore).Methods("POST")

	email.HandleFunc("/verify", v1.VerifyEmail).Methods("POST")
	//router.HandleFunc("/api/auth/verify-email", v1.VerifyEmail).Methods("POST")
	//mailR := router.PathPrefix("/verify").Methods(http.MethodPost).Subrouter()
	//mailR.HandleFunc("/mail", uh.VerifyMail)
	//mailR.HandleFunc("/password-reset", uh.VerifyPasswordReset)
	//mailR.Use(uh.MiddlewareValidateVerificationData)
	//router.HandleFunc("/api/auth/create-user", ).Methods("POST")
	//router.HandleFunc("api/something", utils.CheckTokenHandler(v1.GetSomething)).Methods("GET")
	//router.HandleFunc("/socket", WsEndpoint)
	log.Fatal(http.ListenAndServe(":9100", setHeaders(router)))

}
func setHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//for k, v := range r.Header {
		//	_, err := fmt.Printf("Header field %q, Value %s\n", k, v[0])
		//	if err != nil {
		//		fmt.Printf(err.Error())
		//	}
		//}
		//anyone can make a CORS request (not recommended in production)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//only allow GET, POST, and OPTIONS
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		//Since I was building a REST API that returned JSON, I set the content type to JSON here.
		w.Header().Set("Content-Type", "application/json")
		//Allow requests to have the following headers
		w.Header().Set("Access-Control-Allow-Headers", "Sec-Fetch-Dest, Cache-Control, Access-Control-Request-Method, Pragma, Access-Control-Request-Headers, Origin, Sec-Fetch-Site, Connection, Sec-Fetch-Mode, Referer, Accept-Language, Accept,Accept-Encoding, authorization")

		//if it's just an OPTIONS request, nothing other than the headers in the response is needed.
		//This is essential because you don't need to handle the OPTIONS requests in your handlers now
		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
