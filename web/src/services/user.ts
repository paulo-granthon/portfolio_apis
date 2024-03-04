import {
  UserSchema,
  RegisterUserSchema,
  PostUserSchema,
  UpdateUserSchema,
} from "../schemas/user";

import { YearSemester } from "../schemas/yearSemester";

const API_URL = import.meta.env.VITE_API_URL;

/* eslint-disable  @typescript-eslint/no-explicit-any */
export function mapUsers(data: any): UserSchema[] {
  return data.map((user: any) => mapUser(user));
}

/* eslint-disable  @typescript-eslint/no-explicit-any */
export function mapUser(data: any): UserSchema {
  return {
    id: data.id,
    name: data.name,
    summary: data.summary,
    semesterMatriculed: new YearSemester(
      data.semesterMatriculed.year,
      data.semesterMatriculed.semester,
    ),
    githubUsername: data.githubUsername,
  };
}

export async function getUsers(): Promise<UserSchema[]> {
  return fetch(API_URL + "/users")
    .then((response) => response.json())
    .then((data) => mapUsers(data));
}

export async function getUser(id: number): Promise<UserSchema> {
  return fetch(API_URL + "/users/" + id)
    .then((response) => response.json())
    .then((data) => mapUser(data));
}

export async function registerUser(payload: RegisterUserSchema) {
  return fetch(API_URL + "/register", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(payload),
  });
}

export async function createUser(user: PostUserSchema) {
  return fetch(API_URL + "/users", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(user),
  });
}

export async function updateUser(id: number, user: UpdateUserSchema) {
  return fetch(API_URL + "/users/" + id, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(user),
  });
}

export async function deleteUser(id: number) {
  return fetch(API_URL + "/users/" + id, {
    method: "DELETE",
  });
}
