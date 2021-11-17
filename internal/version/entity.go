package version

// version represents data about a record service.
type Version struct {
    ID          string  `json:"id"`
    Name        string  `json:"name"`
    ServiceId   int     `json:"service_id"`
    Enabled     bool    `json:"enabled"`
}