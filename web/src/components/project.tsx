
import NoteList from "./noteList";
import { ProjectSchema } from "../schemas/project";
import * as styles from "../styles/project";

interface ProjectProps {
  userId: number;
  project: ProjectSchema;
}

export default function Project({
    userId,
    project,
}: ProjectProps) {
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
        <p>{project.summary}</p>
        <a href={`https://${project.url}`}>{project.url}</a>
      </div>
      <p>{project.description}</p>
      <NoteList projectId={project.id} userId={userId} />
    </div>
  );
}
