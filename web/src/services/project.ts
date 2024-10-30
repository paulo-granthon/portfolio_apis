import {
  PostProjectSchema,
  ProjectSchema,
  UpdateProjectSchema,
} from '../schemas/project';

const API_URL = import.meta.env.VITE_API_URL;

export async function getProjects(): Promise<ProjectSchema[]> {
  return fetch(API_URL + '/projects')
    .then(response => response.json())
    .then(data => data);
}

export async function getProjectsOfUser(
  userId: number,
): Promise<ProjectSchema[]> {
  return userId
    ? fetch(API_URL + '/users/' + userId + '/projects')
        .then(response => response.json())
        .then(data => data)
    : undefined;
}

export async function getProject(id: number): Promise<ProjectSchema> {
  return id
    ? fetch(API_URL + '/projects/' + id)
        .then(response => response.json())
        .then(data => data)
    : undefined;
}

export async function createProject(project: PostProjectSchema) {
  return project
    ? fetch(API_URL + '/projects', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(project),
      })
    : undefined;
}

export async function updateProject(id: number, project: UpdateProjectSchema) {
  return !!id && !!project
    ? fetch(API_URL + '/projects/' + id, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(project),
      })
    : undefined;
}

export async function deleteProject(id: number) {
  return id
    ? fetch(API_URL + '/projects/' + id, {
        method: 'DELETE',
      })
    : undefined;
}
