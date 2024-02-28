import { ProjectSchema } from "../schemas/project";

export async function getProjects() {
  return fetch(import.meta.env.VITE_API_URL + "/projects")
    .then((response) => response.json())
    .then((data) => data as ProjectSchema[]);
}
