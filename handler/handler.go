package handler

import (
	"net/http"

	app "github.com/danisbagus/matchoshop/app/api"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	app.StartApp()
}
