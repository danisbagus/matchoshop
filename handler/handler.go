package handler

import (
	"fmt"
	"net/http"
)

// func Handler(w http.ResponseWriter, r *http.Request) {
// 	app.StartApp()
// }

func Handler() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "🚀 Hello from Golang on Vercel!")
	})
	http.ListenAndServe(":8080", nil) // Wajib ada server HTTP
}
