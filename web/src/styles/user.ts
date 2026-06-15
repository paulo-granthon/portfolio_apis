import * as stylex from '@stylexjs/stylex';
import * as base from './base';
export { base };

const styles = stylex.create({
  card: {
    position: 'relative',
    maxWidth: 'var(--maxw)',
    width: '100%',
    margin: '0 auto',
    boxSizing: 'border-box',
    display: 'flex',
    flexFlow: 'row wrap',
    gap: '28px',
    alignItems: 'center',
    padding: '28px',
    backgroundColor: 'var(--bg-elev)',
    borderWidth: '1px',
    borderStyle: 'solid',
    borderColor: 'var(--border)',
    borderRadius: 'var(--radius)',
    boxShadow: '0 24px 60px -30px rgba(0, 0, 0, 0.8)',
  },
  // gradient ring wrapper around the avatar
  cardLeft: {
    flexShrink: 0,
    padding: '3px',
    borderRadius: '20px',
    backgroundImage: 'linear-gradient(140deg, var(--pink), var(--green))',
    boxShadow: '0 0 34px -6px var(--pink-glow)',
  },
  cardLeftPicture: {
    display: 'block',
    width: '116px',
    height: '116px',
    objectFit: 'cover',
    borderRadius: '17px',
    backgroundColor: 'var(--bg)',
  },
  cardRight: {
    display: 'flex',
    flexFlow: 'column',
    gap: '12px',
    flex: '1 1 320px',
  },
  cardRightHeader: {
    display: 'flex',
    flexFlow: 'row wrap',
    justifyContent: 'space-between',
    alignItems: 'center',
    gap: '14px',
    width: '100%',
  },
  cardRightSemester: {
    flexShrink: 0,
    fontFamily: 'var(--font-mono)',
    fontSize: '0.78rem',
    fontWeight: 500,
    color: 'var(--green)',
    backgroundColor: 'rgba(46, 230, 166, 0.07)',
    borderWidth: '1px',
    borderStyle: 'solid',
    borderColor: 'rgba(46, 230, 166, 0.3)',
    borderRadius: '999px',
    padding: '6px 14px',
  },
  cardRightSummary: {
    margin: 0,
    maxWidth: '64ch',
    color: 'var(--text-dim)',
  },
  sidebarCard: {
    maxWidth: 'none',
    margin: 0,
    padding: 0,
    backgroundColor: 'transparent',
    borderWidth: 0,
    borderStyle: 'none',
    borderColor: 'transparent',
    borderRadius: 0,
    boxShadow: 'none',
    flexDirection: 'column',
    alignItems: 'flex-start',
    gap: '16px',
  },
  sidebarRight: {
    flex: '1 1 auto',
    minWidth: 0,
  },
  sidebarPicture: {
    width: '80px',
    height: '80px',
    borderRadius: '14px',
  },
});

export const card = stylex.props(styles.card);
export const cardCompact = stylex.props(styles.card, styles.sidebarCard);
export const cardLeft = stylex.props(styles.cardLeft);
export const cardLeftPicture = stylex.props(styles.cardLeftPicture);
export const cardLeftPictureCompact = stylex.props(styles.cardLeftPicture, styles.sidebarPicture);
export const cardRight = stylex.props(styles.cardRight);
export const cardRightCompact = stylex.props(styles.cardRight, styles.sidebarRight);
export const cardRightHeader = stylex.props(styles.cardRightHeader);
export const cardRightSemester = stylex.props(styles.cardRightSemester);
export const cardRightSummary = stylex.props(styles.cardRightSummary);
