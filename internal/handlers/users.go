package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sankalpmukim/url-shortener-go/internal/database"
	"github.com/sankalpmukim/url-shortener-go/pkg/logs"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := database.DB.GetUsers()
	if err != nil {
		logs.Error("Error getting users", err)
		w.Write([]byte("Error"))
	}

	logs.Info("Users", users)

	// return users json
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		logs.Error("Error encoding users", err)
		w.Write([]byte("Error"))
	}
}
