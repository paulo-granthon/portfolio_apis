import { useParams } from 'react-router-dom';
import { useEffect, useState } from 'react';
import User from '../components/user';
import ProjectList from '../components/projectList';
import { PortfolioSchema } from '../schemas/portfolio';
import { getPortfolio, portfolioMarkdownUrl } from '../services/portfolio';
import { useScrollScenes } from '../theme/useScrollTheme';
import { applyPalette } from '../theme/palettes';
import { BackgroundFx } from '../theme/BackgroundFx';
import * as styles from '../styles/portfolio';

export default function Portfolio() {
  const [portfolio, setPortfolio] = useState<PortfolioSchema | undefined>();
  const [scrollEl, setScrollEl] = useState<HTMLElement | null>(null);

  const params = useParams<{ userId: string }>();

  useEffect(() => {
    const id = params.userId ? parseInt(params.userId, 10) : undefined;
    if (!id) return;

    getPortfolio(id).then(portfolio => setPortfolio(portfolio));
  }, [params.userId]);

  const { active, current, next, pct } = useScrollScenes(
    portfolio?.projects.length,
    scrollEl,
  );
  useEffect(() => applyPalette(active), [active]);

  if (!portfolio) {
    return <p>Usuário não encontrado</p>;
  }

  return (
    <>
      <BackgroundFx current={current} next={next} pct={pct} />
      <div {...styles.layout}>
        <aside {...styles.sidebar}>
          <header {...styles.portfolio}>
            <p {...styles.kicker}>// portfólio · banco de dados · fatec-sjc</p>
            <div {...styles.headerRow}>
              <h1 {...styles.headerTitle}>Portfólio</h1>
              <a
                {...styles.downloadButton}
                href={portfolioMarkdownUrl(portfolio.user.id)}
                download
              >
                › baixar.md
              </a>
            </div>
          </header>
          <User user={portfolio.user} compact />
        </aside>
        <main {...styles.scrollPane} ref={setScrollEl as (el: HTMLElement | null) => void}>
          <ProjectList
            projects={portfolio.projects}
            githubUsername={portfolio.user.githubUsername ?? undefined}
          />
        </main>
      </div>
    </>
  );
}
