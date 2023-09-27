package schema

type Personal struct {
	FirstName string
	LastName  string
	Email     string
	Age       int
	Gender    string
	Phone     string
	Address   string // rever essa parada
}

func (p Personal) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"firstName": p.FirstName,
		"lastName":  p.LastName,
		"email":     p.Email,
		"age":       p.Age,
		"gender":    p.Gender,
		"phone":     p.Phone,
		"address":   p.Address,
	}
}
