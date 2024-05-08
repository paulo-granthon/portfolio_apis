export interface ContributionSchema {
  id: number;
  skill: string;
  projectId: number;
  userId: number;
  content: string;
}

export type PostContributionSchema = Omit<ContributionSchema, "id">;

export type UpdateContributionSchema = Omit<ContributionSchema, "id">;
