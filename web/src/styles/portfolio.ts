import * as stylex from "@stylexjs/stylex";

const styles = stylex.create({
  portfolio: {
    margin: ".5vw",
  },
  projects: {
    display: "flex",
    flexFlow: "column",
    justifyContent: "center",
    maxWidth: "80em",
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
