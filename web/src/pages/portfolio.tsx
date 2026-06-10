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
          <div {...styles.portfolio}>
            <h1>Portfolio</h1>
            <a href={portfolioMarkdownUrl(portfolio.user.id)} download>
              Download Markdown
            </a>
          </div>
          <User user={portfolio.user} />
          <ProjectList projects={portfolio.projects} />
        </div>
      ) : (
        <p>Usuário não encontrado</p>
      )}
    </>
  );
}
