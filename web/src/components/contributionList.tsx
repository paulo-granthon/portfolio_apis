import Contribution from '../components/contribution';
import { PortfolioContributionSchema } from '../schemas/portfolio';
import * as styles from '../styles/contribution';

interface ContributionListProps {
  contributions: PortfolioContributionSchema[];
}

export default function ContributionList({
  contributions,
}: ContributionListProps) {
  return (
    <div {...styles.contributions}>
      <h3 {...styles.contributionsTitle}>Contribuições</h3>
      <div {...styles.contributionList}>
        {!!contributions && !!contributions.length ? (
          <>
            {contributions.map(contribution => (
              <Contribution key={contribution.id} contribution={contribution} />
            ))}
          </>
        ) : (
          <p>Nenhuma contribuição especificada</p>
        )}
      </div>
    </div>
  );
}
