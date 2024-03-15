import { useEffect, useState } from "react";
import Note from "../components/note";
import { NoteSchema } from "../schemas/note";
import * as styles from "../styles/note";
import { getNotesOfUserProject } from "../services/note";

interface NoteListProps {
  projectId: number;
  userId: number;
}

export default function NoteList({ projectId, userId }: NoteListProps) {
  const [notes, setNotes] = useState<NoteSchema[]>([]);

  useEffect(() => {
    getNotesOfUserProject(userId, projectId).then(notes => setNotes(notes));
  }, [projectId, userId]);

  return (
    <div {...styles.notes}>
      {notes.map((note) => (
        <Note key={note.id} note={note} />
      ))}
    </div>
  );
}
