package service

import (
	"fmt"
	"github.com/paulo-granthon/portfolio_apis/models"
	"strings"
)

// RenderMarkdown serializes a Portfolio document to markdown. It is a pure
// function (document in, string out) so the HTTP handler and the CLI share the
// exact same output, and so it can be unit-tested without a database.
func RenderMarkdown(portfolio models.Portfolio) string {
	var b strings.Builder

	fmt.Fprintf(&b, "# Portfólio — %s\n", portfolio.User.Name)

	if u := portfolio.User.GithubUsername; u != nil && *u != "" {
		fmt.Fprintf(&b, "\n![%s](https://github.com/%s.png?size=200)\n", portfolio.User.Name, *u)
	}

	if portfolio.User.SemesterMatriculed != nil {
		m := portfolio.User.SemesterMatriculed
		fmt.Fprintf(&b, "\nMatriculado em %d/%dº semestre.\n", m.Year, m.Semester)
	}

	if portfolio.User.Summary != nil && *portfolio.User.Summary != "" {
		fmt.Fprintf(&b, "\n%s\n", *portfolio.User.Summary)
	}

	for _, project := range portfolio.Projects {
		fmt.Fprintf(&b, "\n## %dº Semestre — %s\n", project.Semester, project.Name)

		if project.Image != nil && *project.Image != "" {
			fmt.Fprintf(&b, "\n![%s](%s)\n", project.Name, *project.Image)
		}

		fmt.Fprintf(&b, "\n- **Empresa:** %s\n", project.Company)
		if project.Url != "" {
			fmt.Fprintf(&b, "- **Repositório:** <https://%s>\n", strings.TrimPrefix(strings.TrimPrefix(project.Url, "https://"), "http://"))
		}

		if project.Summary != "" {
			fmt.Fprintf(&b, "\n%s\n", project.Summary)
		}

		if project.Description != "" {
			fmt.Fprintf(&b, "\n%s\n", project.Description)
		}

		b.WriteString("\n### Contribuições\n")
		if len(project.Contributions) == 0 {
			b.WriteString("\n_Nenhuma contribuição especificada._\n")
			continue
		}

		for _, contribution := range project.Contributions {
			fmt.Fprintf(&b, "\n#### %s\n", contribution.Title)
			if contribution.Content != "" {
				fmt.Fprintf(&b, "\n%s\n", contribution.Content)
			}
			if len(contribution.Skills) > 0 {
				fmt.Fprintf(&b, "\n**Habilidades:** %s\n", strings.Join(contribution.Skills, ", "))
			}
		}
	}

	return b.String()
}
