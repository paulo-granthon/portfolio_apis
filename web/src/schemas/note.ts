export interface NoteSchema {
  id: number;
  skill: string;
  projectId: number;
  userId: number;
  content: string;
}

export type PostNoteSchema = Omit<NoteSchema, "id">;

export type UpdateNoteSchema = Omit<NoteSchema, "id">;
