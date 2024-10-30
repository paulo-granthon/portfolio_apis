import { ContributionSchema } from '../schemas/contribution';
import * as styles from '../styles/contribution';

interface ContributionProps {
  contribution: ContributionSchema;
}

export default function Contribution({ contribution }: ContributionProps) {
  return (
    <div {...styles.contribution}>
      <div {...styles.contributionHeader}>
        {contribution.title ? (
          <h4 {...styles.title}>{contribution.title}</h4>
        ) : (
          <h4 {...styles.title}>[Untitled]</h4>
        )}
        <p {...styles.skills}>
          <b {...styles.base.hidden}>Conhecimentos exercitados: </b>
          {!!contribution.skills && !!contribution.skills.length ? (
            <>
              {contribution.skills?.map((skill, index) => (
                <code {...styles.skill} key={index}>
                  {skill + ' '}
                </code>
              ))}
            </>
          ) : (
            <code {...styles.skill}>[Nenhum especificado]</code>
          )}
        </p>
      </div>
      <p {...styles.content}>{contribution.content}</p>
    </div>
  );
}
