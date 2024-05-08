import { ContributionSchema } from "../schemas/contribution";
import * as styles from "../styles/contribution";

interface ContributionProps {
  contribution: ContributionSchema;
}

export default function Contribution({ contribution }: ContributionProps) {
  return (
    <div {...styles.contribution}>
      {contribution.skill ? (
        <p {...styles.title}>{contribution.skill}</p>
      ) : (
        <p {...styles.title}>[Skill Undefined]</p>
      )}

      <p {...styles.content}>{contribution.content}</p>
    </div>
  );
}
