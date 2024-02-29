import * as stylex from "@stylexjs/stylex";

const styles = stylex.create({
  profilePicture: {
    width: 200,
    height: 200,
    borderRadius: 100,
  },
});

export const profilePicture = stylex.props(styles.profilePicture);
