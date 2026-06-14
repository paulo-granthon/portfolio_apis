import Project from '../components/project';
import { PortfolioProjectSchema } from '../schemas/portfolio';
import * as styles from '../styles/portfolio';

interface ProjectListProps {
  projects: PortfolioProjectSchema[];
}

export default function ProjectList({ projects }: ProjectListProps) {
  return (
    <div {...styles.projects}>
      <h2 {...styles.projectsHeader}>
        projetos · {projects.length} semestres
      </h2>
      <div>
        {projects.map(project => (
          <Project key={project.name} project={project} />
        ))}
      </div>
    </div>
  );
}
