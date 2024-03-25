import { NoteSchema } from "../schemas/note";

const API_URL = import.meta.env.VITE_API_URL;

/* eslint-disable  @typescript-eslint/no-explicit-any */
export function mapNotes(data: any): NoteSchema[] {
  return data
    .map((note: any) => mapNote(note))
    .filter((note: NoteSchema | undefined) => !!note);
}

/* eslint-disable  @typescript-eslint/no-explicit-any */
export function mapNote(data: any): NoteSchema | undefined {
  if (!data) return undefined;
  return {
    id: data.id,
    skill: data.skill,
    projectId: data.projectId,
    userId: data.userId,
    content: data.content,
  };
}

export async function getNotesOfUserProject(
  userId: number,
  projectId: number,
): Promise<NoteSchema[]> {
  if (userId === 0 || projectId === 0) {
    return [];
  }
  return fetch(`${API_URL}/notes?user=${userId}&project=${projectId}`)
    .then((response) => response.json())
    .then((data) => mapNotes(data));
}
