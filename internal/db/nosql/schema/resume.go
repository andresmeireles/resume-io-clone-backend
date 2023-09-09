package schema

type Resume struct {
	Personal     Personal
	Profissional Professional
	Skills       []Skill
	Education    []Education
	Experience   []Experience
	Social       []Social
}
