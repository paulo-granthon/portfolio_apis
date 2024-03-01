import {
    UserSchema,
    RegisterUserSchema,
    PostUserSchema,
    UpdateUserSchema,
} from "../schemas/user";

const API_URL = import.meta.env.VITE_API_URL;

export async function getUsers(): Promise<UserSchema[]> {
  return fetch(API_URL + "/users")
    .then((response) => response.json())
    .then((data) => data);
}

export async function getUser(id: number): Promise<UserSchema> {
  return fetch(API_URL + "/users/" + id)
    .then((response) => response.json())
    .then((data) => data);
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
