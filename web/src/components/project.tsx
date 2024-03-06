import { ProjectSchema } from "../schemas/project";
import * as styles from "../styles/project";

interface ProjectProps {
  project: ProjectSchema;
}

export default function Project({ project }: ProjectProps) {
  return (
    <div {...styles.project}>
      <div {...styles.projectHeader}>
        <h3 {...styles.projectHeaderTitle}>{project.name}</h3>
        <div {...styles.projectHeaderExtra}>
          <p {...styles.projectHeaderExtraItem}>{project.company}</p>
          <p {...styles.projectHeaderExtraItem}>{project.semester}º Sem</p>
        </div>
      </div>
      <div {...styles.projectSubHeader}>
        <p>{project.url}</p>
      </div>
      <p>{project.summary}</p>
    </div>
  );
}
