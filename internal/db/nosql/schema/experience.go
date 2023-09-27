package schema

type Experience struct {
	Job         string
	Company     string
	StartDate   string
	EndDate     string
	Description string
}

func (e *Experience) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"job":         e.Job,
		"company":     e.Company,
		"start_date":  e.StartDate,
		"end_date":    e.EndDate,
		"description": e.Description,
	}
}
