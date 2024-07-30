package main

import (
	"database/sql"
	"fmt"
	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/fatih/color"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const port = "8080"

func main() {
	// connect to the database
	db := initDB()
	err := db.Ping()
	if err != nil {
		return
	}

	// create logger
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// create sessions
	session := initSession()

	// create channels

	// create wait group
	waitGroup := sync.WaitGroup{}

	// set up the application config
	app := Config{
		Session:  session,
		DB:       db,
		Wait:     &waitGroup,
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	// set up mail

	// listen for signal
	go app.listenForShutdown()

	// listen for web connection
	app.serve()
}

func (app *Config) serve() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}

	app.InfoLog.Println("Starting web server.")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func initDB() *sql.DB {
	connect := connectToDB()
	if connect == nil {
		color.Red("ERROR: Can't connect to database")
	}
	return connect
}

func connectToDB() *sql.DB {
	counts := 0

	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			color.Red("ERROR: Connection time out")
		} else {
			color.Green("INFO: Connected to database!")
			return connection
		}
		if counts > 10 {
			return nil
		}

		color.Green("INFO: Backing off for second.")
		time.Sleep(1 * time.Second)
		counts++

		continue
	}
}

func openDB(dns string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dns)
	if err != nil {
		return nil, err
	}

	return db, err
}

func initSession() *scs.SessionManager {
	session := scs.New()
	session.Store = redisstore.New(initRedis())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}

func initRedis() *redis.Pool {
	redisPool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS"))
		},
	}
	return redisPool
}

func (app *Config) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.shutdown()
	os.Exit(1)
}

func (app *Config) shutdown() {
	app.InfoLog.Println("Would run cleanup task.")
	app.Wait.Wait()
	app.InfoLog.Println("Shutdown!")
}
