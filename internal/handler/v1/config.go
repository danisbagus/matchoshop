package v1

import (
	"net/http"
	"os"

	"github.com/danisbagus/go-common-packages/http/response"
)

type ConfigHandler struct {
}

func (h ConfigHandler) GetPaypalConfig(w http.ResponseWriter, r *http.Request) {
	paypalClientID := os.Getenv("PAYPAL_CLIENT_ID")
	resData := map[string]interface{}{
		"client_id": paypalClientID,
	}
	response.Write(w, http.StatusOK, resData)
}
