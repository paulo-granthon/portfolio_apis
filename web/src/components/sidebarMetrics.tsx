import { PortfolioProjectSchema } from '../schemas/portfolio';
import { ContributionTimeline } from './contributionTimeline';
import * as s from '../styles/sidebarMetrics';

interface Props {
  project: PortfolioProjectSchema | null;
  githubUsername?: string;
}

export function SidebarMetrics({ project, githubUsername }: Props) {
  if (!project) {
    return (
      <div {...s.root}>
        <p {...s.placeholder}>// role no projeto para ver métricas</p>
      </div>
    );
  }

  const uniqueSkills = new Set(
    project.contributions.flatMap(c => c.skills),
  ).size;

  return (
    <div {...s.root}>
      <span {...s.ghost}>{project.semester}</span>
      <h3 {...s.name}>{project.name}</h3>
      <hr {...s.divider} />
      <div {...s.stats}>
        <div {...s.statRow}>
          <span {...s.statNum}>{project.contributions.length}</span>
          <span {...s.statLabel}>contribuições</span>
        </div>
        <div {...s.statRow}>
          <span {...s.statNum}>{uniqueSkills}</span>
          <span {...s.statLabel}>habilidades únicas</span>
        </div>
      </div>
      {githubUsername && project.url && (
        <ContributionTimeline repoUrl={project.url} author={githubUsername} />
      )}
    </div>
  );
}
