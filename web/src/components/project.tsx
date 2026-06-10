import ContributionList from './contributionList';
import { PortfolioProjectSchema } from '../schemas/portfolio';
import * as styles from '../styles/project';

interface ProjectProps {
  project: PortfolioProjectSchema;
}

export default function Project({ project }: ProjectProps) {
  return (
    <div {...styles.project}>
      <div {...styles.projectImageContainer}>
        <img
          {...styles.projectImage}
          src={project.image}
          alt={`${project.name} Image`}
        />
      </div>
      <div {...styles.projectHeader}>
        <h3 {...styles.projectHeaderTitle}>{project.name}</h3>
        <div {...styles.projectHeaderExtra}>
          <p {...styles.projectHeaderExtraItem}>
            <b {...styles.base.hidden}>Empresa: </b>
            {project.company}
          </p>
          <p {...styles.projectHeaderExtraItem}>
            <b {...styles.base.hidden}>Semestre: </b>
            {project.semester}º Sem
          </p>
        </div>
      </div>
      <div {...styles.projectSubHeader}>
        <p>{project.summary}</p>
        <a href={`https://${project.url}`}>{project.url}</a>
      </div>
      <p>{project.description}</p>
      <ContributionList contributions={project.contributions} />
    </div>
  );
}
