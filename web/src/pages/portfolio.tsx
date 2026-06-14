import { useParams } from 'react-router-dom';
import { useEffect, useState } from 'react';
import User from '../components/user';
import ProjectList from '../components/projectList';
import { PortfolioSchema } from '../schemas/portfolio';
import { getPortfolio, portfolioMarkdownUrl } from '../services/portfolio';
import * as styles from '../styles/portfolio';

export default function Portfolio() {
  const [portfolio, setPortfolio] = useState<PortfolioSchema | undefined>();

  const params = useParams<{ userId: string }>();

  useEffect(() => {
    const id = params.userId ? parseInt(params.userId, 10) : undefined;
    if (!id) return;

    getPortfolio(id).then(portfolio => setPortfolio(portfolio));
  }, [params.userId]);

  return (
    <>
      {portfolio ? (
        <div>
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
          <User user={portfolio.user} />
          <ProjectList projects={portfolio.projects} />
        </div>
      ) : (
        <p>Usuário não encontrado</p>
      )}
    </>
  );
}
