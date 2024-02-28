import { ProjectSchema } from "../schemas/project";

const API_URL = import.meta.env.VITE_API_URL;

export async function getProjects(): Promise<ProjectSchema[]> {
  return fetch(API_URL + "/projects")
    .then((response) => response.json())
    .then((data) => data);
}

export async function getProject(id: string): Promise<ProjectSchema> {
  return fetch(API_URL + "/projects/" + id)
    .then((response) => response.json())
    .then((data) => data);
}

export async function createProject(
  project: ProjectSchema,
): Promise<ProjectSchema> {
  return fetch(API_URL + "/projects", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(project),
  })
    .then((response) => response.json())
    .then((data) => data);
}

export async function updateProject(
  id: string,
  project: ProjectSchema,
): Promise<ProjectSchema> {
  return fetch(API_URL + "/projects/" + id, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(project),
  })
    .then((response) => response.json())
    .then((data) => data);
}

export async function deleteProject(id: string): Promise<void> {
  return fetch(API_URL + "/projects/" + id, {
    method: "DELETE",
  }).then((response) => response.json());
}
