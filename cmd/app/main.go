package main

import (
	"log"
	"net/http"
	"os"
	"api_calculate/internal/api"
	"api_calculate/internal/api/calc"
	"api_calculate/internal/api/greetings"
	"api_calculate/pkg/router"
)

func main() {
	logger := log.New(os.Stdout, "[LOG] ", log.Lshortfile)

	r := router.New(logger)

	calcH := calc.NewCalc(logger)

	r.Add(
		router.POST("/calc", calcH.Handle).SetErrHandler(api.ErrHandler),
	)

	r.Add(
		router.NewGroup("/greetings",
			router.POST("/hello", greetings.Hello),
			router.NewGroup("/sub",
				router.POST("/hello", greetings.Hello),
				router.GET("/hello", greetings.HelloGet),
			),
		),
	)

	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: r,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}