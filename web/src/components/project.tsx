import ContributionList from './contributionList';
import { ContributionTimeline } from './contributionTimeline';
import { PortfolioProjectSchema } from '../schemas/portfolio';
import { renderRichText } from './richText';
import * as styles from '../styles/project';

interface ProjectProps {
  project: PortfolioProjectSchema;
  githubUsername?: string;
}

export default function Project({ project, githubUsername }: ProjectProps) {
  return (
    <article {...styles.project} data-project={project.name}>
      {project.image && (
        <div {...styles.bannerSticky}>
          <img
            {...styles.bannerImg}
            src={project.image}
            alt={`${project.name} banner`}
          />
        </div>
      )}
      <div {...styles.projectBody}>
        <span {...styles.semesterGhost}>{project.semester}</span>

        <div {...styles.projectHeader}>
          <h3 {...styles.projectHeaderTitle}>{project.name}</h3>
          <div {...styles.projectHeaderExtra}>
            <p {...styles.projectHeaderExtraItem}>{project.company}</p>
            <p {...styles.projectHeaderExtraItem}>
              {project.semester}º semestre
            </p>
          </div>
        </div>

        <div {...styles.projectSubHeader}>
          <p {...styles.projectSummary}>{project.summary}</p>
          <a
            {...styles.projectRepo}
            href={`https://${project.url}`}
            target="_blank"
            rel="noreferrer"
          >
            {project.url}
          </a>
        </div>

        <p {...styles.projectDescription}>
          {renderRichText(project.description)}
        </p>

        {project.participation && (
          <div {...styles.participation}>
            <h4 {...styles.participationTitle}>Minha participação</h4>
            <p {...styles.participationText}>
              {renderRichText(project.participation)}
            </p>
          </div>
        )}

        <ContributionList contributions={project.contributions} />

        {githubUsername && project.url && (
          <ContributionTimeline repoUrl={project.url} author={githubUsername} />
        )}
      </div>
    </article>
  );
}
