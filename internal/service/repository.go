package service

import (
	"log"
	"fmt"

	"database/sql"
    _ "github.com/lib/pq"
)

// repository persists services in database
type repository struct {
	db     *sql.DB
	logger *log.Logger
}

// NewRepository creates a new service repository
func NewRepository(db *sql.DB, logger *log.Logger) repository {
	return repository{db, logger}
}

func (r repository) list(filter, sort string, offset, limit int) ([]Service, error) {
    // A services slice to hold data from returned rows.
    var services []Service

	sql := "SELECT s.id, s.name, s.description, s.user_id, v.name FROM services s LEFT JOIN versions v ON s.id = v.service_id WHERE v.enabled = true"
	if (filter != "") {
		sql = fmt.Sprintf("%s AND s.name ILIKE '%%%s%%' OR s.description ILIKE '%%%s%%'", sql, filter, filter)
	}
	if (sort != "") {
		sql = fmt.Sprintf("%s ORDER BY s.%s", sql, sort )
	}
	
	sql = fmt.Sprintf("%s OFFSET %d LIMIT %d", sql, offset, limit)

    rows, err := r.db.Query(sql)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var service Service
        if err := rows.Scan(&service.ID, &service.Name, &service.Description, &service.UserId, &service.Version); err != nil {
            return nil, err
        }
        services = append(services, service)
    }

	return services, nil
}

func (r repository) get(id int) (Service, error) {
	// A service to hold data from the returned row.
	var service Service

	row := r.db.QueryRow("SELECT s.id, s.name, s.description, s.user_id, v.name FROM services s LEFT JOIN versions v ON s.id = v.service_id WHERE v.enabled = true AND s.id = $1", id)

    if err := row.Scan(&service.ID, &service.Name, &service.Description, &service.UserId, &service.Version); err != nil {
		return Service{}, err
    }
	return service, nil
}

func (r repository) create(newService Service) (int, error) {
    var id int
    err := r.db.QueryRow("INSERT INTO services (name, description, user_id) VALUES ($1, $2, $3) RETURNING id", newService.Name, newService.Description, newService.UserId).Scan(&id)

	return id, err
}

func (r repository) count() (int, error) {
    var count int
    err := r.db.QueryRow("SELECT COUNT(*) FROM services").Scan(&count)

	return count, err
}