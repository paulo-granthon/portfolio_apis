import * as stylex from "@stylexjs/stylex";

const profilePictureSize: string | number = "10vw";

const styles = stylex.create({
  card: {
    display: "flex",
    flexFlow: "row",
    gap: "1.25vw",
  },
  cardLeft: {
    margin: "1.5vw",
  },
  cardLeftPicture: {
    width: profilePictureSize,
    height: profilePictureSize,
    borderRadius: "50%",
  },
  cardRight: {
    display: "flex",
    flexFlow: "column",
    width: "100%",
  },
  cardRightHeader: {
    display: "flex",
    flexFlow: "row",
    justifyContent: "space-between",
    alignItems: "center",
    width: "100%",
  },
  cardRightSemester: {
    display: "flex",
    flexFlow: "row",
    justifyContent: "space-between",
    alignItems: "center",
    width: "50%",
  },
  cardRightSummary: {
    margin: "1.5vw",
  },
});

export const card = stylex.props(styles.card);
export const cardLeft = stylex.props(styles.cardLeft);
export const cardLeftPicture = stylex.props(styles.cardLeftPicture);
export const cardRight = stylex.props(styles.cardRight);
export const cardRightHeader = stylex.props(styles.cardRightHeader);
export const cardRightSemester = stylex.props(styles.cardRightSemester);
export const cardRightSummary = stylex.props(styles.cardRightSummary);
