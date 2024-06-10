import * as stylex from "@stylexjs/stylex";

const styles = stylex.create({
  hidden: {
    display: "none",
  },
});

export const hidden = stylex.props(styles.hidden);
