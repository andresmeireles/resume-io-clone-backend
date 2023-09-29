package schema

type Education struct {
	Degree      string
	Institution string
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Description string
	Hide        bool
}

func (e *Education) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"degree":      e.Degree,
		"institution": e.Institution,
		"start_date":  e.StartDate,
		"end_date":    e.EndDate,
		"description": e.Description,
		"hide":        e.Hide,
	}
}
