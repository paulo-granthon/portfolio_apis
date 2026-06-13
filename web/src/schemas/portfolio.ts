import { ContributionSchema } from './contribution';
import { ProjectSchema } from './project';
import { UserSchema } from './user';

export type PortfolioContributionSchema = Pick<
  ContributionSchema,
  'id' | 'title' | 'content' | 'skills'
>;

export interface PortfolioProjectSchema extends ProjectSchema {
  contributions: PortfolioContributionSchema[];
}

export interface PortfolioSchema {
  user: UserSchema;
  projects: PortfolioProjectSchema[];
}
