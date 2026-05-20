package students

type Student struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	City      string `json:"city"`
	Course    string `json:"course"`
	IsMale    bool   `json:"is_male"`
}