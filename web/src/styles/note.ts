import * as stylex from "@stylexjs/stylex";

const styles = stylex.create({
  notes: {
    color: "white",
  },
  note: {
    display: "flex",
    gap: "4rem",
  },
  title: {
    fontWeight: "bold",
  },
  content: {
    color: "white",
  },
});

export const notes = stylex.props(styles.notes);
export const note = stylex.props(styles.note);
export const title = stylex.props(styles.title);
export const content = stylex.props(styles.content);
