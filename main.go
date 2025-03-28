package main

import (
	"net/http"

	app "github.com/danisbagus/matchoshop/app/api"
)

// func main() {
// 	app.StartApp()
// }

func Main(w http.ResponseWriter, r *http.Request) {
	app.StartApp()
}
