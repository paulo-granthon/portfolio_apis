import * as stylex from "@stylexjs/stylex";

const styles = stylex.create({
  project: {
    gap: 0,
    padding: 0,
  },
  projectHeader: {
    width: "100%",
    display: "flex",
    alignItems: "center",
  },
  projectHeaderTitle: {
    margin: "0.875rem",
  },
  projectHeaderExtra: {
    margin: "auto",
    width: '50%',
    display: "flex",
    justifyContent: "end",
    textAlign: "center",
  },
  projectHeaderExtraItem: {
    margin: "auto",
  },
  projectSubHeader: {
      margin: 0,
      padding: 0,
  },
});

export const project = stylex.props(styles.project);
export const projectHeader = stylex.props(styles.projectHeader);
export const projectSubHeader = stylex.props(styles.projectSubHeader);
export const projectHeaderTitle = stylex.props(styles.projectHeaderTitle);
export const projectHeaderExtra = stylex.props(styles.projectHeaderExtra);
export const projectHeaderExtraItem = stylex.props(
  styles.projectHeaderExtraItem,
);
