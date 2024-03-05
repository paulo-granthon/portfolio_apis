import * as stylex from "@stylexjs/stylex";

const profilePictureSize: string | number = "18.75vw";

const styles = stylex.create({
  userCard: {
    display: "flex",
    flexFlow: "row",
    gap: "1.25vw",
  },
  userCardLeft: {
    margin: "1.5vw",
  },
  profilePicture: {
    width: profilePictureSize,
    height: profilePictureSize,
    borderRadius: "50%",
  },
  userCardRight: {
    display: "flex",
    flexFlow: "column",
  },
  userCardSemester: {
    display: "flex",
    flexFlow: "row",
  },
});

export const userCard = stylex.props(styles.userCard);
export const userCardLeft = stylex.props(styles.userCardLeft);
export const profilePicture = stylex.props(styles.profilePicture);
export const userCardRight = stylex.props(styles.userCardRight);
export const userCardSemester = stylex.props(styles.userCardSemester);
