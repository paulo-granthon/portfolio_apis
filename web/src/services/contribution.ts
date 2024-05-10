import { ContributionSchema } from "../schemas/contribution";

const API_URL = import.meta.env.VITE_API_URL;

/* eslint-disable  @typescript-eslint/no-explicit-any */
export function mapContributions(data: any): ContributionSchema[] {
  if (!data || !!data.Error) return [];
  return data
    .map((contribution: any) => mapContribution(contribution))
    .filter((contribution: ContributionSchema | undefined) => !!contribution);
}

/* eslint-disable  @typescript-eslint/no-explicit-any */
export function mapContribution(data: any): ContributionSchema | undefined {
  if (!data) return undefined;
  return {
    id: data.id,
    title: data.title,
    projectId: data.projectId,
    userId: data.userId,
    content: data.content,
    skills: data.skills,
  };
}

export async function getContributionsOfUserProject(
  userId: number,
  projectId: number,
): Promise<ContributionSchema[]> {
  if (userId === 0 || projectId === 0) {
    return [];
  }
  return fetch(`${API_URL}/contributions?user=${userId}&project=${projectId}`)
    .then((response) => response.json())
    .then((data) => mapContributions(data));
}
