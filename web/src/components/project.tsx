import { ProjectSchema } from "../schemas/project";

export default function Project(project: ProjectSchema) {
  return (
    <div>
      <div>
        <p>{project.name}</p>
        <p>{project.semester}º Sem</p>
      </div>
      <p>{project.url}</p>
      <p>{project.company}</p>
      <p>{project.summary}</p>
    </div>
  );
}
