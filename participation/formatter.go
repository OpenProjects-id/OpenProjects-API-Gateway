package participation

import "time"

type ProjectParticipationsFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatProjectParticipation(participation Participation) ProjectParticipationsFormatter {
	formatter := ProjectParticipationsFormatter{}
	formatter.ID = participation.ID
	formatter.Name = participation.User.Name
	formatter.CreatedAt = participation.CreatedAt

	return formatter
}

func FormatProjectParticipations(participations []Participation) []ProjectParticipationsFormatter {
	if len(participations) == 0 {
		return []ProjectParticipationsFormatter{}
	}

	var participationsFormatter []ProjectParticipationsFormatter

	for _, participation := range participations {
		formatter := FormatProjectParticipation(participation)
		participationsFormatter = append(participationsFormatter, formatter)
	}

	return participationsFormatter
}

type UserParticipationFormatter struct {
	ID        int              `json:"id"`
	Status    string           `json:"status"`
	CreatedAt time.Time        `json:"created_at"`
	Project   ProjectFormatter `json:"project"`
}

type ProjectFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserParticipation(participation Participation) UserParticipationFormatter {
	formatter := UserParticipationFormatter{}
	formatter.ID = participation.ID
	formatter.Status = participation.Status
	formatter.CreatedAt = participation.CreatedAt

	projectFormatter := ProjectFormatter{}
	projectFormatter.Name = participation.Project.Name
	projectFormatter.ImageURL = ""

	if len(participation.Project.ProjectImages) > 0 {
		projectFormatter.ImageURL = participation.Project.ProjectImages[0].FileName
	}

	formatter.Project = projectFormatter

	return formatter
}

func FormatUserParticipations(participations []Participation) []UserParticipationFormatter {
	if len(participations) == 0 {
		return []UserParticipationFormatter{}
	}

	var participationsFormatter []UserParticipationFormatter

	for _, participation := range participations {
		formatter := FormatUserParticipation(participation)
		participationsFormatter = append(participationsFormatter, formatter)
	}

	return participationsFormatter
}
