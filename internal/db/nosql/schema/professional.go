package schema

type Professional struct {
	DesiredJob   string `json:"desired_job"`
	Presentation string
}

func (p *Professional) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"desiredJob":   p.DesiredJob,
		"presentation": p.Presentation,
	}
}
