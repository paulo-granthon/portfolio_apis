import { useEffect, useState } from 'react';
import Project from '../components/project';
import { ProjectSchema } from '../schemas/project';
import { UserSchema } from '../schemas/user';
import * as styles from '../styles/portfolio';
import { getProjectsOfUser } from '../services/project';

interface ProjectListProps {
  user: UserSchema;
}

export default function ProjectList({ user }: ProjectListProps) {
  const [projects, setProjects] = useState<ProjectSchema[]>([]);

  useEffect(() => {
    getProjectsOfUser(user.id).then(projects => setProjects(projects));
  }, [user]);

  return (
    <div {...styles.projects}>
      <h2 {...styles.projectsHeader}>Projects</h2>
      <div>
        {projects.map(project => (
          <Project //
            key={project.name}
            project={project}
            userId={user.id}
          />
        ))}
      </div>
    </div>
  );
}
