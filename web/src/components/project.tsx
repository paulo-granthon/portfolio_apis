import { ProjectSchema } from "../schemas/project";

interface ProjectProps {
  project: ProjectSchema;
}

export default function Project({ project }: ProjectProps) {
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
