package handlers

import (
	"net/http"
)

func GetDocs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/app/docs/swagger.json")
}
