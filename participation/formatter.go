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
