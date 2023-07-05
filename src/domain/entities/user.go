package entities

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Tasks []Task `json:"tasks,omitempty"`
}
