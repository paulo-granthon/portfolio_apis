import * as stylex from '@stylexjs/stylex';
import * as base from './base';
export { base };

const styles = stylex.create({
  layout: {
    display: 'flex',
    flexFlow: 'row nowrap',
    height: '100vh',
    overflow: 'hidden',
    position: 'relative',
  },
  sidebar: {
    flexShrink: 0,
    width: '320px',
    height: '100vh',
    overflowY: 'auto',
    overflowX: 'hidden',
    boxSizing: 'border-box',
    padding: '40px 20px 40px 24px',
    display: 'flex',
    flexDirection: 'column',
    gap: '28px',
    borderRightWidth: '1px',
    borderRightStyle: 'solid',
    borderRightColor: 'var(--border)',
    backdropFilter: 'blur(12px)',
  },
  scrollPane: {
    flex: '1 1 0',
    height: '100vh',
    overflowY: 'auto',
    overflowX: 'hidden',
    position: 'relative',
    paddingTop: '24px',
    boxSizing: 'border-box',
  },
  portfolio: {
    width: '100%',
    padding: 0,
  },
  kicker: {
    fontFamily: 'var(--font-mono)',
    fontSize: '0.8rem',
    fontWeight: 500,
    letterSpacing: '0.16em',
    textTransform: 'uppercase',
    color: 'var(--pink)',
    margin: '0 0 18px',
  },
  headerRow: {
    display: 'flex',
    flexFlow: 'row wrap',
    alignItems: 'flex-end',
    justifyContent: 'space-between',
    gap: '24px',
  },
  headerTitle: {
    margin: 0,
    fontSize: 'clamp(2.6rem, 7vw, 5rem)',
    backgroundImage: 'linear-gradient(102deg, var(--pink), var(--green))',
    backgroundClip: 'text',
    WebkitBackgroundClip: 'text',
    color: 'transparent',
    WebkitTextFillColor: 'transparent',
  },
  downloadButton: {
    fontFamily: 'var(--font-mono)',
    fontSize: '0.9rem',
    fontWeight: 500,
    color: 'var(--green)',
    backgroundColor: 'rgba(46, 230, 166, 0.06)',
    borderWidth: '1px',
    borderStyle: 'solid',
    borderColor: 'rgba(46, 230, 166, 0.4)',
    borderRadius: '11px',
    padding: '13px 22px',
    display: 'inline-flex',
    alignItems: 'center',
    gap: '9px',
    cursor: 'pointer',
    whiteSpace: 'nowrap',
    transition: 'all 0.18s ease',
    ':hover': {
      backgroundColor: 'rgba(46, 230, 166, 0.13)',
      borderColor: 'var(--green)',
      boxShadow: '0 0 28px var(--green-glow)',
      transform: 'translateY(-2px)',
      color: 'var(--green-soft)',
    },
  },
  projects: {
    display: 'flex',
    flexFlow: 'column',
    width: '100%',
    padding: '0 0 96px',
  },
  projectsHeader: {
    fontFamily: 'var(--font-mono)',
    fontSize: '0.8rem',
    fontWeight: 500,
    letterSpacing: '0.16em',
    textTransform: 'uppercase',
    color: 'var(--text-faint)',
    margin: '8px 0 22px',
    paddingBottom: '14px',
    borderBottomWidth: '1px',
    borderBottomStyle: 'solid',
    borderBottomColor: 'var(--border)',
  },
});

export const layout = stylex.props(styles.layout);
export const sidebar = stylex.props(styles.sidebar);
export const scrollPane = stylex.props(styles.scrollPane);
export const portfolio = stylex.props(styles.portfolio);
export const kicker = stylex.props(styles.kicker);
export const headerRow = stylex.props(styles.headerRow);
export const headerTitle = stylex.props(styles.headerTitle);
export const downloadButton = stylex.props(styles.downloadButton);
export const projects = stylex.props(styles.projects);
export const projectsHeader = stylex.props(styles.projectsHeader);
