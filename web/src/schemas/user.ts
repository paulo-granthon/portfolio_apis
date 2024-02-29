export type YearSemester = {
  year: number;
  semester: 1 | 2;
};

export interface UserSchema {
  id: number;
  name: string;
  summary: string;
  semesterMatriculated: YearSemester;
  githubUsername: string;
}

export type PostUserSchema = Partial<Omit<UserSchema, "id">> &
  Pick<UserSchema, "name"> & {
    password: string;
  };

export type UpdateUserSchema = Omit<UserSchema, "id">;
