package handlers

import (
    "database/sql"
    "net/http"
    "github.com/DauletBai/real-estate-app/internal/repository"
    "encoding/json"
)

func GetProperties(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        properties, err := repository.GetAllProperties(db)
        if err != nil {
            http.Error(w, "Unable to fetch properties", http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(properties)
    }
}