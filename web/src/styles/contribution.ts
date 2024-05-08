import * as stylex from "@stylexjs/stylex";

const styles = stylex.create({
  contributions: {
    color: "white",
  },
  contribution: {
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

export const contributions = stylex.props(styles.contributions);
export const contribution = stylex.props(styles.contribution);
export const title = stylex.props(styles.title);
export const content = stylex.props(styles.content);
