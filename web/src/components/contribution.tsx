import { ContributionSchema } from "../schemas/contribution";
import * as styles from "../styles/contribution";

interface ContributionProps {
  contribution: ContributionSchema;
}

export default function Contribution({ contribution }: ContributionProps) {
  return (
    <div {...styles.contribution}>
      {contribution.title ? (
        <p {...styles.title}>{contribution.title}</p>
      ) : (
        <p {...styles.title}>[Untitled]</p>
      )}
      <p {...styles.content}>{contribution.content}</p>
    </div>
  );
}
