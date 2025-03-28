package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// func Handler(w http.ResponseWriter, r *http.Request) {
// 	app.StartApp()
// }

// func Handler() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "🚀 Hello from Golang on Vercel!")
// 	})
// 	http.ListenAndServe(":8080", nil) // Wajib ada server HTTP
// }

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")

	e := echo.New()

	health := e.Group("/health")
	health.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	// http.ListenAndServe(":8080", nil)
	e.Logger.Fatal(e.Start(":" + "8080"))
}
