import { ContributionSchema } from "../schemas/contribution";
import * as styles from "../styles/contribution";

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
        <div {...styles.skills}>
          {contribution.skills?.map((skill, index) => (
            <code {...styles.skill} key={index}>
              {skill + " "}
            </code>
          ))}
        </div>
      </div>
      <p {...styles.content}>{contribution.content}</p>
    </div>
  );
}
