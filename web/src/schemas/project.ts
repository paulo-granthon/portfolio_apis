export interface ProjectSchema {
  id: number;
  name: string;
  semester: number;
  company: string;
  summary: string;
  description: string;
  url: string;
}

export type PostProjectSchema = Omit<ProjectSchema, "id">;

export type UpdateProjectSchema = Omit<ProjectSchema, "id">;
