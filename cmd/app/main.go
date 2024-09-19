package main

import (
    "log"
    "net/http"
    "github.com/DauletBai/real-estate-app/config"
    "github.com/DauletBai/real-estate-app/internal/handlers"
    "github.com/DauletBai/internal/middlewares"
    "github.com/DauletBai/real-estate-app/pkg/db"
    "github.com/go-chi/chi/v5"
)

func main() {
    // Загрузка конфигураций
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Инициализация подключения к базе данных
    dbConn, err := db.Connect(cfg.DatabaseURL)
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    defer dbConn.Close()

    // Инициализация маршрутов
    r := chi.NewRouter()

    // Логирование всех запросов
    r.Use(middlewares.Logger)

    // Пример защищённого маршрута
    r.With(middlewares.AuthMiddleware).Get("/secure", handlers.SecureEndpoint)

    // Пример маршрута для получения списка недвижимости
    r.Get("/properties", handlers.GetProperties(dbConn))

    log.Printf("Server is running on port %s", cfg.ServerPort)
    http.ListenAndServe(cfg.ServerPort, r)
}