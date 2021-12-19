package project

type ProjectFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	TechStacks       string `json:"tech_stacks"`
	ParticipantCount int    `json:"participant_count"`
	ImageURL         string `json:"image_url"`
	TotalBudget      int    `json:"total_budget"`
}

func FormatProject(project Project) ProjectFormatter {
	projectFormatter := ProjectFormatter{}
	projectFormatter.ID = project.ID
	projectFormatter.UserID = project.UserID
	projectFormatter.Name = project.Name
	projectFormatter.ShortDescription = project.ShortDescription
	projectFormatter.Description = project.Description
	projectFormatter.TechStacks = project.TechStacks
	projectFormatter.ParticipantCount = project.ParticipantCount
	projectFormatter.TotalBudget = project.TotalBudget
	projectFormatter.ImageURL = ""

	if len(project.ProjectImages) > 0 {
		projectFormatter.ImageURL = project.ProjectImages[0].FileName
	}

	return projectFormatter
}

func FormatProjects(projects []Project) []ProjectFormatter {
	projectsFormatter := []ProjectFormatter{}

	for _, project := range projects {
		projectFormatter := FormatProject(project)
		projectsFormatter = append(projectsFormatter, projectFormatter)
	}

	return projectsFormatter
}