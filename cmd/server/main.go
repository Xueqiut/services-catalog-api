package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "database/sql"
    _ "github.com/lib/pq"

    "github.com/Xueqiut/services-catalog-api/internal/service"
    "github.com/Xueqiut/services-catalog-api/internal/version"
)

func setupRouter(db *sql.DB, logger *log.Logger) *gin.Engine {
	// Initialize a Gin router using Default.
    r := gin.Default()
    v1 := r.Group("api/v1") 
    {
        service.RegisterHandlers(v1, service.NewRepository(db, logger), logger)
        version.RegisterHandlers(v1, version.NewRepository(db, logger), logger)
    }
    
	return r
}

func main() {
	connStr := "postgres://postgres:password@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }

    logger := log.Default()
    logger.SetPrefix("services-catalog-api")

    r := setupRouter(db, logger)
	r.Run("localhost:8080")
}