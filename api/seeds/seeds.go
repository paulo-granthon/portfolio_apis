package seeds

import (
	"github.com/paulo-granthon/portfolio_apis/models"
	"github.com/paulo-granthon/portfolio_apis/service"
	"github.com/paulo-granthon/portfolio_apis/storage"
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
		DROP TABLE IF EXISTS participations CASCADE;

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

		CREATE TABLE IF NOT EXISTS participations (
			id SERIAL PRIMARY KEY,
			project_id INT NOT NULL,
			user_id INT NOT NULL,
			summary TEXT NOT NULL,
			UNIQUE (project_id, user_id),
			FOREIGN KEY (project_id) REFERENCES projects(id),
			FOREIGN KEY (user_id) REFERENCES users(id)
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

	if err := ParticipationMigrate(storage, *service); err != nil {
		return tracerr.Errorf("failed to run seeds: failed to migrate participation: %w", tracerr.Wrap(err))
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
			strPtr("https://user-images.githubusercontent.com/111442399/194777358-24905c4f-e62b-414d-8754-b3ccaf878547.png"),
			1, "FATEC", 1,
			"Plataforma desktop para aplicação do método de Avaliação 360°.",
			"Aplicação desktop em Python (Tkinter) que viabiliza a Avaliação 360° entre alunos e instrutores da instituição de ensino fictícia PBLTeX, voltada a cursos com metodologia PBL. Gerencia cadastro e autenticação de usuários, grupos, times e papéis, organiza Sprints com períodos avaliativos, coleta avaliações com feedback obrigatório para notas baixas e gera dashboards de desempenho individual, por time e por grupo, com persistência em arquivos CSV.",
			"github.com/projetoKhali/Khali",
		),
		models.NewCreateProject(
			"API2Semestre",
			strPtr("https://raw.githubusercontent.com/projetoKhali/API2Semestre/main/Docs/Banners/Novobanner.png"),
			2, "2RP Net", 1,
			"Sistema desktop em JavaFX para apontamento e controle de horas extras e sobreaviso.",
			"Aplicação desktop em Java/JavaFX desenvolvida para a 2RP Net. Permite que colaboradores registrem horas extras e sobreaviso vinculados a centros de resultado e projetos, enquanto gestores (PO) aprovam ou rejeitam os apontamentos e administradores parametrizam regras de fechamento, percentuais e adicional noturno, geram relatórios em CSV e acompanham um dashboard analítico. O acesso é baseado em papéis (colaborador, gestor e administrador) com persistência em banco relacional.",
			"github.com/projetoKhali/API2Semestre",
		),
		models.NewCreateProject(
			"api3",
			strPtr("https://raw.githubusercontent.com/projetoKhali/api3/main/docs/Banners/Api.png"),
			3, "2RP Net", 1,
			"Sistema web para controle de horas extras e sobreavisos, evolução web do aplicativo desktop anterior.",
			"Sistema web full stack (React + TypeScript no front-end, Java + Spring no back-end, PostgreSQL e Docker) para controle da jornada de trabalho, identificando e classificando horas extras e sobreavisos. Possui perfis de acesso distintos (administrador, gestor e colaborador), apontamento e aprovação de horas, cadastro de clientes, projetos e centros de resultado, parametrização de verbas, extração de relatórios em CSV e dashboards. Desenvolvido pela equipe Khali em parceria com a 2RP Net.",
			"github.com/projetoKhali/api3",
		),
		models.NewCreateProject(
			"api4",
			strPtr("https://github.com/projetoKhali/api4/assets/108769169/ecda074a-ef3f-4ca5-9cf0-d4b559bcbec5"),
			4, "Oracle", 1,
			"Plataforma de analytics para consultores de alianças acompanharem parceiros e produtos no ecossistema Oracle.",
			"Plataforma web de analytics desenvolvida para a Oracle. Oferece a consultores de aliança dashboards individuais e comparativos para monitorar a evolução de parceiros em trilhas, expertises e qualificações, além de métricas de desempenho de produtos por região. Inclui cadastro de parceiros e usuários, relatórios filtráveis com exportação em CSV e visualização de dados. Front-end em Vue 3 + TypeScript e back-end em Java/Spring com PostgreSQL, orquestrados via Docker.",
			"github.com/projetoKhali/api4",
		),
		models.NewCreateProject(
			"api5",
			strPtr("https://github.com/user-attachments/assets/94aecab2-e751-4ab4-a2a8-1a6589b4eb01"),
			5, "Pro4tech", 1,
			"Dashboard interativo de inteligência em recrutamento e seleção para decisões estratégicas de contratação.",
			"Plataforma de otimização de recrutamento e seleção desenvolvida para a Pro4tech. Centraliza e visualiza dados fragmentados dos processos seletivos em um dashboard interativo, com métricas em tempo real (tempo médio de contratação, status de vagas, processos em andamento), filtros personalizados, extração de relatórios e gestão de usuários com grupos de acesso. Arquitetura em três serviços: back-end em Go, front-end em TypeScript e microsserviço de previsão em Python.",
			"github.com/projetoKhali/api5",
		),
		models.NewCreateProject(
			"api6",
			strPtr("https://github.com/user-attachments/assets/d0217c10-db11-470b-a029-f8b664cf4cd2"),
			6, "Kersys", 1,
			"Sistema inteligente de planejamento e monitoramento de reflorestamento (SIPMR).",
			"Plataforma web (SIPMR) desenvolvida para a Kersys que monitora e gerencia plantios para otimizar a recuperação ambiental. Reúne dashboards de métricas e produtividade, gestão de eventos de plantio, simulador de cenários e previsões com machine learning, geração de relatórios e gestão de usuários com autenticação. Arquitetura em monorepo Nx com front-end React/TypeScript, API em Python/Flask, serviço de autenticação em Rust/Actix-web e bancos PostgreSQL e MongoDB.",
			"github.com/projetoKhali/api6",
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

	skillNames := []string{
		// Methodologies & tooling
		"Scrum", "Git", "GitHub Actions", "CI/CD", "Docker", "Docker Compose",
		"Makefile", "Bash", "Batch", "Nx", "Poetry", "Husky", "Swagger",
		"ETL", "Testes de Integração", "Testes Unitários", "Análise de Dados",
		"Modelagem de Dados",
		// Languages
		"Python", "Java", "TypeScript", "Go", "Rust", "SQL", "CSS",
		// Frameworks & libraries
		"Tkinter", "matplotlib", "JavaFX", "FXML", "Maven", "Spring", "React",
		"React Native", "Expo", "Vue", "Vue Router", "Flask", "Actix-web",
		"Ent ORM", "SeaORM", "Pydantic", "Jest", "ESLint", "Prettier",
		// Data
		"JDBC", "JPA", "PostgreSQL", "MongoDB",
		// Concepts
		"JWT", "REST", "Autenticação", "Controle de Acesso", "Criptografia",
		"Fernet", "Componentização", "Paginação", "UI/UX", "Refatoração", "DevOps",
	}

	for _, name := range skillNames {
		if _, err := skillModule.Create(models.NewCreateSkill(name)); err != nil {
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
		// ----------------------------------------------------------------
		// Semester 1 - Khali (Python / Tkinter)
		// ----------------------------------------------------------------
		models.NewCreateContributionByNames(
			"Khali", "paulo-granthon",
			"Back-end de cadastro, autenticação e criptografia de senhas",
			"Implementei o núcleo de back-end do sistema de cadastro e autenticação (módulo [`Authentication`](https://github.com/projetoKhali/Khali/blob/main/Authentication.py)), incluindo registro de usuários, validação de nomes e senhas e a relação entre Grupos, Times e Papéis (Roles). Desenvolvi o módulo [`Gerar_Senha`](https://github.com/projetoKhali/Khali/blob/main/Utils/Gerar_Senha.py), responsável pela criptografia de senhas e geração automática de credenciais, integrando-o ao fluxo de registro e ao sistema de envio de e-mail. Tratei as verificações de [`group_id`](https://github.com/projetoKhali/Khali/blob/main/Authentication.py) e [`role_id`](https://github.com/projetoKhali/Khali/blob/main/Authentication.py) no registro e estruturei a persistência por meio do [`CSVHandler`](https://github.com/projetoKhali/Khali/blob/main/CSV/CSVHandler.py).",
			[]string{"Python", "Autenticação", "Modelagem de Dados"},
		),
		models.NewCreateContributionByNames(
			"Khali", "paulo-granthon",
			"Camada de modelos de dados e persistência em CSV (CSVHandler)",
			"Construí e mantive a camada de acesso a dados do projeto, incluindo os Models (User, Group, Team, Sprint, Role) com verificações de nulidade e funções de conversão `to_<model>`. Desenvolvi o [`CSVHandler`](https://github.com/projetoKhali/Khali/blob/main/CSV/CSVHandler.py), responsável pela leitura e escrita em arquivos CSV, incluindo a função `get_data_list_by_fields_value_csv` para consultas filtradas por campos. Atualizei a estrutura de dados de Grupos e adicionei funções de acesso direto à base nos próprios modelos.",
			[]string{"Python", "Modelagem de Dados"},
		),
		models.NewCreateContributionByNames(
			"Khali", "paulo-granthon",
			"Módulo de Sprints e períodos avaliativos",
			"Implementei a lógica de Sprints que governa os prazos das avaliações, incluindo as funções [`current_sprint`](https://github.com/projetoKhali/Khali/blob/main/Models/Sprint.py#L41-L48), [`previous_sprint`](https://github.com/projetoKhali/Khali/blob/main/Models/Sprint.py#L57-L62), [`current_rating_period`](https://github.com/projetoKhali/Khali/blob/main/Models/Sprint.py#L49-L56), [`rating_period_start`](https://github.com/projetoKhali/Khali/blob/main/Models/Sprint.py#L21-L22) e [`rating_period_end`](https://github.com/projetoKhali/Khali/blob/main/Models/Sprint.py#L23-L25), que determinam a Sprint vigente e os períodos avaliativos ativos. Esses cálculos alimentam tanto a tela inicial (informações da Sprint ativa e prazo) quanto o controle de quando as avaliações podem ser enviadas.",
			[]string{"Python"},
		),
		models.NewCreateContributionByNames(
			"Khali", "paulo-granthon",
			"KML - Khali Markdown Language (protótipo de UI declarativa)",
			"Prototipei a [`Khali Markdown Language`](https://github.com/projetoKhali/Khali/tree/main/KML) (KML), uma linguagem de marcação própria, em estilo XML, para descrever telas do Tkinter de forma declarativa: criei o parser e o sistema de Tags ([`window`](https://github.com/projetoKhali/Khali/blob/main/KML/KML.py#L46-L49), [`module`](https://github.com/projetoKhali/Khali/blob/main/KML/KML.py#L50-L54), [`frame`](https://github.com/projetoKhali/Khali/blob/main/KML/KML.py#L55-L57), [`label`](https://github.com/projetoKhali/Khali/blob/main/KML/KML.py#L58-L60), [`entry`](https://github.com/projetoKhali/Khali/blob/main/KML/KML.py#L61-L63), [`button`](https://github.com/projetoKhali/Khali/blob/main/KML/KML.py#L64-L66), [`img`](https://github.com/projetoKhali/Khali/blob/main/KML/KML.py#L67-L70), [`loop`](https://github.com/projetoKhali/Khali/blob/main/KML/KML.py#L71-L77), [`id`](https://github.com/projetoKhali/Khali/blob/main/KML/KML.py)) que são interpretados recursivamente e convertidos nos widgets correspondentes. Cheguei a renderizar alguns componentes com sucesso e validei a abordagem comparando KML com Tkinter puro, mas a exploração não foi adotada na versão final do produto, nenhuma tela exposta ao usuário acabou construída com KML. Ainda assim, foi um exercício relevante de design de linguagem e parsing.",
			[]string{"Python", "Tkinter"},
		),
		models.NewCreateContributionByNames(
			"Khali", "paulo-granthon",
			"Sistema de Eventos e gerenciamento de janelas (WindowManager)",
			"Desenvolvi um sistema de Eventos (módulo [`Events`](https://github.com/projetoKhali/Khali/blob/main/Events.py)) com registro e desregistro de callbacks, usado para sincronizar estado entre componentes da interface (por exemplo, o armazenamento temporário de formulários ao alterar a quantidade de formulários exibidos). Implementei também o [`WindowManager`](https://github.com/projetoKhali/Khali/blob/main/Front/WindowManager.py) para o controle de navegação entre telas, integrando-o à autenticação e à home.",
			[]string{"Python", "Tkinter"},
		),
		models.NewCreateContributionByNames(
			"Khali", "paulo-granthon",
			"Dashboards de desempenho e cálculo de médias",
			"Fui o principal autor do módulo de Dashboards e do integrador de médias, implementando gráficos com matplotlib (barras horizontais, radar/pentágono do usuário e gráficos de pizza). Criei os cálculos de média por usuário, por papel ([`role_media`](https://github.com/projetoKhali/Khali/blob/main/graficos/Charts.py#L261-L296)), por time ([`users_media_team`](https://github.com/projetoKhali/Khali/blob/main/graficos/Charts.py#L297-L323), [`team_media_sprints`](https://github.com/projetoKhali/Khali/blob/main/graficos/Charts.py#L180-L203)) e por grupo ([`group_media_sprints`](https://github.com/projetoKhali/Khali/blob/main/graficos/Charts.py#L324-L347)), além de visões comparativas como [`user_media_x_team`](https://github.com/projetoKhali/Khali/blob/main/graficos/Charts.py#L151-L179), [`team_media_x_group`](https://github.com/projetoKhali/Khali/blob/main/graficos/Charts.py#L227-L260) e [`group_media_x_groups`](https://github.com/projetoKhali/Khali/blob/main/graficos/Charts.py#L348-L384), com tratamento de matrizes vazias e dados nulos.",
			[]string{"Python", "matplotlib", "Análise de Dados"},
		),
		models.NewCreateContributionByNames(
			"Khali", "paulo-granthon",
			"Telas de cadastro, avaliação e edição de times",
			"Implementei e refatorei diversas telas em Tkinter: cadastro do Administrador (com edição e exclusão de grupos), cadastro pelo Líder do Grupo (com validações e popup de sucesso), a tela de avaliação (com desativação de campos no envio e integração ao sistema de Eventos), a home (com seletor de módulos) e a edição de times (com inclusão e remoção de usuários, edição e exclusão). Desenvolvi também componentes reutilizáveis como Scrollbar e funções gerais de Tkinter no [`Front.Core`](https://github.com/projetoKhali/Khali/blob/main/Front/Core.py).",
			[]string{"Python", "Tkinter", "UI/UX"},
		),
		models.NewCreateContributionByNames(
			"Khali", "paulo-granthon",
			"Integração geral, refatorações e correções de bugs",
			"Atuei de forma transversal integrando back-end e front-end (login, home e avaliação) e fui responsável por grande volume de refatorações e correções de bugs ao longo das quatro Sprints, incluindo merges de múltiplas branches da equipe, correções no [`CSVHandler`](https://github.com/projetoKhali/Khali/blob/main/CSV/CSVHandler.py), em referências de papéis, em seletores de Sprint do perfil, em dropdowns e em casos de dados nulos nos dashboards. Mantive o README e abri o PR de fechamento, consolidando a branch principal e a documentação.",
			[]string{"Python", "Git"},
		),

		models.NewCreateContributionByNames(
			"Khali", "paulo-granthon",
			"Front.Core - design system e helpers de UI em Tkinter",
			"Construí o módulo [`Front/Core.py`](https://github.com/projetoKhali/Khali/blob/main/Front/Core.py), a fundação visual e funcional de toda a interface. Centralizei a paleta de cores do projeto e expus funções genéricas que abstraem a verbosidade do Tkinter ([`criar_frame`](https://github.com/projetoKhali/Khali/blob/main/Front/Core.py#L9-L14), [`criar_label`](https://github.com/projetoKhali/Khali/blob/main/Front/Core.py#L15-L20), [`criar_button`](https://github.com/projetoKhali/Khali/blob/main/Front/Core.py#L21-L26), [`criar_entry`](https://github.com/projetoKhali/Khali/blob/main/Front/Core.py#L27-L33)), além de helpers avançados como [`bind_entry_placeholder`](https://github.com/projetoKhali/Khali/blob/main/Front/Core.py#L34-L50) (texto-fantasma que some ao focar o campo), [`create_dropdown`](https://github.com/projetoKhali/Khali/blob/main/Front/Core.py#L51-L61) (menus integrados ao sistema de eventos para atualizações reativas) e [`bind_edit_label`](https://github.com/projetoKhali/Khali/blob/main/Front/Core.py#L62-L95) (transforma labels em campos editáveis ao clicar). Isso deu consistência de estilo e layout a todas as telas e eliminou boilerplate de interface repetido em toda a base de código.",
			[]string{"Python", "Tkinter", "Componentização", "UI/UX", "Refatoração"},
		),
		models.NewCreateContributionByNames(
			"Khali", "paulo-granthon",
			"Componente de área rolável (Scrollbar) com Canvas",
			"Implementei o [`Front/Scrollbar.py`](https://github.com/projetoKhali/Khali/blob/main/Front/Scrollbar.py), um container rolável montado sobre a combinação Canvas + Frame + Scrollbar do Tkinter. Inseri o conteúdo no Canvas via [`create_window`](https://github.com/projetoKhali/Khali/blob/main/Front/Scrollbar.py) e usei o evento [`<Configure>`](https://github.com/projetoKhali/Khali/blob/main/Front/Scrollbar.py) para ajustar dinamicamente largura/altura e recalcular a [`scrollregion`](https://github.com/projetoKhali/Khali/blob/main/Front/Scrollbar.py) conforme o conteúdo muda. Sincronizei a barra de rolagem (`yview`/`yscrollcommand`) e adicionei suporte à roda do mouse com binding global ativado ao entrar no frame e removido ao sair, integrado ao sistema de eventos para desabilitar a rolagem quando submódulos abrem.",
			[]string{"Python", "Tkinter", "Componentização", "UI/UX"},
		),
		models.NewCreateContributionByNames(
			"Khali", "paulo-granthon",
			"ModulesManager e seletor de módulos da Home",
			"Desenvolvi o sistema de módulos carregados dentro da Home ([`ModulesManager`](https://github.com/projetoKhali/Khali/blob/main/Front/Modules/ModulesManager.py) + [`home_front`](https://github.com/projetoKhali/Khali/blob/main/Front/Windows/home_front.py)), incluindo o seletor de módulos estilizado, a verificação de permissões por papel ([`role_id`](https://github.com/projetoKhali/Khali/blob/main/Front/Modules/ModulesManager.py)) e o tratamento de casos de grupo nulo. Esse sistema permite que diferentes telas (lista de usuários, avaliação, dashboards, edição de times) sejam plugadas como módulos dentro da janela principal de forma condicional ao perfil do usuário.",
			[]string{"Python", "Tkinter", "Controle de Acesso", "Componentização"},
		),

		// ----------------------------------------------------------------
		// Semester 2 - API2Semestre (Java / JavaFX)
		// ----------------------------------------------------------------
		models.NewCreateContributionByNames(
			"API2Semestre", "paulo-granthon",
			"Camada de persistência e abstração de queries SQL (Query.java)",
			"Refatorei a camada de acesso a dados do sistema introduzindo um construtor de queries (PR [#87](https://github.com/projetoKhali/API2Semestre/pull/87)). Criei a classe [`Query.java`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/database/query/Query.java#L8) junto de [`QueryParam`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/database/query/QueryParam.java) e os enums [`QueryType`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/database/query/QueryType.java), [`QueryTable`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/database/query/QueryTable.java) e [`TableProperty`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/database/query/TableProperty.java), eliminando a escrita de strings SQL cruas dentro de [`QueryLibs`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/database/QueryLibs.java). Com isso, as consultas passaram a ser montadas de forma parametrizada e tipada, reduzindo erros e duplicação de código. Também atuei nos métodos de conexão com o banco ([`SQLConnection`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/database/SQLConnection.java)), no reaproveitamento de conexões e na remoção de parâmetros redundantes.",
			[]string{"Java", "JDBC", "SQL", "Modelagem de Dados"},
		),
		models.NewCreateContributionByNames(
			"API2Semestre", "paulo-granthon",
			"Autenticação, sessão e fluxo de aprovação de apontamentos",
			"Implementei o mecanismo de autenticação e o gerenciamento de sessão do sistema (PR [#71](https://github.com/projetoKhali/API2Semestre/pull/71)), incluindo a tela de login com identidade visual e o protótipo do [`ViewManager`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/view/manager/ViewsManager.java) para o controle de navegação entre telas conforme o usuário logado. Desenvolvi também a lógica de aprovação e rejeição de apontamentos pelo gestor (PO), com as condições de rejeição, e complementei o trabalho de permissões e configuração de visualização por papel (`ViewConfig`/`Permissions`).",
			[]string{"Java", "JavaFX", "Autenticação", "Controle de Acesso"},
		),
		models.NewCreateContributionByNames(
			"API2Semestre", "paulo-granthon",
			"Tela de Relatórios com exportação em CSV",
			"Desenvolvi a tela de extração de relatórios (PR [#102](https://github.com/projetoKhali/API2Semestre/pull/102)), com tabela de pré-visualização dos dados e exportação em CSV. Implementei o [`ReportController`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/view/controllers/views/Report.java) e o modelo [`Report.java`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/view/controllers/views/Report.java) usando `PropertyValueFactory` para popular a tabela, além de checkboxes de seleção de colunas que alternam os campos exibidos na prévia. Realizei ajustes posteriores na formatação de datas em [`ReportIntervalWrapper`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/view/utils/wrappers/ReportIntervalWrapper.java) e na soma de [`TotalHours`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/view/controllers/templates/TotalHours.java).",
			[]string{"Java", "JavaFX"},
		),
		models.NewCreateContributionByNames(
			"API2Semestre", "paulo-granthon",
			"Dashboard analítico com gráfico de Volume por Hora do Dia",
			"Criei a estrutura base da tela de Dashboard (PR [#124](https://github.com/projetoKhali/API2Semestre/pull/124)), integrando-a ao gerenciador de views, e desenvolvi a classe utilitária [`ChartGenerator.java`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/view/utils/ChartGenerator.java#L36) para gerar gráficos dinâmicos. Implementei o gráfico de Volume por Hora do Dia com lógica de interseção de intervalos de apontamento em janelas de 24 horas, avaliando três estratégias de comparação e escolhendo a que gerava menor distorção visual. Adicionei também filtros ao Dashboard e ajustes de estilo.",
			[]string{"Java", "JavaFX", "Análise de Dados"},
		),
		models.NewCreateContributionByNames(
			"API2Semestre", "paulo-granthon",
			"Parametrização de regras de horas e cadastro de Centro de Resultado",
			"Implementei a tela e a lógica de parametrização (PR [#104](https://github.com/projetoKhali/API2Semestre/pull/104)), com o controller [`Parametrization.java`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/view/controllers/views/Parametrization.java#L38), o wrapper [`IntervalFeeWrapper`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/view/utils/wrappers/IntervalFeeWrapper.java) e ajustes em [`Expedient.java`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/data/Expedient.java#L13) (dia de fechamento) e [`IntervalFee`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/report/IntervalFee.java) (tipo e duração do apontamento), incluindo a criação da tabela de parametrização e a validação de formato de horário. Também entreguei o cadastro de Centro de Resultado e a edição de verbas.",
			[]string{"Java", "JavaFX", "Modelagem de Dados"},
		),
		models.NewCreateContributionByNames(
			"API2Semestre", "paulo-granthon",
			"Componente genérico reutilizável de busca e seleção (LookupTextField)",
			"Desenvolvi o componente genérico [`LookupTextField`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/view/controllers/custom_tags/LookupTextField.java#L27) (PR [#115](https://github.com/projetoKhali/API2Semestre/pull/115)), uma caixa de pesquisa e seleção reutilizável para padronizar buscas em diferentes telas do sistema, em vez de múltiplas soluções customizadas. Também implementei a tela de Listagem e contribuí na reestruturação do sistema de arquivos do projeto e na atualização do diagrama de entidades.",
			[]string{"Java", "JavaFX", "Componentização", "UI/UX"},
		),

		models.NewCreateContributionByNames(
			"API2Semestre", "paulo-granthon",
			"Reestruturação para o padrão Maven e configuração do módulo JavaFX",
			"Liderei a reorganização completa do sistema de arquivos do projeto (PRs [#52](https://github.com/projetoKhali/API2Semestre/pull/52) e [#84](https://github.com/projetoKhali/API2Semestre/pull/84)), migrando para o layout convencional do Maven com [`pom.xml`](https://github.com/projetoKhali/API2Semestre/blob/main/pom.xml) e [`src/main/java/module-info.java`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/module-info.java) para configurar dependências externas e a compatibilidade com o JavaFX. Realoquei todo o código Java para [`src/main/java/org/openjfx/API2Semestre/`](https://github.com/projetoKhali/API2Semestre/tree/main/src/main/java/org/openjfx/API2Semestre) e os FXML para [`src/main/resources/`](https://github.com/projetoKhali/API2Semestre/tree/main/src/main/resources), removi arquivos legados e limpei artefatos de build no [`.gitignore`](https://github.com/projetoKhali/API2Semestre/blob/main/.gitignore). Essa base garantiu compatibilidade com o NetBeans e um build padronizado para todo o time.",
			[]string{"Maven", "JavaFX", "Java", "Refatoração"},
		),
		models.NewCreateContributionByNames(
			"API2Semestre", "paulo-granthon",
			"Sistema de configuração de telas via tags FXML customizadas",
			"No PR [#62](https://github.com/projetoKhali/API2Semestre/pull/62) (alto impacto), implementei tags FXML customizadas e específicas do projeto para o gerenciamento da configuração de telas (ViewConfig), criando uma abordagem padronizada para a montagem da interface, e acoplei a ela o tratamento de permissões de visualização por papel. Conduzi também a refatoração de pastas seguindo convenções Java. O trabalho foi reconhecido pelo time como de alta qualidade e serviu de base para as demais telas.",
			[]string{"JavaFX", "FXML", "Java", "Controle de Acesso", "Refatoração"},
		),
		models.NewCreateContributionByNames(
			"API2Semestre", "paulo-granthon",
			"Macro reutilizável de células editáveis em tabelas e edição de Verbas",
			"Criei [`view.macros.TableMacros.enableEditableCells`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/view/macros/TableMacros.java#L223-L291), uma macro que torna editável qualquer coluna de qualquer tabela, recebendo parâmetros para validação, atualização de itens e exibição (PR [#131](https://github.com/projetoKhali/API2Semestre/pull/131)). Sobre essa abstração, entreguei a edição de Verbas na tela de Parametrização, e a mesma macro passou a sustentar a edição nas telas de Usuários, Centros de Resultado e Clientes, garantindo consistência e reduzindo duplicação em todo o sistema.",
			[]string{"JavaFX", "Java", "Componentização"},
		),
		models.NewCreateContributionByNames(
			"API2Semestre", "paulo-granthon",
			"Tela de Listagem e Aprovação com filtros dinâmicos por coluna",
			"No PR [#67](https://github.com/projetoKhali/API2Semestre/pull/67) (alto impacto, resolvendo seis issues), construí a listagem dinâmica de apontamentos em tabela, com seleção de linhas e filtragem textual por coluna via macros ([`AppointmentFilter`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/view/utils/filters/AppointmentFilter.java), [`AppointmentWrapper`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/view/utils/wrappers/AppointmentWrapper.java)), e a tela de Aprovação com checkboxes para aprovar/rejeitar apontamentos e popups de feedback. Refatorei o `buildTable`, corrigi o `collaboratorSelect` na criação de apontamentos e integrei tudo à camada de persistência, atravessando os controllers e os FXML.",
			[]string{"JavaFX", "Java", "Componentização", "Controle de Acesso"},
		),
		models.NewCreateContributionByNames(
			"API2Semestre", "paulo-granthon",
			"Modelagem inicial do domínio de apontamentos",
			"No PR [#25](https://github.com/projetoKhali/API2Semestre/pull/25) entreguei as classes de modelo fundamentais do domínio: [`Appointment.java`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/appointment/Appointment.java#L9), que armazena as informações dos registros de apontamento, e o enum [`AppointmentType`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/appointment/AppointmentType.java#L5-L42), que distingue Hora-Extra e Sobreaviso. Essa estrutura serviu de base para toda a gestão de apontamentos do sistema de controle de horas.",
			[]string{"Java", "Modelagem de Dados"},
		),
		models.NewCreateContributionByNames(
			"API2Semestre", "paulo-granthon",
			"Otimização do acesso ao banco com reuso de conexões",
			"No PR [#146](https://github.com/projetoKhali/API2Semestre/pull/146) (consolidando 25 commits) conduzi melhorias de performance no acesso ao banco, adicionando reuso de conexão em [`Appointments.java`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/view/controllers/views/Appointments.java) e otimizações em [`QueryLibs`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/database/QueryLibs.java) e [`SQLConnector`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/database/SQLConnection.java). Junto disso, corrigi o cálculo de soma de [`TotalHours`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/view/controllers/templates/TotalHours.java), acessores do Dashboard e Expedient, a formatação de relatórios e fiz o debugging de diversas telas, além de limpar imports e TODOs obsoletos.",
			[]string{"Java", "JDBC", "SQL", "Refatoração"},
		),

		// ----------------------------------------------------------------
		// Semester 3 - api3 (React + TypeScript / Java + Spring)
		// ----------------------------------------------------------------
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Inclusão das entidades iniciais do back-end",
			"Modelei e implementei as entidades iniciais do domínio em Java, estabelecendo a base do back-end. Criei as classes [`User`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/user/User.java#L39-L54) e [`UserType`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/user/UserType.java), [`Appointment`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/appointment/Appointment.java#L49-L99) com os enums [`AppointmentStatus`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/appointment/AppointmentStatus.java) e [`AppointmentType`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/appointment/AppointmentType.java), [`Client`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/client/Client.java#L28-L43) e [`PayRateRule`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/pay_rate_rule/PayRateRule.java) (regras de verba), além de restrições de unicidade na entidade [`User`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/user/User.java#L39-L54). Esse conjunto de entidades serviu de fundação para os repositórios, serviços e controllers desenvolvidos ao longo das sprints.",
			[]string{"Java", "Spring", "Modelagem de Dados", "PostgreSQL"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Dockerização do projeto",
			"Estruturei todo o ambiente de execução conteinerizado do projeto. Criei os [`Dockerfile`](https://github.com/projetoKhali/api3/blob/main/api/Dockerfile) do back-end e do front-end, o [`docker-compose.yml`](https://github.com/projetoKhali/api3/blob/main/docker-compose.yml) orquestrando os containers de API, front-end e banco de dados PostgreSQL (com política de auto restart da API), e configurei o [`application.properties`](https://github.com/projetoKhali/api3/blob/main/api/src/main/resources/application.properties) para o ambiente Docker. Adicionei ainda scripts de execução para Unix e Windows, um [Makefile](https://github.com/projetoKhali/api3/blob/main/Makefile), o script [`pgconnect.sh`](https://github.com/projetoKhali/api3/blob/main/pgconnect.sh) e arquivos [`.env.schema`](https://github.com/projetoKhali/api3/blob/main/.env.schema), padronizando a inicialização do ambiente para toda a equipe.",
			[]string{"Docker", "Docker Compose", "Makefile", "Bash", "Batch", "PostgreSQL"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Sistema de permissões de acesso às telas",
			"Implementei o sistema de controle de acesso baseado em perfis (administrador, gestor e colaborador), abrangendo back-end e front-end. No back-end criei a entidade [`Permission`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/permission/Permission.java#L9-L23) e a rota [`/{id}/permissions`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/controllers/UserController.java#L106) no [`UserController`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/controllers/UserController.java#L41-L173). No front-end desenvolvi o serviço [`Access`](https://github.com/projetoKhali/api3/blob/main/web/src/services/Access.ts), com a função [`getUserSideMenuItems`](https://github.com/projetoKhali/api3/blob/main/web/src/services/Access.ts#L97-L98) que gera dinamicamente os itens do menu lateral conforme as permissões. Configurei também o [`CorsConfig`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/config/CorsConfig.java#L8-L16) com os [`allowedOrigins`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/config/CorsConfig.java#L8-L16) corretos para viabilizar a comunicação entre front e back.",
			[]string{"Java", "Spring", "React", "TypeScript", "Controle de Acesso"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Tela de Apontamentos",
			"Desenvolvi a tela de apontamentos de horas extras e sobreaviso no front-end React/TypeScript. Criei o schema [`Appointment`](https://github.com/projetoKhali/api3/blob/main/web/src/schemas/Appointment.ts), o [`AppointmentService`](https://github.com/projetoKhali/api3/blob/main/web/src/services/AppointmentService.ts) para consumo da API, o componente [`AppointmentForm`](https://github.com/projetoKhali/api3/blob/main/web/src/components/AppointmentForm.tsx) e a página `Appointments`, integrando-os à rota [`/appointments`](https://github.com/projetoKhali/api3/blob/main/web/src/App.tsx). Implementei validação no [`handleSubmit`](https://github.com/projetoKhali/api3/blob/main/web/src/components/AppointmentForm.tsx), tratamento de erros via [`errorCallback`](https://github.com/projetoKhali/api3/blob/main/web/src/components/AppointmentForm.tsx), conversão de campos de texto para dropdowns e correções de bugs como recarregamento de página em submit inválido e loop contínuo de `useEffect`.",
			[]string{"React", "TypeScript", "REST"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Back-end e tela de Clientes",
			"Implementei a funcionalidade completa de cadastro de clientes, de ponta a ponta. No back-end criei o [`ClientController`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/controllers/ClientController.java) com seus endpoints REST. No front-end desenvolvi o [`ClientService`](https://github.com/projetoKhali/api3/blob/main/web/src/services/ClientService.ts), o schema [`Client`](https://github.com/projetoKhali/api3/blob/main/web/src/schemas/Client.ts), o componente [`ClientForm`](https://github.com/projetoKhali/api3/blob/main/web/src/components/ClientForm.tsx) e a página [`Clients`](https://github.com/projetoKhali/api3/blob/main/web/src/pages/Clients.tsx), mapeada à rota [`/clients`](https://github.com/projetoKhali/api3/blob/main/web/src/App.tsx). Também ajustei campos das entidades, como a inclusão de [`insertDate`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/client/Client.java#L39) e [`expireDate`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/client/Client.java#L42) no [`Client`](https://github.com/projetoKhali/api3/blob/main/web/src/schemas/Client.ts).",
			[]string{"Java", "Spring", "React", "TypeScript", "REST"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Componente personalizado de Dropdown",
			"Desenvolvi um componente [`Dropdown`](https://github.com/projetoKhali/api3/blob/main/web/src/components/Dropdown.tsx#L11-L66) reutilizável em React/TypeScript, usado em diversos formulários do sistema, como o de apontamentos. Defini a tipagem das [`DropdownProps`](https://github.com/projetoKhali/api3/blob/main/web/src/components/Dropdown.tsx#L5) (incluindo placeholder), implementei o comportamento de fechar o dropdown após a seleção, adicionei o ícone caret-down e apliquei estilização CSS dedicada. O componente substituiu campos de texto por seleção controlada em vários formulários.",
			[]string{"React", "TypeScript", "Componentização", "CSS", "UI/UX"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Componente de LookUpTextField",
			"Criei o componente [`LookUpTextField`](https://github.com/projetoKhali/api3/blob/main/web/src/components/LookUpTextField.tsx#L11-L51) em React/TypeScript, um campo de texto com funcionalidade de busca/lookup para seleção assistida de valores em formulários. Desenvolvi o componente e apliquei estilização CSS própria, integrando-o aos formulários do sistema em que era necessário selecionar um valor existente entre muitos disponíveis.",
			[]string{"React", "TypeScript", "Componentização", "CSS", "UI/UX"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Componentes de Célula de Tabela Editável e de Botão",
			"Desenvolvi componentes reutilizáveis de célula de tabela para o front-end. Criei o [`EditableTableCell`](https://github.com/projetoKhali/api3/blob/main/web/src/components/EditableTableCell.tsx), célula que permite edição inline de valores (populando o input com o valor atual ao iniciar a edição), e o [`ButtonTableCell`](https://github.com/projetoKhali/api3/blob/main/web/src/components/ButtonTableCell.tsx), célula que renderiza um botão de ação. Realizei refatorações de nomenclatura, padronizando os componentes como `EditableTableColumn` e `ButtonTableColumn`, e os integrei às páginas do sistema.",
			[]string{"React", "TypeScript", "Componentização", "UI/UX"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Fluxo de inclusão de usuários à ResultCenter",
			"Implementei o fluxo de associação de usuários (membros) aos Centros de Resultado ([`ResultCenter`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/resultCenter/ResultCenter.java#L43-L50)), abrangendo back-end e front-end. No back-end criei o [`MemberController`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/controllers/MemberController.java), a função [`findByUserType`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/repositories/UserRepository.java) no [`UserRepository`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/repositories/UserRepository.java) e [`getResultCentersOfUser`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/controllers/ResultCenterController.java) no [`ResultCenterController`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/controllers/ResultCenterController.java). No front-end desenvolvi o [`MemberService`](https://github.com/projetoKhali/api3/blob/main/web/src/services/MemberService.ts), o [`MemberSchema`](https://github.com/projetoKhali/api3/blob/main/web/src/schemas/Member.ts) e o componente [`SchemaList`](https://github.com/projetoKhali/api3/blob/main/web/src/components/SchemaList.tsx#L8-L21), permitindo buscar usuários por tipo e vinculá-los ao centro de resultado durante sua criação.",
			[]string{"Java", "Spring", "React", "TypeScript", "REST"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Motor de cálculo de fatias para relatórios de horas",
			"Desenvolvi o motor que segmenta os apontamentos em fatias (slices) para a extração de relatórios e a aplicação das regras de verba. No back-end criei as classes [`Slice`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/slice/Slice.java#L23-L26), [`SliceCalculator`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/util/SliceCalculator.java#L24-L267), [`SliceService`](https://github.com/projetoKhali/api3/blob/main/web/src/services/SliceService.ts) e [`SliceController`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/controllers/SliceController.java), a utilitária [`Pair`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/util/Pair.java) e o conceito de [`Week`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/util/Week.java), além de funções como [`getShiftTimeRange`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/services/PayRateRuleService.java) e [`IntegratedPayRateRule`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/pay_rate_rule/IntegratedPayRateRule.java) no [`PayRateRuleService`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/services/PayRateRuleService.java). No front-end criei o [`SliceSchema`](https://github.com/projetoKhali/api3/blob/main/web/src/schemas/Slice.ts) e o [`SliceService`](https://github.com/projetoKhali/api3/blob/main/web/src/services/SliceService.ts). Esse fluxo viabilizou o cálculo correto de horas extras e sobreaviso considerando turnos e percentuais.",
			[]string{"Java", "Spring", "React", "TypeScript"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Parametrização de regras de verba (PayRateRule)",
			"Implementei a tela e a integração de parametrização das regras de verba. No front-end desenvolvi funções no [`PayRateRulesService`](https://github.com/projetoKhali/api3/blob/main/web/src/services/PayRateRulesService.ts) ([`postPayRateRules`](https://github.com/projetoKhali/api3/blob/main/web/src/services/PayRateRulesService.ts), [`validatePayRateRules`](https://github.com/projetoKhali/api3/blob/main/web/src/services/PayRateRulesService.ts)), inicializei a página [`Parametrization`](https://github.com/projetoKhali/api3/blob/main/web/src/pages/Parametrization.tsx#L24-L198) com os parâmetros carregados e adicionei o campo `daysOfWeek` ao schema. No back-end criei o conversor [`DaysOfWeekConverter`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/converters/days_of_week/DaysOfWeekConverter.java#L8-L29), a entidade `Expedient`, a configuração [`JpaConfig`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/config/JpaConfig.java#L8-L15) e a função `findDefault` no repositório.",
			[]string{"Java", "Spring", "JPA", "React", "TypeScript"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Cadastro de usuários e autenticação (Login)",
			"Implementei o cadastro de usuários e o fluxo de login. No back-end criei o [`UserController`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/controllers/UserController.java#L41-L173) com seus endpoints, movi [`getUsers`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/services/UserService.java) para o [`UserService`](https://github.com/projetoKhali/api3/blob/main/web/src/services/UserService.ts) e configurei a documentação da API com SpringFox/Swagger. No front-end criei a página [`Login`](https://github.com/projetoKhali/api3/blob/main/web/src/pages/Login.tsx#L13-L74), adicionei a função [`requestLogin`](https://github.com/projetoKhali/api3/blob/main/web/src/services/UserService.ts) ao [`UserService`](https://github.com/projetoKhali/api3/blob/main/web/src/services/UserService.ts), a interface [`UserData`](https://github.com/projetoKhali/api3/blob/main/web/src/services/UserService.ts) e ajustes na página de usuários, incluindo ocultar a senha no input.",
			[]string{"Java", "Spring", "Swagger", "React", "TypeScript", "Autenticação"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Padronização visual e estilização do sistema",
			"Liderei a frente de padronização visual do front-end. Padronizei a paleta de cores, organizei os arquivos CSS em um diretório [`/styles`](https://github.com/projetoKhali/api3/tree/main/web/src/styles) dedicado, criei estilos para páginas e componentes (login, popup, menu lateral colapsável, formulários genéricos, `AppointmentForm`, `Dropdown` e `LookUpTextField`), integrei a biblioteca de ícones boxicons, adicionei o logotipo Khali à tela de login e ajustei espaçamentos e tamanhos dos componentes.",
			[]string{"CSS", "React", "TypeScript", "UI/UX"},
		),

		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Documentação interativa da API com Swagger (SpringFox)",
			"Integrei o SpringFox ao back-end Spring Boot para gerar documentação interativa da API, adicionando as dependências [`springfox-boot-starter`](https://github.com/projetoKhali/api3/blob/main/api/pom.xml)/[`springfox-swagger2`](https://github.com/projetoKhali/api3/blob/main/api/pom.xml) ao [`pom.xml`](https://github.com/projetoKhali/api3/blob/main/api/pom.xml) e criando a classe `SpringFoxConfig` (PR [#68](https://github.com/projetoKhali/api3/pull/68)). Isso disponibilizou uma Swagger UI navegável documentando automaticamente todos os endpoints REST (apontamentos, usuários, clientes, centros de resultado, parametrização), facilitando o consumo pelo front-end e os testes manuais.",
			[]string{"Java", "Spring", "Swagger", "REST"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Componente genérico de lista editável de itens (SchemaList)",
			"Criei o componente reutilizável [`SchemaList`](https://github.com/projetoKhali/api3/blob/main/web/src/components/SchemaList.tsx#L8-L21) em React+TypeScript, usado para gerenciar coleções dinâmicas de itens dentro de formulários (por exemplo, a lista de membros ao criar um Centro de Resultado). O componente abstrai a adição, exibição e remoção de itens tipados por schema, padronizando o tratamento de listas dinâmicas e servindo de base para fluxos de cadastro com múltiplas entradas.",
			[]string{"React", "TypeScript", "Componentização", "UI/UX"},
		),
		models.NewCreateContributionByNames(
			"api3", "paulo-granthon",
			"Back-end de Apontamentos: serviço, repositório e filtros",
			"Desenvolvi a camada de back-end dos apontamentos: o [`AppointmentService`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/services/AppointmentService.java) (insert, get, getById, update) e o [`AppointmentRepository`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/repositories/AppointmentRepository.java) com queries nativas para filtrar por usuário, data e intervalo de tempo, além de [`updateStatusAppointment`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/repositories/AppointmentRepository.java) e [`findByActive`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/repositories/AppointmentRepository.java)/[`findByInactive`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/repositories/AppointmentRepository.java). Tratei detalhes de mapeamento ORM ([`insertDate`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/appointment/Appointment.java#L49-L99)/[`expireDate`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/appointment/Appointment.java#L49-L99), [`@Transient`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/appointment/Appointment.java#L49-L99) no campo [`status`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/appointment/Appointment.java#L49-L99), função [`copy`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/domain/appointment/Appointment.java#L49-L99)) e a validação no [`AppointmentController`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/controllers/AppointmentController.java).",
			[]string{"Java", "Spring", "JPA", "SQL", "REST"},
		),

		// ----------------------------------------------------------------
		// Semester 4 - api4 (Vue / Java + Spring) - Paulo: Scrum Master
		// ----------------------------------------------------------------
		models.NewCreateContributionByNames(
			"api4", "paulo-granthon",
			"Scrum Master e arquitetura do monorepo e da infraestrutura Docker/PostgreSQL",
			"Como Scrum Master da equipe Khali, liderei a organização técnica do projeto, estruturado em três repositórios (umbrella [`api4`](https://github.com/projetoKhali/api4), [`api4back`](https://github.com/projetoKhali/api4back) e [`api4front`](https://github.com/projetoKhali/api4front) via submódulos). Configurei a execução integrada do monorepo com o pacote [`concurrently`](https://github.com/projetoKhali/api4/blob/main/package.json#L13), criei os templates de pull request e de issue e padronizei o fluxo de contribuição. No back-end, montei toda a infraestrutura de banco de dados em container, criando o [`docker-compose`](https://github.com/projetoKhali/api4back/blob/main/docker-compose.yml) do PostgreSQL e configurando o [`application.properties`](https://github.com/projetoKhali/api4back/blob/main/src/main/resources/application.properties) com as variáveis de conexão, além de ajustar o [`pom.xml`](https://github.com/projetoKhali/api4back/blob/main/pom.xml) do projeto Spring.",
			[]string{"Scrum", "Git", "Docker", "PostgreSQL", "Java", "Spring", "Vue"},
		),
		models.NewCreateContributionByNames(
			"api4", "paulo-granthon",
			"Modelagem e seeds do banco de dados PostgreSQL",
			"Trabalhei diretamente na camada de banco de dados do back-end, escrevendo e corrigindo os scripts DDL ([`postgres_version`](https://github.com/projetoKhali/api4back/blob/main/db/1_postgres_version.sql)), de seeds ([`postgres_seeds`](https://github.com/projetoKhali/api4back/blob/main/db/5_postgres_seeds.sql)) e de limpeza ([`postgres_drop_all`](https://github.com/projetoKhali/api4back/blob/main/db/0_postgres_drop_all.sql)). Adicionei a constraint UNIQUE em [`usr_login`](https://github.com/projetoKhali/api4back/blob/main/db/1_postgres_version.sql#L117) na tabela Users, corrigi chaves primárias da tabela Expertise, ajustei colunas da tabela Track e apliquei constraints de unicidade. Criei seeds com dados oficiais para Track, Expertise e Users, usando [`ON CONFLICT DO NOTHING`](https://github.com/projetoKhali/api4back/blob/main/db/5_postgres_seeds.sql) para evitar erros na carga.",
			[]string{"PostgreSQL", "SQL", "Modelagem de Dados"},
		),
		models.NewCreateContributionByNames(
			"api4", "paulo-granthon",
			"Endpoints, paginação e validação no back-end Java/Spring",
			"No back-end Spring, implementei paginação em endpoints de relatórios e métricas ([`allPartnerReports`](https://github.com/projetoKhali/api4back/blob/main/src/main/java/com/fatec/springapi4/controller/PartnerReportsController.java#L38-L41), [`allPartnerMetrics`](https://github.com/projetoKhali/api4back/blob/main/src/main/java/com/fatec/springapi4/controller/PartnerMetricsController.java#L24-L26)), inclusive adicionando uma coluna [`id`](https://github.com/projetoKhali/api4back/blob/main/db/3_postgres_views.sql#L150) dedicada via [`row_number`](https://github.com/projetoKhali/api4back/blob/main/db/3_postgres_views.sql#L150) para ordenação estável. Criei a classe [`Validation`](https://github.com/projetoKhali/api4back/blob/main/src/main/java/com/fatec/springapi4/Validation.java#L8-L64) com as funções [`validatePartner`](https://github.com/projetoKhali/api4back/blob/main/src/main/java/com/fatec/springapi4/Validation.java#L10-L58) e [`validateUser`](https://github.com/projetoKhali/api4back/blob/main/src/main/java/com/fatec/springapi4/Validation.java#L60-L64), integrando-as às rotinas de [`saveAndUpdatePartner`](https://github.com/projetoKhali/api4back/blob/main/src/main/java/com/fatec/springapi4/controller/PartnerController.java#L110-L112) e [`saveAndUpdateUser`](https://github.com/projetoKhali/api4back/blob/main/src/main/java/com/fatec/springapi4/controller/UserController.java#L75-L77) para validação de dados em requisições POST. Também realizei refatorações de legibilidade.",
			[]string{"Java", "Spring", "REST", "Paginação"},
		),
		models.NewCreateContributionByNames(
			"api4", "paulo-granthon",
			"Biblioteca de componentes Vue: Table, Form, Popup e Filter",
			"No front-end Vue 3 + TypeScript, construí boa parte da biblioteca de componentes reutilizáveis. Desenvolvi o componente [`Table`](https://github.com/projetoKhali/api4front/blob/main/src/components/Table.vue#L1-L255) (incluindo células-botão com roteamento e atualização manual de dados), os componentes de formulário [`Form`](https://github.com/projetoKhali/api4front/blob/main/src/components/form/Form.vue#L1-L156) e [`FormPopup`](https://github.com/projetoKhali/api4front/blob/main/src/components/form/FormPopup.vue#L1-L85), agrupei os componentes de [`Popup`](https://github.com/projetoKhali/api4front/blob/main/src/components/popup/Popup.vue#L1-L46) e criei o tipo [`PopupProps`](https://github.com/projetoKhali/api4front/blob/main/src/components/popup/NotificationPopup.vue#L3-L7), além do componente [`Filter`](https://github.com/projetoKhali/api4front/blob/main/src/components/Filter.vue#L1-L84). Apliquei tipagem genérica em inputs e centralizei a responsabilidade de notificação na função [`openNotifPopup`](https://github.com/projetoKhali/api4front/blob/main/src/views/ListPartner.vue#L104-L111) das views.",
			[]string{"Vue", "TypeScript", "Componentização"},
		),
		models.NewCreateContributionByNames(
			"api4", "paulo-granthon",
			"Roteamento, NavBar recursiva e geração de relatórios CSV",
			"Implementei o sistema de roteamento de páginas e o componente [`NavBar`](https://github.com/projetoKhali/api4front/blob/main/src/components/nav/NavBar.vue), evoluindo-o para rotas aninhadas com submenus por meio do componente [`RecursiveRouterLink`](https://github.com/projetoKhali/api4front/blob/main/src/components/nav/RecursiveRouterLink.vue#L1-L61). Implementei a geração e o download de relatórios em CSV para Partner, criando as funções genérica [`downloadCSV`](https://github.com/projetoKhali/api4front/blob/main/src/report/csv.ts#L6-L9) e específica [`downloadPartnerCSV`](https://github.com/projetoKhali/api4front/blob/main/src/report/partner.ts#L8-L23) (usando papaparse e file-saver). Também criei utilitários como [`removeSpecialCharacters`](https://github.com/projetoKhali/api4front/blob/main/src/util/index.ts#L47-L50) e [`getDisplayName`](https://github.com/projetoKhali/api4front/blob/main/src/util/index.ts#L1-L45) e configurei as regras de estilo de código (ESLint + Prettier) do front-end.",
			[]string{"Vue", "TypeScript"},
		),

		models.NewCreateContributionByNames(
			"api4", "paulo-granthon",
			"Camada de serviços tipada e paginada do front-end (classe Page)",
			"Estabeleci o padrão da camada de serviços do front criando o [`UserService`](https://github.com/projetoKhali/api4front/blob/main/src/service/UserService.ts) com os schemas [`UserSchema`](https://github.com/projetoKhali/api4front/blob/main/src/schemas/User.ts#L8-L13)/[`UserPatchSchema`](https://github.com/projetoKhali/api4front/blob/main/src/schemas/User.ts#L14) e mapeadores de resposta (PR [#39](https://github.com/projetoKhali/api4front/pull/39)); depois criei a classe [`Page`](https://github.com/projetoKhali/api4front/blob/main/src/schemas/Page.ts#L1-L44) definindo o contrato de respostas paginadas e adaptei [`getPartners`](https://github.com/projetoKhali/api4front/blob/main/src/service/PartnerService.ts#L63-L71)/[`getUsers`](https://github.com/projetoKhali/api4front/blob/main/src/service/UserService.ts#L36-L43) para receber [`page`](https://github.com/projetoKhali/api4front/blob/main/src/schemas/Page.ts#L1-L44) e [`size`](https://github.com/projetoKhali/api4front/blob/main/src/schemas/Page.ts#L1-L44), refatorando [`ListPartner`](https://github.com/projetoKhali/api4front/blob/main/src/views/ListPartner.vue)/[`ListUser`](https://github.com/projetoKhali/api4front/blob/main/src/views/ListUser.vue) para consumir os serviços (em vez de `fetch` direto) e usar os metadados de paginação (PR [#76](https://github.com/projetoKhali/api4front/pull/76)).",
			[]string{"Vue", "TypeScript", "Paginação", "Refatoração"},
		),
		models.NewCreateContributionByNames(
			"api4", "paulo-granthon",
			"Componente NotificationPopup e feedback centralizado de operações",
			"Criei o componente reutilizável [`NotificationPopup`](https://github.com/projetoKhali/api4front/blob/main/src/components/popup/NotificationPopup.vue#L1-L99), que exibe mensagens de sucesso/erro com fechamento automático, e a função centralizada [`openNotificationPopup`](https://github.com/projetoKhali/api4front/blob/main/src/views/ListPartner.vue#L104-L111) com duração configurável (PR [#107](https://github.com/projetoKhali/api4front/pull/107)). Integrei o feedback às views de listagem e exportação, refatorando as chamadas de serviço para `try/catch` com `async/await` e exibindo as mensagens de erro vindas das respostas da API.",
			[]string{"Vue", "TypeScript", "Componentização", "UI/UX"},
		),
		models.NewCreateContributionByNames(
			"api4", "paulo-granthon",
			"Dashboard comparativo de parceiros com gráficos empilhados",
			"Desenvolvi a view [`ComparativePartner`](https://github.com/projetoKhali/api4front/blob/main/src/views/ComparativePartner.vue#L1-L449) e o componente `StackedBarChartByTracks` para comparar métricas de parceiros por trilhas, com filtro por nome parcial, paginação e seleção via `CheckBoxTableCell`. Migrei [`getPartnerMetrics`](https://github.com/projetoKhali/api4front/blob/main/src/service/PartnerMetricsService.ts#L60-L63) de GET para POST, corrigi o loop infinito de [`getPageData`](https://github.com/projetoKhali/api4front/blob/main/src/views/ListPartner.vue#L240) e padronizei nomes de classes CSS, funções e schemas para consistência.",
			[]string{"Vue", "TypeScript", "Análise de Dados", "Componentização"},
		),
		models.NewCreateContributionByNames(
			"api4", "paulo-granthon",
			"Paginação do endpoint de métricas de parceiros (PartnerMetrics)",
			"Adicionei paginação à rota de métricas de parceiros, convertendo o endpoint de GET para POST para receber os parâmetros via um novo schema `PartnerMetricsRequest`, e corrigi nomes de campos da entidade [`PartnerMetrics`](https://github.com/projetoKhali/api4back/blob/main/src/main/java/com/fatec/springapi4/entity/PartnerMetrics.java#L27-L44) (PR [#63](https://github.com/projetoKhali/api4back/pull/63)). A mudança alinhou o contrato de métricas ao mesmo padrão paginado já adotado nos demais endpoints.",
			[]string{"Java", "Spring", "REST", "Paginação"},
		),
		models.NewCreateContributionByNames(
			"api4", "paulo-granthon",
			"Configuração de ESLint/Prettier e padronização de estilo de código",
			"Configurei o tooling de estilo do front-end (ESLint + Prettier), criando os arquivos de configuração e os scripts (incluindo [`lint:fix`](https://github.com/projetoKhali/api4front/blob/main/package.json)), e apliquei as regras à base existente (PR [#14](https://github.com/projetoKhali/api4front/pull/14)). Depois liderei a padronização de estilo também no back-end (PR [#78](https://github.com/projetoKhali/api4back/pull/78)), incluindo regra de `switch case` no ESLint e lint de markdown, alinhando os dois repositórios do projeto.",
			[]string{"ESLint", "Prettier", "TypeScript", "DevOps"},
		),

		// ----------------------------------------------------------------
		// Semester 5 - api5 (Go / TypeScript / Python)
		// ----------------------------------------------------------------
		models.NewCreateContributionByNames(
			"api5", "paulo-granthon",
			"Modelagem do data warehouse e entidades do back-end em Go com Ent ORM",
			"No repositório [`api5back`](https://github.com/projetoKhali/api5back), modelei e implementei as entidades do data warehouse de recrutamento utilizando o Ent ORM em Go, criando estruturas como [`HiringProcessCandidate`](https://github.com/projetoKhali/api5back/blob/main/src/schema/fact_hiring_process.go), [`DimCandidateStatus`](https://github.com/projetoKhali/api5back/blob/main/src/property/dim_candidate_status.go#L9), [`DimProcess`](https://github.com/projetoKhali/api5back/blob/main/src/schema/dim_process.go#L17) e [`DimVacancy`](https://github.com/projetoKhali/api5back/blob/main/src/schema/dim_vacancy.go#L16). Adicionei identificadores únicos ([`db_id`](https://github.com/projetoKhali/api5back/pull/54)) às entidades do DW (PR [#54](https://github.com/projetoKhali/api5back/pull/54)), padronizei o esquema de candidatos (PR [#67](https://github.com/projetoKhali/api5back/pull/67)) e renomeei [`GroupAccess`](https://github.com/projetoKhali/api5back/blob/main/src/schema/access_group.go#L14) para [`AccessGroup`](https://github.com/projetoKhali/api5back/blob/main/src/schema/access_group.go#L14) em todo o código. Também construí a função [`VacancyStatusSummary`](https://github.com/projetoKhali/api5back/blob/main/src/processing/vacancy_status.go#L11-L15) (PR [#22](https://github.com/projetoKhali/api5back/pull/22)), base para os gráficos do dashboard.",
			[]string{"Go", "Ent ORM", "Modelagem de Dados", "PostgreSQL"},
		),
		models.NewCreateContributionByNames(
			"api5", "paulo-granthon",
			"Endpoint de dashboard e consultas analíticas do processo seletivo",
			"Desenvolvi o endpoint do dashboard de processo seletivo (PR [#32](https://github.com/projetoKhali/api5back/pull/32)) no [`api5back`](https://github.com/projetoKhali/api5back), responsável por agregar as métricas exibidas na aplicação, incluindo o cálculo de tempo médio de contratação. Implementei a lógica de filtros das consultas ([`applyFactHiringProcessQueryFilters`](https://github.com/projetoKhali/api5back/blob/main/src/service/dashboard.go#L39-L113)), com verificação de nulidade de [`DepartmentIds`](https://github.com/projetoKhali/api5back/blob/main/src/service/dashboard.go#L39-L113), e otimizei as queries com `DISTINCT ON` por [`db_id`](https://github.com/projetoKhali/api5back/blob/main/src/service/dashboard.go#L39-L113) para retornar os estados mais recentes das entidades.",
			[]string{"Go", "SQL", "Análise de Dados"},
		),
		models.NewCreateContributionByNames(
			"api5", "paulo-granthon",
			"Endpoints de sugestões com paginação reutilizável",
			"No [`api5back`](https://github.com/projetoKhali/api5back), criei os endpoints de [`/suggestions`](https://github.com/projetoKhali/api5back/tree/main/src/service) com paginação (PR [#73](https://github.com/projetoKhali/api5back/pull/73)). Extraí a lógica de paginação para um módulo dedicado ([`ParsePageRequest`](https://github.com/projetoKhali/api5back/blob/main/src/pagination/pagination.go#L18-L46), [`ParseOffsetAndTotalPages`](https://github.com/projetoKhali/api5back/blob/main/src/pagination/pagination.go)), padronizei os nomes das funções do serviço de sugestões e adicionei os tratamentos de nulidade necessários. No front-end, garanti a compatibilidade do cliente com a paginação dos endpoints de sugestões (PR [#86](https://github.com/projetoKhali/api5front/pull/86)).",
			[]string{"Go", "Paginação"},
		),
		models.NewCreateContributionByNames(
			"api5", "paulo-granthon",
			"Suíte de testes de integração com containers de banco dedicados",
			"Estruturei a estratégia de testes do back-end em Go, criando testes de integração que executam um container de banco de dados dedicado (PR [#20](https://github.com/projetoKhali/api5back/pull/20)) e compartilham um único container entre múltiplos testes (PR [#23](https://github.com/projetoKhali/api5back/pull/23)) para reduzir o tempo de execução. Escrevi testes unitários de paginação, de propriedades e dos serviços de sugestões, consolidando a confiabilidade da camada de dados.",
			[]string{"Go", "Testes de Integração", "Testes Unitários", "Docker"},
		),
		models.NewCreateContributionByNames(
			"api5", "paulo-granthon",
			"Scripts de seeds e ETL do data warehouse",
			"No [`api5back`](https://github.com/projetoKhali/api5back), implementei os scripts de seeds do data warehouse (PRs [#24](https://github.com/projetoKhali/api5back/pull/24), #47, #65), refatorando os dados de [`dw_base`](https://github.com/projetoKhali/api5back/blob/main/seeds/dw_base.go) em constantes públicas reutilizáveis, de-duplicando nomes de vagas e definindo o banco-alvo via [`SeedsPreset`](https://github.com/projetoKhali/api5back/tree/main/seeds). Criei o script [`drop-all`](https://github.com/projetoKhali/api5back/blob/main/drop-all.bash) para reset do ambiente e corrigi a planilha do ETL quanto a status de candidatos inválidos, além de mensagens de erro mais informativas.",
			[]string{"Go", "ETL", "Modelagem de Dados"},
		),
		models.NewCreateContributionByNames(
			"api5", "paulo-granthon",
			"Pipeline de CI e configuração de ambiente do back-end",
			"Criei o workflow de CI do [`api5back`](https://github.com/projetoKhali/api5back) e ajustei o branch-alvo do deploy automático. Adicionei a leitura de [`SSLMODE`](https://github.com/projetoKhali/api5back/blob/main/src/database/credentials.go#L82) a partir de variáveis de ambiente, restaurei a configuração de CORS no [`main`](https://github.com/projetoKhali/api5back/blob/main/main.go), incluí a dependência `gin-metrics` para métricas, ajustei o build tag de produção no Makefile e gerei a documentação Swagger dos endpoints, integrando os comandos do Swag à rotina de desenvolvimento.",
			[]string{"CI/CD", "GitHub Actions", "Go", "Swagger"},
		),
		models.NewCreateContributionByNames(
			"api5", "paulo-granthon",
			"Filtros e integração com a API no front-end em TypeScript",
			"No [`api5front`](https://github.com/projetoKhali/api5front), implementei os filtros de status do dashboard, o componente de multi-seleção (com botões de limpar lista e limpar filtros) e a aplicação automática dos filtros ao alterar valores. Refatorei a camada de serviços para consumir a URL da API via variável de ambiente, criei o tipo [`Method`](https://github.com/projetoKhali/api5front/blob/main/src/service/base.ts#L13) no serviço base para suportar métodos além de POST e corrigi a navegação para o dashboard no login.",
			[]string{"TypeScript", "Componentização"},
		),
		models.NewCreateContributionByNames(
			"api5", "paulo-granthon",
			"Configuração do monorepo e padronização de commits e hooks",
			"No repositório umbrella [`api5`](https://github.com/projetoKhali/api5), estruturei o monorepo com submódulos Git para back-end e front-end e criei comandos de orquestração no [`package.json`](https://github.com/projetoKhali/api5/blob/main/package.json). Configurei os git hooks com Husky para garantir estilo de código, execução de testes e mensagens de commit semânticas, corrigindo a sintaxe da chamada de testes no hook [`pre-commit`](https://github.com/projetoKhali/api5/tree/main/hooks) e permitindo mensagens padrão de merge no hook [`commit-msg`](https://github.com/projetoKhali/api5/tree/main/hooks).",
			[]string{"Git", "Husky"},
		),

		models.NewCreateContributionByNames(
			"api5", "paulo-granthon",
			"Configuração de SSL e ambiente para conexão de banco em produção",
			"Implementei a leitura do parâmetro [`SSLMODE`](https://github.com/projetoKhali/api5back/blob/main/src/database/credentials.go#L82) a partir de variáveis de ambiente na camada de conexão do back-end Go (PR [#72](https://github.com/projetoKhali/api5back/pull/72)), permitindo conexão segura ao banco gerenciado em produção sem hardcode. A mudança parametriza o modo SSL/TLS por ambiente (desenvolvimento, homologação e produção), viabilizando o deploy real do serviço contra o PostgreSQL de produção.",
			[]string{"Go", "PostgreSQL", "DevOps"},
		),
		models.NewCreateContributionByNames(
			"api5", "paulo-granthon",
			"Testes unitários do módulo property e do enum DimCandidateStatus",
			"Criei testes unitários para o módulo [`property`](https://github.com/projetoKhali/api5back/tree/main/src/property), incluindo a cobertura do enum [`DimCandidateStatus`](https://github.com/projetoKhali/api5back/blob/main/src/property/dim_candidate_status.go#L9), e refatorei as estruturas de casos de teste para um arquivo dedicado dentro do módulo (PR [#84](https://github.com/projetoKhali/api5back/pull/84)), ampliando a cobertura para além dos testes de integração com containers de banco.",
			[]string{"Go", "Testes Unitários", "Refatoração"},
		),
		models.NewCreateContributionByNames(
			"api5", "paulo-granthon",
			"Camada de serviços HTTP do front-end com método parametrizável",
			"Refatorei a camada base de serviços do front-end introduzindo um tipo [`Method`](https://github.com/projetoKhali/api5front/blob/main/src/service/base.ts#L13) para permitir verbos HTTP além de POST e um utilitário [`processRequest`](https://github.com/projetoKhali/api5front/blob/main/src/service/base.ts#L18-L24) para padronizar as chamadas à API, incluindo o serviço de Login (PRs [#95](https://github.com/projetoKhali/api5front/pull/95) e [#86](https://github.com/projetoKhali/api5front/pull/86)). Corrigi imports de schemas de [`Suggestion`](https://github.com/projetoKhali/api5front/blob/main/src/service/Suggestions.ts), ajustei a versão da API em URLs de template e renomeei [`GroupAccess`](https://github.com/projetoKhali/api5front/blob/main/src/service/AccessGroup.ts) para [`AccessGroup`](https://github.com/projetoKhali/api5front/blob/main/src/service/AccessGroup.ts) por clareza semântica.",
			[]string{"TypeScript", "React Native", "Refatoração", "REST"},
		),
		models.NewCreateContributionByNames(
			"api5", "paulo-granthon",
			"Infraestrutura de testes Jest no front-end",
			"Configurei a infraestrutura de testes do front-end (PR [#85](https://github.com/projetoKhali/api5front/pull/85)): corrigi o comando de teste, atualizei o [`tsconfig`](https://github.com/projetoKhali/api5front/blob/main/tsconfig.json) para reconhecer o diretório [`__tests__`](https://github.com/projetoKhali/api5front/tree/main/__tests__), configurei o [Jest](https://github.com/projetoKhali/api5front/blob/main/jest.config.js) com arquivos de configuração dedicados em vez de comandos inline e refinei os comandos de watch, preparando o projeto para o desenvolvimento de testes.",
			[]string{"TypeScript", "Jest", "DevOps"},
		),
		models.NewCreateContributionByNames(
			"api5", "paulo-granthon",
			"Setup inicial do front-end com Expo/React Native e git hooks",
			"Realizei o setup inicial do front-end com projeto Expo/React Native e roteamento via `expo-router` (PR [#6](https://github.com/projetoKhali/api5front/pull/6)), e configurei a rotina de qualidade com Husky e lint-staged: hook de pre-commit para formatação e testes, hook de commit-msg para commits semânticos e script de postinstall para inicializar os hooks.",
			[]string{"TypeScript", "React Native", "Expo", "Husky"},
		),

		// ----------------------------------------------------------------
		// Semester 6 - api6 (Nx monorepo: React / Flask / Rust)
		// ----------------------------------------------------------------
		models.NewCreateContributionByNames(
			"api6", "paulo-granthon",
			"Serviço de autenticação em Rust/Actix-web com JWT",
			"Concebi e implementei do zero o app [`auth`](https://github.com/projetoKhali/api6/tree/main/apps/auth), um microsserviço de autenticação em Rust com o framework Actix-web (PRs #64 e #70). O serviço oferece login por credenciais, geração e validação de tokens JWT com assinatura criptográfica, revogação de tokens via tabela [`revoked_tokens`](https://github.com/projetoKhali/api6/blob/main/apps/auth/src/entities/revoked_token.rs#L3-L5) no PostgreSQL, middleware de proteção de rotas e documentação OpenAPI/Swagger. Também o integrei ao restante do sistema com entry point Docker, retorno de permissões no login e correções de segurança como impedir o login de usuário inativo.",
			[]string{"Rust", "Actix-web", "JWT", "Autenticação", "PostgreSQL", "Docker"},
		),
		models.NewCreateContributionByNames(
			"api6", "paulo-granthon",
			"Autenticação de clientes externos e portabilidade de dados",
			"Projetei e implementei todo o fluxo de clientes externos e portabilidade de dados (PR #117). Modelei a tabela [`external_clients`](https://github.com/projetoKhali/api6/blob/main/apps/auth/src/entities/external_client.rs#L5-L6) e refatorei [`user_key`](https://github.com/projetoKhali/api6/blob/main/apps/auth/src/entities/entity_key.rs) para suportar chaves criptográficas de múltiplas entidades, com campo [`entity_type`](https://github.com/projetoKhali/api6/blob/main/apps/auth/src/entities/entity_key.rs#L9) e restrição de unicidade. Implementei o CRUD e a autenticação de clientes externos, ampliei o JWT para suportar dupla entidade e construí o fluxo de portabilidade em três etapas (botão, tela de autorização com geração de token, e troca do token por dados). Entreguei ainda um app de exemplo simulando o cliente externo.",
			[]string{"Rust", "JWT", "Modelagem de Dados", "PostgreSQL"},
		),
		models.NewCreateContributionByNames(
			"api6", "paulo-granthon",
			"Infraestrutura de CI/CD, monorepo Nx e tooling de testes",
			"Estabeleci a base de infraestrutura do projeto (PR #8) com pipeline de CI/CD em GitHub Actions executando testes e validações automatizadas, orquestração de tarefas via Nx (install, test, test-integration, lint, lock) e Nx Cloud. Configurei o gerenciamento de dependências Python com Poetry e ambiente virtual compartilhado, adicionei pytest, testcontainers e Black, e criei o [`.env.example`](https://github.com/projetoKhali/api6/blob/main/.env.example) para PostgreSQL e MongoDB. Também tratei ajustes de Nx ao longo do projeto e a unificação dos comandos de seed.",
			[]string{"GitHub Actions", "CI/CD", "Nx", "Poetry", "Docker"},
		),
		models.NewCreateContributionByNames(
			"api6", "paulo-granthon",
			"Modelo e rotas de produtividade (yield) na API Python/Flask",
			"Implementei o modelo de eventos de produtividade ([`YieldEvent`](https://github.com/projetoKhali/api6/blob/main/apps/api/src/api/models/yield_model.py#L5-L18)) usando Pydantic na API Python (PR #30), com suíte de testes unitários usando MagicMock e mock de MongoDB. Em seguida desenvolvi o mecanismo de rotas de yield (PR #34), reorganizando a lógica de servidor do módulo [`server`](https://github.com/projetoKhali/api6/blob/main/apps/api/src/api/server.py) para [`api`](https://github.com/projetoKhali/api6/tree/main/apps/api/src/api), criando o registro padronizado de blueprints Flask, aplicando CORS em toda a aplicação e corrigindo a inicialização do MongoDB (variável de database, [`authSource=admin`](https://github.com/projetoKhali/api6/blob/main/apps/db/src/db/mongo.py#L24) e teste de conexão antes de instanciar o Flask).",
			[]string{"Python", "Flask", "Pydantic", "MongoDB", "Testes Unitários", "REST"},
		),
		models.NewCreateContributionByNames(
			"api6", "paulo-granthon",
			"Componente React de paginação reutilizável e gestão de usuários",
			"Desenvolvi o componente React/TypeScript [`Pagination`](https://github.com/projetoKhali/api6/blob/main/apps/web/src/components/Pagination.tsx#L14-L125) (PR #38), com botões de primeira/última página, controles anterior/próximo e suporte a sobrescrita de estilos para reuso. Posteriormente integrei a gestão de usuários ponta a ponta (PR #93), construindo a [`UserManagementPage`](https://github.com/projetoKhali/api6/blob/main/apps/web/src/pages/UserManagementPage.tsx#L115-L301) com listagem, paginação e exclusão com confirmação, e no back-end Rust adicionei paginação ao endpoint [`/users/`](https://github.com/projetoKhali/api6/blob/main/apps/auth/src/routes/user.rs#L60) com a struct [`PaginatedResponse`](https://github.com/projetoKhali/api6/blob/main/apps/auth/src/models/pagination.rs#L11-L16) e conversão automática snake_case/camelCase.",
			[]string{"React", "TypeScript", "Componentização", "Rust", "Paginação"},
		),
		models.NewCreateContributionByNames(
			"api6", "paulo-granthon",
			"Criptografia de campos sensíveis com Fernet no serviço de autenticação",
			"Implementei um módulo [`fernet`](https://github.com/projetoKhali/api6/blob/main/apps/auth/src/service/fernet.rs) no app de autenticação em Rust, responsável por cifrar e decifrar os campos sensíveis dos usuários. Criei a função [`encrypt_field`](https://github.com/projetoKhali/api6/blob/main/apps/auth/src/service/fernet.rs#L119-L122), integrei a criptografia nos endpoints de [`register`](https://github.com/projetoKhali/api6/blob/main/apps/auth/src/routes/auth.rs#L89), [`login`](https://github.com/projetoKhali/api6/blob/main/apps/auth/src/routes/auth.rs#L153) e [`update_user`](https://github.com/projetoKhali/api6/blob/main/apps/auth/src/routes/user.rs#L217) e a decriptação nas rotas GET, recuperando a chave de cada usuário a partir de uma entidade [`keys`](https://github.com/projetoKhali/api6/blob/main/apps/db/src/db/keys.py) dedicada. Tratei a codificação base64, defini quais campos não deveriam ser cifrados (como [`permission_id`](https://github.com/projetoKhali/api6/blob/main/apps/auth/src/entities/permission.rs)) e modelei um banco de chaves ([`api6_keys`](https://github.com/projetoKhali/api6/blob/main/apps/db/src/db/keys.py#L30)) com conexão própria, separando os dados criptográficos do banco principal.",
			[]string{"Rust", "Actix-web", "Criptografia", "Fernet", "SeaORM", "PostgreSQL"},
		),
		models.NewCreateContributionByNames(
			"api6", "paulo-granthon",
			"Tratamento de erros customizado e modo de desenvolvimento na API de auth",
			"Criei a struct [`CustomError`](https://github.com/projetoKhali/api6/blob/main/apps/auth/src/routes/common/error.rs#L8-L17) e renomeei [`ErrorType`](https://github.com/projetoKhali/api6/blob/main/apps/auth/src/routes/common/error.rs#L19-L25) para [`ServerErrorType`](https://github.com/projetoKhali/api6/blob/main/apps/auth/src/routes/common/error.rs#L19-L25), estruturando o tratamento de erros do serviço Rust. Adicionei detalhamento de erros de endpoint em modo de desenvolvimento (impressão no terminal e inclusão dos detalhes na resposta) e log condicional de `DatabaseError` quando em [`dev_mode`](https://github.com/projetoKhali/api6/blob/main/apps/auth/src/infra/config.rs#L39), melhorando a observabilidade durante o desenvolvimento.",
			[]string{"Rust", "Actix-web", "Refatoração"},
		),
		models.NewCreateContributionByNames(
			"api6", "paulo-granthon",
			"Centralização de configuração e infraestrutura de logging do serviço Rust",
			"Centralizei o parsing de variáveis de ambiente em um módulo [`config`](https://github.com/projetoKhali/api6/blob/main/apps/auth/src/infra/config.rs), configurei a infraestrutura de logging do app de autenticação e estruturei as conexões de banco via SeaORM, incluindo a conexão dedicada ao banco [`api6_keys`](https://github.com/projetoKhali/api6/blob/main/apps/db/src/db/keys.py#L30). Usei o nome do container como host padrão dos bancos para funcionar em ambiente Docker e tratei os erros das structs de configuração.",
			[]string{"Rust", "SeaORM", "Docker", "DevOps"},
		),
		models.NewCreateContributionByNames(
			"api6", "paulo-granthon",
			"Containerização do serviço de autenticação com Docker e alvo Nx",
			"Criei o [`Dockerfile`](https://github.com/projetoKhali/api6/blob/main/apps/auth/Dockerfile) e o [`docker-compose`](https://github.com/projetoKhali/api6/blob/main/apps/auth/docker-compose.yml) do app de autenticação em Rust e o comando Nx [`serve-docker`](https://github.com/projetoKhali/api6/blob/main/apps/auth/project.json#L11) para subir o serviço containerizado (PR #70). Garanti a inclusão das variáveis de ambiente do banco de chaves no build Docker e padronizei os containers de banco no docker-compose.",
			[]string{"Docker", "Docker Compose", "Nx", "Rust", "DevOps"},
		),
		models.NewCreateContributionByNames(
			"api6", "paulo-granthon",
			"Seeds de PostgreSQL e MongoDB com criptografia de credenciais",
			"Desenvolvi os seeds de banco do app [`db`](https://github.com/projetoKhali/api6/tree/main/apps/db): o seed do PostgreSQL com inserção de usuários abstraída e reutilizável, atribuição de [`permission_id`](https://github.com/projetoKhali/api6/blob/main/apps/db/seeds/postgres_seeds.py)/[`role_id`](https://github.com/projetoKhali/api6/blob/main/apps/db/seeds/postgres_seeds.py#L151) e criptografia das senhas e campos sensíveis já na carga inicial. Criei o comando Nx unificado [`seeds`](https://github.com/projetoKhali/api6/blob/main/apps/db/project.json#L57) orquestrando a inicialização de Postgres e Mongo (com [`init-keys`](https://github.com/projetoKhali/api6/blob/main/apps/db/project.json#L37) e [`init-postgres`](https://github.com/projetoKhali/api6/blob/main/apps/db/project.json#L23) como dependências), restringi a execução automática do seed no init do Mongo e ajustei a inicialização de sessão e os commits transacionais.",
			[]string{"Python", "PostgreSQL", "MongoDB", "Criptografia", "Nx"},
		),
		models.NewCreateContributionByNames(
			"api6", "paulo-granthon",
			"Página de cadastro de produtividade e camada de serviços/AuthService no front React",
			"Criei a página de registro de produtividade ([`/register/yield`](https://github.com/projetoKhali/api6/blob/main/apps/web/src/pages/YieldRegister.tsx#L15-L78)) em React, com componentes de formulário e de tabela e os schemas [`Yield`](https://github.com/projetoKhali/api6/blob/main/apps/web/src/schemas/yield.ts#L3-L13)/[`Season`](https://github.com/projetoKhali/api6/blob/main/apps/web/src/schemas/yield.ts#L1) com suas funções de serviço (PR #36). Estabeleci a camada de serviços HTTP genérica ([`processRequest`](https://github.com/projetoKhali/api6/blob/main/apps/web/src/service/service.ts#L48)/[`processGET`](https://github.com/projetoKhali/api6/blob/main/apps/web/src/service/service.ts#L67)) tipada e o [`AuthService`](https://github.com/projetoKhali/api6/blob/main/apps/web/src/service/AuthService.ts) usado pela página de Login, integrando o fluxo de autenticação ao serviço Rust, com leitura do token do `localStorage`, unificação da pasta de serviços e correção da lógica de redirecionamento.",
			[]string{"React", "TypeScript", "Autenticação", "REST", "Refatoração"},
		),
	}

	for _, n := range exampleContributions {
		if _, err := service.ContributionService.Create(n); err != nil {
			return tracerr.Errorf("failed to create contribution: %w", tracerr.Wrap(err))
		}
	}

	return nil
}

func ParticipationMigrate(
	storage storage.Storage,
	service service.Service,
) error {
	projectModule, err := storage.GetProjectModule()
	if err != nil {
		return tracerr.Errorf("failed to get project module: %w", tracerr.Wrap(err))
	}

	userModule, err := storage.GetUserModule()
	if err != nil {
		return tracerr.Errorf("failed to get user module: %w", tracerr.Wrap(err))
	}

	participationModule, err := storage.GetParticipationModule()
	if err != nil {
		return tracerr.Errorf("failed to get participation module: %w", tracerr.Wrap(err))
	}

	summaries := []models.CreateParticipationByNames{
		models.NewCreateParticipationByNames(
			"Khali", "paulo-granthon",
			"Tive uma participação intensa e tecnicamente central no projeto Khali, com grande volume de contribuições ao longo de todo o desenvolvimento. Atuei como engenheiro das fundações do front-end, integrador e desenvolvedor full-stack: criei os alicerces da interface (a biblioteca de helpers e o design system do [`Front.Core`](https://github.com/projetoKhali/Khali/blob/main/Front/Core.py), o componente de área rolável com Canvas e o [`WindowManager`](https://github.com/projetoKhali/Khali/blob/main/Front/WindowManager.py) que governa a navegação como máquina de estados) e o sistema de módulos da Home com controle de acesso por papel. Trabalhei intensamente na camada de dados e visualização, escrevendo os cálculos de médias multidimensionais (por critério, sprint, time, grupo e papel) e os gráficos dos dashboards, e implementei telas centrais como a de avaliação 360° e os cadastros de grupos e times. Meus diferenciais foram a prototipação da KML (uma DSL declarativa para Tkinter) e o sistema de eventos publish-subscribe que desacoplou as telas. Por fim, fui o integrador de fato do repositório, conduzindo inúmeros merges e o fechamento do release. Onde mais me destaquei: arquitetura, abstrações reutilizáveis e a cola entre as partes do sistema.",
		),
		models.NewCreateParticipationByNames(
			"API2Semestre", "paulo-granthon",
			"No segundo semestre fui um dos pilares técnicos do projeto, com 32 PRs (a maioria de alto impacto) e centenas de commits no sistema desktop Java/JavaFX. Minha marca foi a engenharia de fundações reutilizáveis e a definição da arquitetura: reestruturei todo o projeto para o padrão Maven com módulo JavaFX, criei o sistema de tags FXML para configuração de telas com permissões e produzi abstrações que o time inteiro consumiu (a macro de células editáveis [`TableMacros`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/view/macros/TableMacros.java), o [`ChartGenerator`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/view/utils/ChartGenerator.java#L36), o [`LookupTextField`](https://github.com/projetoKhali/API2Semestre/blob/main/src/main/java/org/openjfx/API2Semestre/view/controllers/custom_tags/LookupTextField.java#L27) e as macros de filtro). Atuei full-stack dentro do desktop, modelando o domínio, construindo as telas centrais (Listagem, Aprovação, Dashboard, Relatórios e Parametrização) e otimizando a camada de dados com reuso de conexões. Tomei decisões técnicas fundamentadas e assumi o papel de integrador e mantenedor da qualidade na reta final. Onde mais brilhei: na criação de componentes e abstrações reaproveitáveis e na arquitetura do projeto.",
		),
		models.NewCreateParticipationByNames(
			"api3", "paulo-granthon",
			"Tive uma participação ampla e versátil no api3, com cerca de 366 commits e 32 PRs cobrindo toda a stack (React e TypeScript no front, Java e Spring no back, PostgreSQL e Docker). Estabeleci as fundações do projeto, do scaffolding inicial à dockerização completa, ao [`CorsConfig`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/config/CorsConfig.java#L8-L16), à documentação Swagger e às ferramentas de produtividade. No back-end fui o principal responsável pela camada de apontamentos e pelo sistema de permissões; meu trabalho de maior destaque técnico foi o motor de cálculo de horas ([`SliceCalculator`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/util/SliceCalculator.java#L24-L267)) e a persistência engenhosa dos dias da semana das regras de verba como bitmask via [`DaysOfWeekConverter`](https://github.com/projetoKhali/api3/blob/main/api/src/main/java/com/khali/api3/converters/days_of_week/DaysOfWeekConverter.java#L8-L29). No front-end criei um arsenal de componentes reutilizáveis (Dropdown, LookUpTextField, células de tabela editáveis, SchemaList), o serviço de menu dinâmico por permissão e várias telas, e liderei a padronização visual. Onde mais me destaquei: na transversalidade, resolvendo algoritmos de negócio e modelagem de banco e, ao mesmo tempo, entregando UI polida e infraestrutura, atuando de ponta a ponta como desenvolvedor full-stack.",
		),
		models.NewCreateParticipationByNames(
			"api4", "paulo-granthon",
			"Como Scrum Master da equipe no quarto semestre, tive forte presença técnica no projeto Oracle, com grande volume de entregas nos dois repositórios de código ([api4back](https://github.com/projetoKhali/api4back) e [api4front](https://github.com/projetoKhali/api4front)) e conduzindo a integração no repositório-umbrella. Atuei como arquiteto transversal: no front Vue/TypeScript construí praticamente toda a biblioteca de componentes genéricos (Table, Form/FormPopup, Filter, NotificationPopup), o roteamento, a camada de serviços tipada e paginada, a geração de relatórios CSV e o dashboard comparativo de parceiros; no back Java/Spring entreguei validação, paginação determinística de endpoints e modelagem e seeds do PostgreSQL. Meu trabalho mais singular foi a fundação técnica e o ferramental: ESLint e Prettier e padronização de estilo nos dois repositórios, setup do monorepo com [`concurrently`](https://github.com/projetoKhali/api4/blob/main/package.json#L13) e submódulos, e os templates de PR e issue que organizaram o processo. Onde mais me destaquei: na criação de abstrações reutilizáveis no front e na garantia de qualidade e consistência transversal, combinando a liderança de processo com uma execução técnica intensa.",
		),
		models.NewCreateParticipationByNames(
			"api5", "paulo-granthon",
			"No quinto semestre fui um dos principais engenheiros de back-end e de dados do projeto Pro4tech, com grande volume de entregas (24 PRs no [api5back](https://github.com/projetoKhali/api5back), 16 no [api5front](https://github.com/projetoKhali/api5front) e 6 no monorepo). Meu trabalho de maior peso concentrou-se no back-end Go e no data warehouse: modelei e evoluí as entidades dimensionais (criação e renomeação de [`DimCandidate`](https://github.com/projetoKhali/api5back/blob/main/src/schema/dim_candidate.go#L16) e a introdução de [`db_id`](https://github.com/projetoKhali/api5back/pull/54)), desenhei as queries analíticas do dashboard e as funções de agregação, e estruturei os endpoints de sugestões com paginação reutilizável. Destaquei-me na engenharia de qualidade e infraestrutura, sendo responsável pela suíte de testes de integração com containers de banco, pela automação de Swagger, pelos scripts de seeds e ETL e pela configuração de ambiente (SSLMODE, CORS e URL da API) que viabilizou o deploy. No front-end TypeScript/React Native tive papel relevante porém mais secundário, focado em peças de fundação ([MultiSelectFilter](https://github.com/projetoKhali/api5front/blob/main/src/components/MultiSelectFilter.tsx#L67-L326), camada de serviços, setup do projeto e testes com [Jest](https://github.com/projetoKhali/api5front/blob/main/jest.config.js)). Em todos os repositórios fui o autor do setup do monorepo, dos git hooks e da padronização de commits. Onde mais brilhei: na modelagem e nas consultas do data warehouse em Go e na infraestrutura que sustentou o projeto.",
		),
		models.NewCreateParticipationByNames(
			"api6", "paulo-granthon",
			"Atuando como Product Owner da equipe no semestre final, tive uma atuação técnica bastante abrangente no projeto SIPMR (Kersys), presente em todas as camadas ao longo de mais de 500 commits e 18 PRs. Fui o principal responsável pelo serviço de autenticação em Rust/Actix-web (JWT, validação, criptografia Fernet de campos sensíveis, tratamento de erros, configuração e logging), pela modelagem dos bancos PostgreSQL e MongoDB (entidades SeaORM, chaves, clientes externos e tokens revogados) e pelos seeds com criptografia de credenciais. Na API Python/Flask modelei o domínio de produtividade com Pydantic e MongoDB e criei o middleware de autenticação; no front React entreguei componentes reutilizáveis, páginas de cadastro e gestão de usuários e a camada de serviços/AuthService. Fui ainda o arquiteto da infraestrutura: o monorepo Nx com Poetry, o pipeline de CI e a containerização Docker. Meu trabalho mais singular foi o fluxo completo de clientes externos e portabilidade de dados, que exigiu integrar Rust, banco e segurança. Onde mais me destaquei: na amplitude e na cola entre camadas, na arquitetura de autenticação e segurança, na modelagem de dados e no tooling do monorepo.",
		),
	}

	for _, p := range summaries {
		project, err := projectModule.GetByName(p.Project)
		if err != nil {
			return tracerr.Errorf("failed to get project %q for participation: %w", p.Project, tracerr.Wrap(err))
		}

		user, err := userModule.GetByUsername(p.User)
		if err != nil {
			return tracerr.Errorf("failed to get user %q for participation: %w", p.User, tracerr.Wrap(err))
		}

		if _, err := participationModule.Create(
			models.NewCreateParticipation(project.Id, user.Id, p.Summary),
		); err != nil {
			return tracerr.Errorf("failed to create participation: %w", tracerr.Wrap(err))
		}
	}

	return nil
}
