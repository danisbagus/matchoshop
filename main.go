package main

import (
	"net/http"

	app "github.com/danisbagus/matchoshop/app/api"
)

// func main() {
// 	app.StartApp()
// }

func Handler(w http.ResponseWriter, r *http.Request) {
	app.StartApp()
}
