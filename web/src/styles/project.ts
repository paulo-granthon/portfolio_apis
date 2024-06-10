import * as stylex from "@stylexjs/stylex";
import * as base from "./base";
export { base };

const styles = stylex.create({
  project: {
    gap: 0,
    width: "100%",
    border: "1px solid black",
    borderRadius: "8px",
    padding: "0.5em",
    margin: "0.5em",
  },
  projectHeader: {
    width: "100%",
    display: "flex",
    alignItems: "center",
    justifyContent: "space-between",
  },
  projectHeaderTitle: {
    margin: "0.875rem",
  },
  projectHeaderExtra: {
    width: "50%",
    display: "flex",
    justifyContent: "end",
    textAlign: "center",
  },
  projectHeaderHiddenLabel: {
    display: "none",
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
export const projectHeaderHiddenLabel = stylex.props(
  styles.projectHeaderHiddenLabel,
);
