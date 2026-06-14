package service

import (
	"github.com/paulo-granthon/portfolio_apis/models"
	"strings"
	"testing"
)

func strptr(s string) *string { return &s }

func sampleProject(semester uint8, name string) models.Project {
	return models.Project{
		Id:          uint64(semester),
		Name:        name,
		Semester:    semester,
		Company:     "FATEC",
		Summary:     name + " summary",
		Description: name + " description",
		Url:         "github.com/test/" + name,
	}
}

func TestRenderMarkdown_StructureAndOrder(t *testing.T) {
	image := "https://example.com/banner.png"
	summary := "A backend developer."
	github := "paulo-granthon"
	portfolio := models.Portfolio{
		User: models.User{
			Name:               "Paulo Granthon",
			Summary:            &summary,
			SemesterMatriculed: &models.YearSemester{Year: 2022, Semester: 2},
			GithubUsername:     &github,
		},
		Projects: []models.PortfolioProject{
			{
				Project:       func() models.Project { p := sampleProject(1, "Khali"); p.Image = &image; return p }(),
				Contributions: []models.PortfolioContribution{{Id: 1, Title: "Dockerization", Content: "Did docker.", Skills: []string{"Docker", "Bash"}}},
			},
			{
				Project:       sampleProject(2, "api2"),
				Contributions: nil,
			},
		},
	}

	md := RenderMarkdown(portfolio)

	for _, want := range []string{
		"# Portfólio — Paulo Granthon",
		"![Paulo Granthon](https://github.com/paulo-granthon.png?size=200)",
		"Matriculado em 2022/2º semestre.",
		"A backend developer.",
		"## 1º Semestre — Khali",
		"![Khali](https://example.com/banner.png)",
		"- **Empresa:** FATEC",
		"- **Repositório:** <https://github.com/test/Khali>",
		"#### Dockerization",
		"**Habilidades:** Docker, Bash",
		"## 2º Semestre — api2",
		"_Nenhuma contribuição especificada._",
	} {
		if !strings.Contains(md, want) {
			t.Errorf("markdown missing %q\n---\n%s", want, md)
		}
	}

	// Semester 1 must be rendered before semester 2.
	if strings.Index(md, "## 1º Semestre") > strings.Index(md, "## 2º Semestre") {
		t.Errorf("projects not ordered by semester:\n%s", md)
	}
}

func TestRenderMarkdown_OmitsOptionalFields(t *testing.T) {
	// No summary, no matriculation, no image, no url, no skills.
	portfolio := models.Portfolio{
		User: models.User{Name: "No Frills"},
		Projects: []models.PortfolioProject{
			{
				Project:       models.Project{Name: "bare", Semester: 3, Company: "ACME"},
				Contributions: []models.PortfolioContribution{{Id: 1, Title: "Untitled work", Content: "content"}},
			},
		},
	}

	md := RenderMarkdown(portfolio)

	for _, absent := range []string{"Matriculado em", "![", "Repositório", "Habilidades"} {
		if strings.Contains(md, absent) {
			t.Errorf("markdown should not contain %q for bare data\n---\n%s", absent, md)
		}
	}
	if !strings.Contains(md, "#### Untitled work") {
		t.Errorf("expected contribution title rendered\n%s", md)
	}
}
