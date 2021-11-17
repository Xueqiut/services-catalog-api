package version

import (
	"log"

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

func (r repository) list(serviceId int) ([]Version, error) {
	var versions []Version

    rows, err := r.db.Query("SELECT * FROM versions WHERE service_id = $1", serviceId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var version Version
        if err := rows.Scan(&version.ID, &version.Name, &version.ServiceId, &version.Enabled); err != nil {
            return nil, err
        }
        versions = append(versions, version)
    }

	return versions, nil
}