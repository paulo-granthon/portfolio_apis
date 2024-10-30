export interface ContributionSchema {
  id: number;
  projectId: number;
  userId: number;
  title: string;
  content: string;
  skills: string[];
}

export type PostContributionSchema = Omit<ContributionSchema, 'id'>;

export type UpdateContributionSchema = Omit<ContributionSchema, 'id'>;
