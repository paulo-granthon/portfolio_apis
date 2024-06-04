import { useEffect, useState } from "react";
import Contribution from "../components/contribution";
import { ContributionSchema } from "../schemas/contribution";
import * as styles from "../styles/contribution";
import { getContributionsOfUserProject } from "../services/contribution";

interface ContributionListProps {
  projectId: number;
  userId: number;
}

export default function ContributionList({
  projectId,
  userId,
}: ContributionListProps) {
  const [contributions, setContributions] = useState<ContributionSchema[]>([]);

  useEffect(() => {
    getContributionsOfUserProject(userId, projectId).then((contributions) =>
      setContributions(contributions),
    );
  }, [projectId, userId]);

  return (
    <div {...styles.contributions}>
      <h3 {...styles.contributionsTitle}>Contribuições</h3>
      <div {...styles.contributionList}>
        {contributions.map((contribution) => (
          <Contribution key={contribution.id} contribution={contribution} />
        ))}
      </div>
    </div>
  );
}
