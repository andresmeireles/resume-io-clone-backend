package collection

import (
	"github.com/andresmeireles/resume/internal/db/nosql/collection/user"
	"github.com/andresmeireles/resume/internal/db/nosql/schema"
)

func GetResumeByUserId(userCollection user.User, userId int) (schema.Resume, error) {
	personal, err := userCollection.GetPersonal(userId)
	socials, err := userCollection.GetSocial(userId)
	skills, err := userCollection.GetAllSkills(userId)
	experiences, err := userCollection.GetExperiences(userId)
	professional, err := userCollection.GetProfessional(userId)
	education, err := userCollection.GetEducation(userId)

	if err != nil {
		return schema.Resume{}, err
	}

	resume := schema.Resume{
		Skills:       skills,
		Personal:     personal,
		Experience:   experiences,
		Social:       socials,
		Profissional: professional,
		Education:    education,
	}
	return resume, nil
}
