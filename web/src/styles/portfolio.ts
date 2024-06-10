import * as stylex from "@stylexjs/stylex";
import * as base from "./base";
export { base };

const styles = stylex.create({
  portfolio: {
    margin: "auto",
    width: "80%",
  },
  projects: {
    display: "flex",
    flexFlow: "column",
    justifyContent: "center",
    width: "80%",
    margin: 'auto',
  },
  projectsHeader: {
    margin: "0",
    textAlign: "center",
    fontSize: "2rem",
  },
});

export const portfolio = stylex.props(styles.portfolio);
export const projects = stylex.props(styles.projects);
export const projectsHeader = stylex.props(styles.projectsHeader);
