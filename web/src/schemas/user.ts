import { YearSemester } from './yearSemester';

export interface UserSchema {
  id: number;
  name: string;
  summary: string;
  semesterMatriculed: YearSemester;
  githubUsername: string;
}

export type PostUserSchema = Partial<Omit<UserSchema, 'id'>> &
  Pick<UserSchema, 'name'> & {
    password: string;
  };

export type RegisterUserSchema = Pick<UserSchema, 'name'> & {
  password: string;
};

export type UpdateUserSchema = Omit<UserSchema, 'id'>;
