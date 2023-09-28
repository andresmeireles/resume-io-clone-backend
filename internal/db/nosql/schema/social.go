package schema

type Social struct {
	Name string
	Link string
	Hide bool
}

func (s *Social) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name": s.Name,
		"link": s.Link,
		"hide": s.Hide,
	}
}
