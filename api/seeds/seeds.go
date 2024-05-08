package seeds

import (
	"models"
	"service"
	"storage"

	"github.com/ztrue/tracerr"
)

func Run() error {
	storage, err := storage.NewPostgreStorage()
	if err != nil {
		return tracerr.Errorf("failed to run seeds: failed to create storage: %w", tracerr.Wrap(err))
	}

	service, nil := service.NewService(storage)
	if err != nil {
		return tracerr.Errorf("failed to run seeds: failed to create service: %w", tracerr.Wrap(err))
	}

	rawDB, err := storage.GetRawDB()
	if err != nil {
		err = tracerr.Errorf("PostgreStorage.Migrate: error getting sql.DB from gorm: %w", err)
		tracerr.PrintSourceColor(err)
		return err
	}

	if _, err := rawDB.Exec(`
		CREATE EXTENSION IF NOT EXISTS uint;

		DROP TABLE IF EXISTS users CASCADE;
		DROP TABLE IF EXISTS teams CASCADE;
		DROP TABLE IF EXISTS team_users CASCADE;
		DROP TABLE IF EXISTS projects CASCADE;
		DROP TABLE IF EXISTS skills CASCADE;
		DROP TABLE IF EXISTS contributions CASCADE;
		DROP TABLE IF EXISTS contribution_skills CASCADE;

		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			summary VARCHAR(200) NULL,
			semester_matriculed JSONB NULL,
			github_username VARCHAR(39) NULL,
			password VARCHAR(50) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS teams (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS team_users (
			team_id INT NOT NULL,
			user_id INT NOT NULL,
			PRIMARY KEY (team_id, user_id),
			FOREIGN KEY (team_id) REFERENCES teams(id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		);

		CREATE TABLE IF NOT EXISTS projects (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			semester UINT1 NOT NULL,
			company VARCHAR(100) NOT NULL,
			team_id INT NOT NULL,
			summary TEXT NOT NULL,
			description VARCHAR(500) NOT NULL,
			url VARCHAR(100) NOT NULL,
			FOREIGN KEY (team_id) REFERENCES teams(id)
		);

		CREATE TABLE IF NOT EXISTS skills (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS contributions (
			id SERIAL PRIMARY KEY,
			project_id INT NOT NULL,
			user_id INT NOT NULL,
			title VARCHAR(50) NOT NULL,
			content TEXT NOT NULL,
			FOREIGN KEY (project_id) REFERENCES projects(id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		);

		CREATE TABLE IF NOT EXISTS contribution_skills (
			contribution_id INT NOT NULL,
			skill_id INT NOT NULL,
			PRIMARY KEY (contribution_id, skill_id),
			FOREIGN KEY (contribution_id) REFERENCES contributions(id),
			FOREIGN KEY (skill_id) REFERENCES skills(id)
		);
	`); err != nil {
		err = tracerr.Errorf("PostgreStorage.Migrate: error executing root migration: %w", err)
		tracerr.PrintSourceColor(err)
		return err
	}

	if err := UserMigrate(storage, *service); err != nil {
		return tracerr.Errorf("failed to run seeds: failed to migrate user: %w", tracerr.Wrap(err))
	}

	if err := TeamMigrate(storage, *service); err != nil {
		return tracerr.Errorf("failed to run seeds: failed to migrate team: %w", tracerr.Wrap(err))
	}

	if err := ProjectMigrate(storage, *service); err != nil {
		return tracerr.Errorf("failed to run seeds: failed to migrate project: %w", tracerr.Wrap(err))
	}

	if err := SkillMigrate(storage, *service); err != nil {
		return tracerr.Errorf("failed to run seeds: failed to migrate skill: %w", tracerr.Wrap(err))
	}

	if err := ContributionMigrate(storage, *service); err != nil {
		return tracerr.Errorf("failed to run seeds: failed to migrate contribution: %w", tracerr.Wrap(err))
	}

	return nil
}

func UserMigrate(
	storage storage.Storage,
	service service.Service,
) error {
	userModule, err := storage.GetUserModule()
	if err != nil {
		return tracerr.Errorf("failed to get user module: %w", tracerr.Wrap(err))
	}

	summary := "Backend developer intern at @gorilainvest | Database technologist student at FATEC | Self titled full-stack developer"
	yearSemesterMatriculed := models.NewYearSemester(uint16(2022), uint8(2))
	githubUsername := "paulo-granthon"

	exampleUsers := []models.CreateUser{
		models.NewCreateUser(
			"Paulo Granthon",
			"123456",
			&summary,
			&yearSemesterMatriculed,
			&githubUsername,
		),
	}

	for _, p := range exampleUsers {
		if _, err := userModule.Create(p); err != nil {
			return tracerr.Errorf("failed to insert user seeds: %w", tracerr.Wrap(err))
		}
	}

	return nil
}

func TeamMigrate(
	storage storage.Storage,
	service service.Service,
) error {
	teamModule, err := storage.GetTeamModule()
	if err != nil {
		return tracerr.Errorf("failed to get team module: %w", tracerr.Wrap(err))
	}

	exampleTeams := []models.CreateTeam{
		models.NewCreateTeam("Khali"),
	}

	var teamIdToAddMember *uint64

	for _, p := range exampleTeams {
		insertedTeamId, err := teamModule.Create(p)
		if err != nil {
			if teamIdToAddMember == nil {
				teamIdToAddMember = insertedTeamId
			}
			return tracerr.Errorf("failed to insert team seeds: %w", tracerr.Wrap(err))
		}
	}

	if teamIdToAddMember == nil {
		existingTeam, err := teamModule.GetById(1)
		if err != nil {
			return tracerr.Errorf("no team was inserted neither found in the database: %w", tracerr.Wrap(err))
		}

		teamIdToAddMember = &existingTeam.Id
	}

	if err := teamModule.AddUsers(*teamIdToAddMember, 1); err != nil {
		return tracerr.Errorf("failed to add users to team: %w", tracerr.Wrap(err))
	}

	return nil
}

func ProjectMigrate(
	storage storage.Storage,
	service service.Service,
) error {
	projectModule, err := storage.GetProjectModule()
	if err != nil {
		return tracerr.Errorf("failed to get project module: %w", tracerr.Wrap(err))
	}

	exampleProjects := []models.CreateProject{
		models.NewCreateProject(
			"Khali", 1, "FATEC", 1,
			"Avaliação 360",
			"A plataforma Khali permite a implementação do método de Avaliação 360° na Instituição de Ensino fictícia PBLTeX. Este projeto de API do 1º Semestre de Banco de Dados da Fatec - São José dos Campos possibilita uma abordagem abrangente na avaliação dos diversos aspectos da instituição, promovendo uma análise holística e aprimorando processos de gestão e desenvolvimento.",
			"github.com/taniacruzz/Khali",
		),
		models.NewCreateProject(
			"API2Semestre", 2, "2RP", 1,
			"Controle de Horas-Extras e Sobreavisos (desktop)",
			"A API desenvolvida no 2° semestre do curso de Banco de Dados na Fatec - SJC proporciona um sistema desktop especializado no registro de horas extras e sobreavisos pelos colaboradores, com funcionalidades de controle tanto para gestores (PO) quanto para administradores (RH e Financeiro). Essa solução oferece uma plataforma integrada e eficiente para gerenciamento de tempo e recursos humanos, contribuindo para uma gestão mais eficaz e transparente dentro da organização.",
			"github.com/projetoKhali/API2Semestre",
		),
		models.NewCreateProject(
			"api3", 3, "2RP", 1,
			"Controle de Horas-Extras e Sobreavisos (web)",
			"Sistema desenvolvido para auxiliar na gestão eficiente das horas trabalhadas pelos colaboradores de uma empresa. Ele automatiza a identificação e classificação de horas extras e sobreavisos, simplificando os processos de controle para os departamentos pessoal e financeiro.",
			"github.com/projetoKhali/api3",
		),
	}

	for _, p := range exampleProjects {
		if _, err := projectModule.Create(p); err != nil {
			return err
		}
	}

	return nil
}

func SkillMigrate(
	storage storage.Storage,
	service service.Service,
) error {
	skillModule, err := storage.GetSkillModule()
	if err != nil {
		return tracerr.Errorf("failed to get skill module: %w", tracerr.Wrap(err))
	}

	exampleSkills := []models.CreateSkill{
		models.NewCreateSkill("Scrum"),
		models.NewCreateSkill("Python"),
		models.NewCreateSkill("TKinter"),
		models.NewCreateSkill("Análise de Dados"),
		models.NewCreateSkill("Java"),
		models.NewCreateSkill("Spring"),
	}

	for _, sk := range exampleSkills {
		if _, err := skillModule.Create(sk); err != nil {
			return err
		}
	}

	return nil
}

func ContributionMigrate(
	storage storage.Storage,
	service service.Service,
) error {
	exampleContributions := []models.CreateContributionByNames{
		models.NewCreateContributionByNames(
			"API2Semestre", "paulo-granthon",
			"Teste de titulo 1", "Teste de conteúdo 1",
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Teste de titulo 2", "Teste de conteúdo 2",
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Teste de titulo 3", "Teste de conteúdo 3",
		),
	}

	for _, n := range exampleContributions {
		if _, err := service.ContributionService.Create(n); err != nil {
			return tracerr.Errorf("failed to create contribution: %w", tracerr.Wrap(err))
		}
	}

	return nil
}
