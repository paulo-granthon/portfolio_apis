import * as stylex from '@stylexjs/stylex';
import * as base from './base';
export { base };

const styles = stylex.create({
  contributions: {
    display: 'flex',
    flexDirection: 'column',
    marginTop: '24px',
    paddingTop: '20px',
    borderTopWidth: '1px',
    borderTopStyle: 'solid',
    borderTopColor: 'var(--border)',
  },
  contributionsTitle: {
    margin: '0 0 14px',
    fontFamily: 'var(--font-mono)',
    fontSize: '0.78rem',
    fontWeight: 600,
    letterSpacing: '0.14em',
    textTransform: 'uppercase',
    color: 'var(--green)',
  },
  contributionList: {
    display: 'flex',
    flexDirection: 'column',
    gap: '12px',
  },
  contribution: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'stretch',
    gap: '6px',
    padding: '16px 18px',
    backgroundColor: 'var(--bg-elev-2)',
    borderWidth: '1px',
    borderStyle: 'solid',
    borderColor: 'var(--border)',
    borderRadius: '11px',
    transition: 'border-color 0.18s ease, box-shadow 0.18s ease',
    ':hover': {
      borderColor: 'rgba(46, 230, 166, 0.5)',
      boxShadow: '0 0 0 1px rgba(46, 230, 166, 0.15), 0 18px 40px -28px var(--green-glow)',
    },
  },
  contributionHeader: {
    width: '100%',
    display: 'flex',
    flexDirection: 'column',
    gap: '8px',
  },
  title: {
    margin: 0,
    fontSize: '1.02rem',
    color: 'var(--text)',
  },
  skills: {
    display: 'flex',
    flexFlow: 'row wrap',
    margin: '2px 0 0',
    gap: '7px',
  },
  // skill tags: green pills, visually distinct from the magenta inline code
  skill: {
    height: 'fit-content',
    fontSize: '0.72rem',
    letterSpacing: '0.02em',
    color: 'var(--green-soft)',
    backgroundColor: 'rgba(46, 230, 166, 0.07)',
    borderColor: 'rgba(46, 230, 166, 0.28)',
    borderRadius: '999px',
    padding: '0.22em 0.7em',
    whiteSpace: 'nowrap',
    transition: 'border-color 0.18s ease, box-shadow 0.18s ease, transform 0.18s ease',
    ':hover': {
      borderColor: 'var(--green)',
      boxShadow: '0 0 16px var(--green-glow)',
      transform: 'translateY(-1px)',
    },
  },
  content: {
    margin: '4px 0 0',
    color: 'var(--text-dim)',
    fontSize: '0.95rem',
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
