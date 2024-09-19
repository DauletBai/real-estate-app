package repository

import (
    "database/sql"
    "github.com/DauletBai/real-estate-app/internal/models"
)

func GetAllProperties(db *sql.DB) ([]models.Property, error) {
    rows, err := db.Query("SELECT id, name, price FROM properties")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var properties []models.Property
    for rows.Next() {
        var property models.Property
        if err := rows.Scan(&property.ID, &property.Name, &property.Price); err != nil {
            return nil, err
        }
        properties = append(properties, property)
    }

    return properties, nil
}