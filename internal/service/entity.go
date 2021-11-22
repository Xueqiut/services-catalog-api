package service

// service represents data about a service.
type Service struct {
  ID          string  `json:"id"`
  Name        string  `json:"name"`
  Description string  `json:"description"`
  UserId      int 	`json:"user_id"`
	Version		string	`json:"version"`
}