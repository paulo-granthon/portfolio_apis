import Project from "../components/project";
import User from "../components/user";
import { ProjectSchema } from "../schemas/project";
import { UserSchema } from "../schemas/user";
import * as styles from "../styles/portfolio";

export default function Portfolio(user: UserSchema) {
  const projects: ProjectSchema[] = [
    {
      id: 1,
      name: "Project 1",
      semester: 1,
      summary: "This is a summary",
      company: "Company 1",
      url: "https://example.com",
    },
  ];

  return (
    <div {...styles.portfolio}>
      <div>
        <h1>Portfolio</h1>
      </div>
      <User user={user} />
      <div {...styles.projects}>
        <h2 {...styles.projectsHeader}>Projects</h2>
        <div>
          {projects.map((project) => (
            <Project key={project.name} project={project} />
          ))}
        </div>
      </div>
    </div>
  );
}
