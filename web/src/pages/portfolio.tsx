import Project from "../components/project";
import User from "../components/user";
import { ProjectSchema } from "../schemas/project";
import { UserSchema } from "../schemas/user";

export default function Portfolio(user: UserSchema) {
  const projects: ProjectSchema[] = [
    {
      id: 1,
      name: "Project 1",
      semester: 1,
      summary: "This is a summary",
      company: "Company 1",
      url: "https://example.com",
    },
  ];

  return (
    <div>
      <div>
        <h1>Portfolio</h1>
      </div>
      <User user={user} />
      <div>
        {projects.map((project) => (
          <Project key={project.name} project={project} />
        ))}
      </div>
    </div>
  );
}
