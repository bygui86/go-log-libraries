package main

import (
	// "fmt"
	// "log"
	"net/http"

	zerolog "github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
	zaplog "go.uber.org/zap"
)

func main() {
	zap, _ := zaplog.NewProduction()
	// zap, _ := zaplog.NewProduction().Sugar()
	defer zap.Sync() // flushes buffer, if any

	// single prints
	// http.HandleFunc("/fmt", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Print("hello")
	// })
	// http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Print("hello")
	// })
	http.HandleFunc("/logrus", func(w http.ResponseWriter, r *http.Request) {
		logrus.Print("hello")
	})
	http.HandleFunc("/zap", func(w http.ResponseWriter, r *http.Request) {
		zap.Info("hello")
	})
	http.HandleFunc("/zerolog", func(w http.ResponseWriter, r *http.Request) {
		zerolog.Print("hello")
	})

	// multiple prints (100)
	// http.HandleFunc("/fmt100", func(w http.ResponseWriter, r *http.Request) {
	// 	for i := 0; i < 100; i++ {
	// 		fmt.Print("hello")
	// 	}
	// })
	// http.HandleFunc("/log100", func(w http.ResponseWriter, r *http.Request) {
	// 	for i := 0; i < 100; i++ {
	// 		log.Print("hello")
	// 	}
	// })
	http.HandleFunc("/logrus100", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < 100; i++ {
			logrus.Print("hello")
		}
	})
	http.HandleFunc("/zap100", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < 100; i++ {
			zap.Info("hello")
		}
	})
	http.HandleFunc("/zerolog100", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < 100; i++ {
			zerolog.Print("hello")
		}
	})

	logrus.Info("Starting log comparison server...")
	logrus.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
