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
			image VARCHAR(200) NULL,
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
			title VARCHAR(100) NOT NULL,
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

	strPtr := func(s string) *string { return &s }

	exampleProjects := []models.CreateProject{
		models.NewCreateProject(
			"Khali",
			strPtr("https://raw.githubusercontent.com/projetoKhali/api3/refs/heads/main/docs/Banners/Api.png"),
			1, "FATEC", 1,
			"Avaliação 360",
			"A plataforma Khali permite a implementação do método de Avaliação 360° na Instituição de Ensino fictícia PBLTeX. Este projeto de API do 1º Semestre de Banco de Dados da Fatec - São José dos Campos possibilita uma abordagem abrangente na avaliação dos diversos aspectos da instituição, promovendo uma análise holística e aprimorando processos de gestão e desenvolvimento.",
			"github.com/taniacruzz/Khali",
		),
		models.NewCreateProject(
			"API2Semestre",
			strPtr("https://raw.githubusercontent.com/projetoKhali/API2Semestre/refs/heads/main/Docs/Banners/Novobanner.png"),
			2, "2RP", 1,
			"Controle de Horas-Extras e Sobreavisos (desktop)",
			"A API desenvolvida no 2° semestre do curso de Banco de Dados na Fatec - SJC proporciona um sistema desktop especializado no registro de horas extras e sobreavisos pelos colaboradores, com funcionalidades de controle tanto para gestores (PO) quanto para administradores (RH e Financeiro). Essa solução oferece uma plataforma integrada e eficiente para gerenciamento de tempo e recursos humanos, contribuindo para uma gestão mais eficaz e transparente dentro da organização.",
			"github.com/projetoKhali/API2Semestre",
		),
		models.NewCreateProject(
			"api3",
			strPtr("https://user-images.githubusercontent.com/111442399/194777358-24905c4f-e62b-414d-8754-b3ccaf878547.png"),
			3, "2RP", 1,
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
		models.NewCreateSkill("Docker"),
		models.NewCreateSkill("Bash"),
		models.NewCreateSkill("Batch"),
		models.NewCreateSkill("TypeScript"),
		models.NewCreateSkill("React"),
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
			[]string{"Python"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Inclusão das entidades iniciais",
			"Adicionei as entidades: User, Client, Appointment, PayRateRule, e os enums necessários. Essa contribuição foi composta por uma revisão do código JavFX do semestre anterior para garantir a compatibilidade com a arquitetura do terceiro semestre. Alguns pontos foram atualizados para seguir a nova arquitetura.",
			[]string{"Java", "Spring"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Dockerização do projeto", "Efetuei a dockerização do projeto, criando arquivos Dockerfile e docker-compose.yml para facilitar a execução do projeto em qualquer ambiente. Ao todo foram incluidos 3 containers no `docker-compose` principal do projeto: um para o banco de dados, um para o back-end e um para o front-end. Para o back-end e front-end, foram criados arquivos `Dockerfile` específicos para cada um, com as dependências necessárias para a execução do projeto.",
			[]string{"Docker", "Bash", "Batch"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Sistema de permissões de acesso às telas", "Implementei um sistema de permissões de acesso às telas do sistema, baseado em regras de acesso por perfil de usuário. As permissões de acesso são definidas pelo perfil do usuário e também por uma análise dos dados do usuário. Por exemplo, o usuário de nível Colaborador, só pode efetuar apontamentos caso pertença a um ResultCenter, caso contrário não faz sentido possuir acesso à tela de apontamentos.",
			[]string{"Java", "Spring", "TypeScript", "React"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Tela de Apontamentos", "Desenvolvi da tela de apontamentos do sistema, incluindo o formulário que permite ao usuário registrar as horas trabalhadas e os sobreavisos, assim como a tabela responsável pela listagem dos apontamentos efetuados previamente. A tela se comunica com o back-end através das funções service do front-end. ",
			[]string{"TypeScript", "React"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Backend e tela de Clientes", "Desenvolvi a tela de Clientes do sistema, incluindo o formulário que permite ao usuário cadastrar novos clientes e a tabela responsável pela listagem dos clientes cadastrados. A tela se comunica com o back-end através das funções service do front-end. ",
			[]string{"TypeScript", "React"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Componentes personalizado de Dropdown", "Desenvolvi componentes personalizados de Dropdown que integram com o funcionamento das telas para disponibilizar abstrações, facilitando a implementação de novas funcionalidades.",
			[]string{"TypeScript", "React"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Componente de LookUpTextField", "Desenvolvi o componente `LookUpTextField`, que permite a pesquisa em uma lista e reduzindo a quantidade de opções exibidas ao usuário ao selecionar um valor dentro de uma lista finita. A pesquisa é realizada em tempo real, filtrando os resultados conforme o usuário digita. O componente foi utilizado em diversas telas do sistema aonde a seleção de um valor existente entre muitos disponíveis é necessária.",
			[]string{"TypeScript", "React"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Componentes de Célula de Tabela Editável e Célula de Tabela Botão", "Desenvolvi o componente `EditableTableCell`, que permite a edição de células de uma tabela diretamente na célula, sem a necessidade de abrir um formulário de edição. O componente `ButtonTableCell` foi desenvolvido para permitir a inclusão de botões em células de uma tabela, facilitando a execução de ações específicas como exibir detalhes, editar ou excluir um registro.",
			[]string{"TypeScript", "React"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Fluxo de inclusão de usuários à ResultCenter", "Desenvolvi o fluxo de inclusão de usuários à ResultCenter, permitindo que um usuário seja vinculado a uma ResultCenter. Para isso, foi desenvolvido o `MemberController` no back-end, responsável pela definição dos endpoints utilizados durante a associação, desassociação e listagem de membros de uma ResultCenter. No front-end, durante a criação de um `ResultCenter` é possível utilizar a `LookUpTextField` para a pesquisa de usuários existentes para a associação. Os usuários podem ser incluídos ou excluídos de uma lista temporária de membros, que é persistida ao concluir o registro.",
			[]string{"Java", "Spring", "TypeScript", "React"},
		),
	}

	for _, n := range exampleContributions {
		if _, err := service.ContributionService.Create(n); err != nil {
			return tracerr.Errorf("failed to create contribution: %w", tracerr.Wrap(err))
		}
	}

	return nil
}
