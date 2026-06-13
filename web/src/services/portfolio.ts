import { PortfolioSchema } from '../schemas/portfolio';
import { mapContributions } from './contribution';
import { mapUser } from './user';

const API_URL = import.meta.env.VITE_API_URL;

/* eslint-disable  @typescript-eslint/no-explicit-any */
export function mapPortfolio(data: any): PortfolioSchema | undefined {
  if (!data || !!data.Error) return undefined;

  const user = mapUser(data.user);
  if (!user) return undefined;

  return {
    user,
    projects: (data.projects ?? []).map((project: any) => ({
      id: project.id,
      name: project.name,
      image: project.image,
      semester: project.semester,
      company: project.company,
      summary: project.summary,
      description: project.description,
      url: project.url,
      contributions: mapContributions(project.contributions),
    })),
  };
}

export async function getPortfolio(
  userId: number,
): Promise<PortfolioSchema | undefined> {
  if (!userId) return undefined;
  return fetch(`${API_URL}/portfolio/${userId}`)
    .then(response => response.json())
    .then(data => mapPortfolio(data));
}

export function portfolioMarkdownUrl(userId: number): string {
  return `${API_URL}/portfolio/${userId}/markdown`;
}
