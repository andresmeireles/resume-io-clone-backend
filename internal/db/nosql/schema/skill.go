package schema

import "strings"

type Skill struct {
	Name  string
	Level string
	Hide  bool
}

func (s *Skill) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":  strings.Trim(s.Name, " "),
		"level": s.Level,
		"hide":  s.Hide,
	}
}
