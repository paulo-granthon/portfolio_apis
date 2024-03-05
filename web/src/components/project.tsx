import { ProjectSchema } from "../schemas/project";
import * as styles from "../styles/project";

interface ProjectProps {
  project: ProjectSchema;
}

export default function Project({ project }: ProjectProps) {
  return (
    <div {...styles.project}>
      <div {...styles.projectHeader}>
        <p>{project.name}</p>
        <p>{project.semester}º Sem</p>
      </div>
      <div {...styles.projectSubHeader}>
          <p>{project.url}</p>
          <p>{project.company}</p>
      </div>
      <p>{project.summary}</p>
    </div>
  );
}
