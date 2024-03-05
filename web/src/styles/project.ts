import * as stylex from "@stylexjs/stylex";

const styles = stylex.create({
  project: {
    gap: 0,
    padding: 0,
  },
  projectHeader: {
    display: "flex",
    justifyContent: "space-around",
  },
  projectSubHeader: {
    display: "flex",
    justifyContent: "space-around",
  },
});

export const project = stylex.props(styles.project);
export const projectHeader = stylex.props(styles.projectHeader);
export const projectSubHeader = stylex.props(styles.projectSubHeader);
