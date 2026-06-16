# Portfólio — Paulo Granthon

![Paulo Granthon](https://github.com/paulo-granthon.png?size=200)

## Introdução

<div align="justify">

Comecei a programar como hobby quase que uma década antes de entrar na FATEC. Nesse período desenvolvi jogos em C# com Unity, escrevi código em Java, Python com Django, HTML e JavaScript, e aprendi a usar Git e GitHub, explorando de forma independente. Quando cheguei ao curso já tinha uma bagagem técnica, mas enxerguei na FATEC algo que o aprendizado solo não consegue oferecer: a experiência de construir software em equipe, sob pressão real de prazos e com responsabilidades compartilhadas.

Entrei também em busca de uma formação oficial para o currículo, de preencher lacunas de conhecimento que só aparecem quando você trabalha com outras pessoas, e de uma rota mais acessível para o mercado — o estágio é uma porta que uma graduação abre com muito mais facilidade do que a caçada direta a vagas de junior.

Dentro dos projetos, minha experiência prévia me colocou naturalmente num papel de apoio técnico para o time — ensinei Git, expliquei conceitos de arquitetura, e fui procurado para ajudar a resolver bugs e tomar decisões técnicas. Mas as decisões sempre foram coletivas: nenhuma escolha passou sem consenso da equipe, e minha opinião não pesava mais do que a de qualquer outro integrante. Esse papel de apoio foi diminuindo de intensidade ao longo dos semestres, conforme o time crescia e as pessoas construíam a própria autonomia — o que era exatamente o que devia acontecer.

Atualmente atuo como desenvolvedor backend na Gorila Invest, aplicando na prática muito do que construí ao longo desses seis semestres.

</div>

## Contatos

[![GitHub](https://img.shields.io/badge/GitHub-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/paulo-granthon)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://linkedin.com/in/paulo-granthon)

## Meus Principais Conhecimentos

- Python / Django
- Tkinter
- Java / JavaFX / Spring Boot
- JavaScript / TypeScript
- React / Vue
- Go
- Rust / Actix Web
- PostgreSQL / MongoDB
- Docker / Docker Compose
- Nx (monorepo)
- Modelagem de Dados
- APIs REST
- Git / GitHub

---

## Meus Projetos

---

<details>
<summary><strong>1º Semestre — Khali</strong></summary>

<br/>

![Khali](https://user-images.githubusercontent.com/111442399/194777358-24905c4f-e62b-414d-8754-b3ccaf878547.png)

- **Empresa:** FATEC / PBLTeX
- **Repositório:** [projetoKhali/Khali](https://github.com/projetoKhali/Khali)

### Missão

Desenvolver uma plataforma desktop para aplicação do método de Avaliação 360°, voltada ao contexto educacional da instituição fictícia PBLTeX.

### Tecnologias Utilizadas

- **Python / Tkinter:** Interface gráfica desktop e toda a lógica da aplicação.
- **Matplotlib:** Geração dos dashboards e gráficos de desempenho.
- **CSV:** Camada de persistência de dados.
- **Git / GitHub:** Controle de versão e colaboração da equipe.

### Minha participação

Trabalhei nas fundações do projeto: o design system do `Front.Core`, o `WindowManager` para navegação entre telas, a camada de persistência com `CSVHandler`, os cálculos de médias e os dashboards com Matplotlib. Também desenvolvi telas centrais como a avaliação 360°, os cadastros e o sistema de módulos da Home com controle de acesso por papel.

### Hard Skills

- **Python:** Sei fazer com autonomia.
- **Tkinter:** Sei fazer com autonomia.
- **Modelagem de Dados:** Sei fazer com autonomia.
- **Análise de Dados:** Sei fazer com autonomia.
- **Componentização:** Sei fazer com autonomia.
- **Git:** Sei fazer com autonomia.

### Soft Skills

O primeiro semestre foi, acima de tudo, um exercício de começar do zero com outras pessoas. Não havia padrão estabelecido de como trabalhar juntos, como dividir tarefas, como comunicar decisões técnicas. A principal soft skill que precisei desenvolver foi a **comunicação técnica**: como explicar para colegas sem o mesmo histórico de programação por que uma abordagem era melhor que outra, ou como usar Git sem travar o repositório. Passei boa parte do projeto ensinando Git e os fundamentos de versionamento para o time, o que exigiu paciência e clareza.

Também foi o semestre em que percebi a importância de estabelecer **padrões antes de escalar**: o que definimos nas primeiras sprints criou uma base que o time inteiro passou a usar. Isso exigiu **proatividade** para criar componentes reutilizáveis e documentação mínima antes mesmo que alguém pedisse. O **trabalho em equipe** aqui significou aceitar que nem sempre o resultado seria o que eu faria sozinho — e que isso era parte do processo.

</details>

---

<details>
<summary><strong>2º Semestre — API2Semestre</strong></summary>

<br/>

![API2Semestre](https://raw.githubusercontent.com/projetoKhali/API2Semestre/main/Docs/Banners/Novobanner.png)

- **Empresa:** 2RP Net
- **Repositório:** [projetoKhali/API2Semestre](https://github.com/projetoKhali/API2Semestre)

### Missão

Desenvolver uma aplicação desktop em Java/JavaFX para registro, aprovação e gestão de horas extras e sobreaviso, com parametrização de regras, relatórios e dashboard analítico.

### Tecnologias Utilizadas

- **Java / JavaFX:** Lógica de negócio e interface desktop.
- **JDBC / PostgreSQL:** Camada de acesso a dados.
- **Maven:** Build e gerenciamento de dependências.
- **SQL:** Consultas e modelagem do banco.
- **Git / GitHub:** Controle de versão.

### Minha participação

Contribuí na reestruturação do projeto para o padrão Maven, criei componentes reutilizáveis para o time (`TableMacros`, `ChartGenerator`, `LookupTextField`), implementei a camada de persistência com `Query.java` tipado, a autenticação, o fluxo de aprovação de apontamentos, dashboards, relatórios e a tela de parametrização de regras.

### Hard Skills

- **Java:** Sei fazer com autonomia.
- **JavaFX:** Sei fazer com autonomia.
- **Maven:** Sei fazer com autonomia.
- **JDBC / SQL:** Sei fazer com autonomia.
- **Arquitetura de software:** Sei fazer com autonomia.

### Soft Skills

Esse semestre consolidou meu papel de **liderança técnica** dentro do time. Por ser o primeiro contato de boa parte da equipe com Java e com um projeto desktop estruturado, precisei criar uma arquitetura que fosse compreensível o suficiente para que todos pudessem contribuir. Isso exigiu **clareza de comunicação**: uma abstração só tem valor se quem vai usá-la consegue entendê-la.

Também desenvolvi minha **organização** ao assumir a responsabilidade por componentes centrais do projeto — qualquer bug nessas peças impactava o trabalho de todo mundo. Precisei ser cuidadoso, documentar decisões e estar disponível para explicar o que havia construído. O **trabalho em equipe** aqui ganhou uma dimensão nova: não era só colaborar, era garantir que meu trabalho habilitasse o trabalho dos outros.

</details>

---

<details>
<summary><strong>3º Semestre — api3</strong></summary>

<br/>

![api3](https://raw.githubusercontent.com/projetoKhali/api3/main/docs/Banners/Api.png)

- **Empresa:** 2RP Net
- **Repositório:** [projetoKhali/api3](https://github.com/projetoKhali/api3)

### Missão

Evoluir a solução desktop do semestre anterior para uma aplicação web full stack, mantendo o domínio de apontamentos, aprovação e parametrização de verbas em uma arquitetura mais distribuída.

### Tecnologias Utilizadas

- **React / TypeScript:** Frontend web.
- **Java / Spring Boot:** Backend REST.
- **PostgreSQL:** Banco de dados relacional.
- **Docker / Docker Compose:** Containerização e ambiente.
- **Swagger:** Documentação da API.

### Minha participação

Atuei de forma transversal em toda a stack: dockerizei o projeto, documentei a API com Swagger, implementei o motor de cálculo de fatias (`SliceCalculator`), o sistema de permissões por perfil, a autenticação, telas centrais no frontend e componentes reutilizáveis. Por já conhecer React, fui o ponto de referência do time para a tecnologia durante toda a migração.

### Hard Skills

- **React / TypeScript:** Sei fazer com autonomia.
- **Spring Boot:** Sei fazer com autonomia.
- **PostgreSQL:** Sei fazer com autonomia.
- **Docker:** Sei fazer com autonomia.
- **Swagger:** Sei fazer com autonomia.

### Soft Skills

A migração do desktop para a web foi um momento de **transferência de conhecimento** intensa. React e TypeScript eram tecnologias que eu já conhecia, mas eram novidade para o restante do time. Precisei explicar conceitos, revisar código, mostrar padrões e, principalmente, fazer isso sem criar dependência — o objetivo era que o time aprendesse, não que o time me pedisse ajuda para sempre.

Isso desenvolveu minha **didática**: é diferente saber fazer algo e conseguir explicar da forma certa para cada pessoa. Também exercitei minha **adaptabilidade**: Spring Boot era novo para mim, então ao mesmo tempo em que ensinava o que sabia, aprendia o que não sabia ainda. O semestre fortaleceu minha capacidade de **colaboração** em ambos os sentidos — como quem ensina e como quem aprende.

</details>

---

<details>
<summary><strong>4º Semestre — api4</strong></summary>

<br/>

![api4](https://github.com/projetoKhali/api4/assets/108769169/ecda074a-ef3f-4ca5-9cf0-d4b559bcbec5)

- **Empresa:** Oracle
- **Repositório:** [projetoKhali/api4](https://github.com/projetoKhali/api4)

### Missão

Construir uma plataforma de analytics para acompanhamento de parceiros, trilhas, expertises, qualificações e indicadores de produtos, com relatórios e visualizações filtráveis.

### Tecnologias Utilizadas

- **Vue 3 / TypeScript:** Frontend web.
- **Java / Spring Boot:** Backend REST.
- **PostgreSQL:** Banco de dados relacional.
- **Docker:** Containerização.
- **ESLint / Prettier:** Padronização de código.

### Minha participação

Atuei como Scrum Master e como desenvolvedor central do projeto. No front-end Vue construí a maior parte da biblioteca de componentes (`Table`, `Form`, `Filter`, `NotificationPopup`), o roteamento com submenus recursivos e a camada de serviços tipada e paginada. No back-end entreguei paginação, validações e modelagem do banco. Também estabeleci o tooling de qualidade e o setup do monorepo.

### Hard Skills

- **Vue 3 / TypeScript:** Sei fazer com autonomia.
- **Spring Boot:** Sei fazer com autonomia.
- **Scrum Master:** Sei fazer com autonomia.
- **ESLint / Prettier:** Sei fazer com autonomia.

### Soft Skills

Exercer o papel de Scrum Master foi o maior aprendizado desse semestre. Além de continuar desenvolvendo, precisei acompanhar o progresso de todos os integrantes, identificar quem estava travado e intervir antes que o problema atrasasse a sprint. Isso exigiu **atenção e escuta ativa**: perceber quando alguém estava com dificuldade às vezes significava ler entre as linhas, não esperar que a pessoa pedisse ajuda.

Também precisei desenvolver minha **organização de processos**: garantir que as histórias estivessem bem definidas antes da sprint começar, que os critérios de aceitação fossem claros e que o time soubesse o que precisava entregar. A **responsabilidade** de coordenar entregas para um cliente real como a Oracle aumentou minha maturidade em relação a prazos e qualidade. Aprendi que liderar tecnicamente e liderar um processo são habilidades distintas — e que a segunda é, em muitos aspectos, mais difícil.

</details>

---

<details>
<summary><strong>5º Semestre — api5</strong></summary>

<br/>

![api5](https://github.com/user-attachments/assets/94aecab2-e751-4ab4-a2a8-1a6589b4eb01)

- **Empresa:** Pro4tech
- **Repositório:** [projetoKhali/api5](https://github.com/projetoKhali/api5)

### Missão

Centralizar e visualizar dados de recrutamento e seleção em um dashboard interativo, apoiando decisões estratégicas com métricas, filtros, relatórios e gestão de acesso.

### Tecnologias Utilizadas

- **Go:** Backend e data warehouse.
- **React Native / TypeScript:** Frontend mobile.
- **PostgreSQL:** Banco de dados e data warehouse.
- **Docker / CI/CD:** Infraestrutura e automação.
- **Jest:** Testes de frontend.

### Minha participação

Concentrei minha atuação no back-end e na modelagem dimensional do data warehouse, nas consultas analíticas do dashboard, nos testes de integração com containers dedicados e na infraestrutura do monorepo. Convenci o time a adotar Go como linguagem do back-end e atuei como referência técnica para a equipe durante toda a transição.

### Hard Skills

- **Go:** Sei fazer com autonomia.
- **PostgreSQL / ETL:** Sei fazer com autonomia.
- **Docker / CI/CD:** Sei fazer com autonomia.
- **Testes de integração:** Sei fazer com autonomia.

### Soft Skills

A decisão de usar Go foi uma que eu propus ao time e precisei defender com argumentos concretos. Desenvolvi minha **capacidade de persuasão técnica**: não é suficiente estar convicto de que algo é melhor — é preciso construir o caso com clareza, entender as preocupações dos colegas e dar espaço para que a decisão seja coletiva. O time topou, e isso trouxe uma responsabilidade adicional: eu precisava garantir que essa escolha não prejudicaria ninguém.

Isso me levou a investir mais em **compartilhamento de conhecimento**: workshops informais, revisões de PR mais detalhadas, disponibilidade para tirar dúvidas. À medida que o projeto avançava e o time ganhava confiança com Go, percebi que essa mentoria foi diminuindo naturalmente — o que era exatamente o objetivo. Esse processo me ensinou que **transferir conhecimento bem** significa tornar a si mesmo progressivamente dispensável, e que isso é um sucesso, não uma perda.

</details>

---

<details>
<summary><strong>6º Semestre — api6</strong></summary>

<br/>

![api6](https://github.com/user-attachments/assets/d0217c10-db11-470b-a029-f8b664cf4cd2)

- **Empresa:** Kersys
- **Repositório:** [projetoKhali/api6](https://github.com/projetoKhali/api6)

### Missão

Desenvolver uma plataforma para monitorar e planejar reflorestamento, com dashboards, simulador de cenários, previsões com machine learning e gestão de usuários.

### Tecnologias Utilizadas

- **React / TypeScript:** Frontend web.
- **Python / Flask:** API principal.
- **Rust / Actix Web:** Microsserviço de autenticação.
- **PostgreSQL / MongoDB:** Persistência relacional e documental.
- **Nx:** Orquestrador de monorepo.
- **Docker / CI/CD:** Containerização e automação.

### Minha participação

Atuei como Product Owner do semestre final, papel que exigiu de mim mais do que qualquer outro ao longo do curso. Tecnicamente, fui responsável pelo serviço de autenticação em Rust (JWT, criptografia Fernet, revogação de tokens, middleware, logging), pela modelagem dos bancos PostgreSQL e MongoDB, pela infraestrutura do monorepo Nx, pelo pipeline de CI/CD e pela containerização de todos os serviços. Também contribuí na API Python/Flask e no frontend React, estando presente em todas as camadas do sistema.

A iniciativa de usar o Nx como orquestrador de monorepo partiu de mim desde o início: foi uma decisão arquitetural que permitiu separar bem os microsserviços enquanto mantinha o repositório gerenciável. E a inclusão de Rust no projeto foi algo que planejei desde que entrei na FATEC — vi nesse semestre a oportunidade perfeita para aplicá-lo em uma parte isolada e bem delimitada do sistema, e trouxe isso para o time como proposta, que foi aceita.

### Hard Skills

- **Rust / Actix Web:** Sei fazer com autonomia.
- **Flask / Python:** Sei fazer com autonomia.
- **React / TypeScript:** Sei fazer com autonomia.
- **PostgreSQL / MongoDB:** Sei fazer com autonomia.
- **Nx (monorepo):** Sei fazer com autonomia.
- **Docker / CI/CD:** Sei fazer com autonomia.
- **Arquitetura de sistemas:** Sei fazer com autonomia.

### Soft Skills

O semestre mais desafiador do curso foi também o que mais me ensinou sobre mim mesmo — especialmente sobre os limites que ainda tenho.

Fui escolhido como Product Owner porque a equipe e os professores avaliaram que meu conhecimento técnico seria uma vantagem para traduzir a linguagem dos desenvolvedores para o cliente e vice-versa. Essa lógica faz sentido no papel: um PO que entende o que está sendo construído consegue negociar escopo com mais precisão e articular requisitos com mais clareza. Na prática, essa **ponte entre mundos técnico e de negócio** funcionou bem — eu conseguia explicar para o cliente Kersys por que determinada funcionalidade era complexa, ou defender para o time por que a demanda do cliente era legítima.

Mas o papel de PO vai muito além disso, e eu precisei encarar minhas limitações com **honestidade**. As responsabilidades do dia a dia do Product Owner — acompanhar o backlog, organizar as reuniões de refinamento, manter a comunicação constante com o cliente e priorizar histórias — eram tarefas que consumiam tempo de uma forma que eu não estava acostumado. Sou, por natureza, uma pessoa de execução técnica: prefiro resolver um problema no código do que gerenciar o processo de decidir qual problema resolver. Reconhecer isso foi necessário, e buscar ajuda foi a decisão certa.

Foi aqui que a **colaboração** com Marcos Malaquias, um dos integrantes do time, fez toda a diferença. Marcos tinha uma habilidade natural para as responsabilidades processuais do papel de PO que eu não tinha. Dividimos o trabalho de forma complementar: eu trazia a visão técnica nas interações com o cliente e nas decisões de escopo, e Marcos sustentava o processo do dia a dia, as cerimônias e a organização do backlog. Não foi uma divisão planejada desde o começo — emergiu da necessidade e da **confiança mútua** que construímos ao longo do projeto. Esse tipo de parceria, em que duas pessoas reconhecem pontos cegos um do outro e trabalham em complemento, foi uma das lições mais valiosas que o curso me deu.

Ao mesmo tempo, não me afastei do código. Criar o monorepo com Nx, arquitetar o serviço de autenticação em Rust, modelar os bancos — essas decisões partiam de mim e eu as executava. Isso criou um equilíbrio que funcionou para o projeto, mas que também me ensinou algo sobre **autoconhecimento profissional**: entendi onde estou mais forte e onde preciso crescer. Ser um bom engenheiro e ser um bom gestor de produto são carreiras que se tocam mas não são a mesma. Essa distinção ficou muito clara nesse semestre.

Por fim, ver o sistema funcionar com Rust na stack — algo que eu queria desde o primeiro dia de aula — foi uma satisfação que vai além da tecnologia. Foi a demonstração de que é possível, mesmo dentro de um projeto acadêmico com restrições, defender uma escolha técnica com argumentos e viabilizá-la com o apoio do time.

</details>

---

## Próximos Passos

<div align="justify">

Concluo a FATEC com uma formação que já estou aplicando na prática: atuando como desenvolvedor backend na Gorila Invest, onde trabalho com sistemas financeiros em produção.

O que quero continuar desenvolvendo é minha capacidade de atuar em sistemas distribuídos de maior escala — arquitetura, observabilidade, modelagem de dados e integração entre serviços. Rust é uma linguagem em que pretendo me aprofundar. E o lado de gestão e comunicação com produto que o papel de PO no último semestre me expôs é algo que quero continuar desenvolvendo, com mais consciência dos meus pontos cegos.

A FATEC me deu o que eu vim buscar: experiência real de colaboração, a formação oficial e as lacunas preenchidas. O próximo passo é consolidar isso no mercado.

</div>
