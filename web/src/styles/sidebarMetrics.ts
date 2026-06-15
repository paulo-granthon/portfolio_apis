import * as stylex from '@stylexjs/stylex';

const s = stylex.create({
  root: {
    display: 'flex',
    flexDirection: 'column',
    gap: '14px',
    padding: '20px',
    position: 'relative',
    backgroundColor: 'var(--bg-elev)',
    borderWidth: '1px',
    borderStyle: 'solid',
    borderColor: 'var(--border)',
    borderRadius: 'var(--radius)',
    overflow: 'hidden',
    transition: 'border-color 0.3s ease',
  },
  ghost: {
    position: 'absolute',
    top: '-12px',
    right: '12px',
    fontFamily: 'var(--font-display)',
    fontWeight: 800,
    fontSize: '6rem',
    lineHeight: 1,
    color: 'var(--pink)',
    opacity: 0.07,
    pointerEvents: 'none',
    userSelect: 'none',
  },
  name: {
    margin: 0,
    fontSize: '1.1rem',
    fontWeight: 700,
    color: 'var(--text)',
    lineHeight: 1.3,
    position: 'relative',
  },
  stats: {
    display: 'flex',
    flexDirection: 'column',
    gap: '8px',
  },
  statRow: {
    display: 'flex',
    alignItems: 'baseline',
    gap: '8px',
  },
  statNum: {
    fontFamily: 'var(--font-display)',
    fontWeight: 800,
    fontSize: '1.6rem',
    color: 'var(--green)',
    lineHeight: 1,
  },
  statLabel: {
    fontFamily: 'var(--font-mono)',
    fontSize: '0.72rem',
    fontWeight: 500,
    color: 'var(--text-faint)',
    letterSpacing: '0.08em',
  },
  divider: {
    height: '1px',
    backgroundColor: 'var(--border)',
    border: 'none',
    margin: '2px 0',
  },
  placeholder: {
    fontFamily: 'var(--font-mono)',
    fontSize: '0.76rem',
    color: 'var(--text-faint)',
    margin: 0,
    opacity: 0.6,
  },
});

export const root = stylex.props(s.root);
export const ghost = stylex.props(s.ghost);
export const name = stylex.props(s.name);
export const stats = stylex.props(s.stats);
export const statRow = stylex.props(s.statRow);
export const statNum = stylex.props(s.statNum);
export const statLabel = stylex.props(s.statLabel);
export const divider = stylex.props(s.divider);
export const placeholder = stylex.props(s.placeholder);
