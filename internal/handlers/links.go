package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sankalpmukim/url-shortener-go/internal/database"
	"github.com/sankalpmukim/url-shortener-go/pkg/logs"
)

func GetLinks(w http.ResponseWriter, r *http.Request) {
	links, err := database.DB.GetLinks()
	if err != nil {
		logs.Error("Error getting links", err)
		w.Write([]byte("Error"))
		return
	}

	logs.Info("Link", links)

	// return links json
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(links); err != nil {
		logs.Error("Error encoding links", err)
		w.Write([]byte("Error"))
	}
}
