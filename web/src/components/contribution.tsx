import { PortfolioContributionSchema } from '../schemas/portfolio';
import { renderRichText } from './richText';
import * as styles from '../styles/contribution';

interface ContributionProps {
  contribution: PortfolioContributionSchema;
}

export default function Contribution({ contribution }: ContributionProps) {
  return (
    <div {...styles.contribution}>
      <div {...styles.contributionHeader}>
        <h4 {...styles.title}>{contribution.title || '[Sem título]'}</h4>
        {!!contribution.skills?.length && (
          <p {...styles.skills}>
            {contribution.skills.map((skill, index) => (
              <code {...styles.skill} key={index}>
                {skill}
              </code>
            ))}
          </p>
        )}
      </div>
      <p {...styles.content}>{renderRichText(contribution.content)}</p>
    </div>
  );
}
