import { NoteSchema } from "../schemas/note";
import * as styles from "../styles/note";

interface NoteProps {
  note: NoteSchema;
}

export default function Note({ note }: NoteProps) {
  return (
    <div {...styles.note}>
      <p {...styles.title}>{note.skill}</p>
      <p {...styles.content}>{note.content}</p>
    </div>
  );
}
