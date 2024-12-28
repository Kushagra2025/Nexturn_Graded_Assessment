package students

type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
	GPA     int    `json:"GPA"`
	Major   string `json:"Major"`
}
