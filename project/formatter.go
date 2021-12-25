package project

import "strings"

type ProjectFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	TechStacks       string `json:"tech_stacks"`
	ParticipantCount int    `json:"participant_count"`
	TotalBudget      int    `json:"total_budget"`
	ImageURL         string `json:"image_url"`
	Slug             string `json:"slug"`
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
	projectFormatter.Slug = project.Slug

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

type ProjectDetailFormatter struct {
	ID               int                     `json:"id"`
	Name             string                  `json:"name"`
	ShortDescription string                  `json:"short_description"`
	Description      string                  `json:"description"`
	TechStacks       string                  `json:"tech_stacks"`
	ParticipantCount int                     `json:"participant_count"`
	ImageURL         string                  `json:"image_url"`
	TotalBudget      int                     `json:"total_budget"`
	UserID           int                     `json:"user_id"`
	Slug             string                  `json:"slug"`
	Perks            []string                `json:"perks"`
	User             ProjectUserFormatter    `json:"user"`
	Images           []ProjectImageFormatter `json:"images"`
}

type ProjectUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type ProjectImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatProjectDetail(project Project) ProjectDetailFormatter {
	projectDetailFormatter := ProjectDetailFormatter{}
	projectDetailFormatter.ID = project.ID
	projectDetailFormatter.Name = project.Name
	projectDetailFormatter.ShortDescription = project.ShortDescription
	projectDetailFormatter.Description = project.Description
	projectDetailFormatter.TechStacks = project.TechStacks
	projectDetailFormatter.ParticipantCount = project.ParticipantCount
	projectDetailFormatter.TotalBudget = project.TotalBudget
	projectDetailFormatter.ImageURL = ""
	projectDetailFormatter.Slug = project.Slug
	projectDetailFormatter.UserID = project.UserID

	if len(project.ProjectImages) > 0 {
		projectDetailFormatter.ImageURL = project.ProjectImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(project.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	projectDetailFormatter.Perks = perks

	user := project.User
	projectUserFormatter := ProjectUserFormatter{}
	projectUserFormatter.Name = user.Name
	projectUserFormatter.ImageURL = user.AvatarFileName

	projectDetailFormatter.User = projectUserFormatter

	images := []ProjectImageFormatter{}

	for _, image := range project.ProjectImages {
		projectImageFormatter := ProjectImageFormatter{}
		projectImageFormatter.ImageURL = image.FileName

		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		}
		projectImageFormatter.IsPrimary = isPrimary

		images = append(images, projectImageFormatter)
	}

	projectDetailFormatter.Images = images

	return projectDetailFormatter
}
