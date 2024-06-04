import * as stylex from "@stylexjs/stylex";

const styles = stylex.create({
  contributions: {
    display: "flex",
    flexDirection: "column",
    justifyContent: "center",
  },
  contributionsTitle: {
    textAlign: "center",
    margin: 0,
  },
  contributionList: {
    display: "flex",
    flexDirection: "column",
  },
  contribution: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    padding: "0.5em",
    margin: "0.5em",
    gap: 0,
    border: "1px solid black",
    borderRadius: "8px",
  },
  contributionHeader: {
    width: "100%",
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    border: "1px solid black",
    borderRadius: "8px",
    gap: 0,
  },
  title: {
    margin: 0,
    fontWeight: "bold",
  },
  skills: {
    display: "flex",
    flexDirection: "row",
    margin: "0.5rem",
    gap: "4rem",
  },
  skill: {
    borderRadius: "4px",
    height: "fit-content",
  },
  content: {
    color: "black",
  },
});

export const contributions = stylex.props(styles.contributions);
export const contributionsTitle = stylex.props(styles.contributionsTitle);
export const contributionList = stylex.props(styles.contributionList);
export const contribution = stylex.props(styles.contribution);
export const contributionHeader = stylex.props(styles.contributionHeader);
export const title = stylex.props(styles.title);
export const content = stylex.props(styles.content);
export const skills = stylex.props(styles.skills);
export const skill = stylex.props(styles.skill);
