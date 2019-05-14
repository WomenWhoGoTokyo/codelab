package taskmanager

// Parameter is the structure of the json parameter.
type Parameter struct {
	ID     int64  `json:"ID"`
	Title  string `json:"title"`
	Status Status `json:"status"`
}
