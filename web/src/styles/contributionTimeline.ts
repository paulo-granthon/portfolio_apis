import * as stylex from '@stylexjs/stylex';

const s = stylex.create({
  root: {
    marginTop: '22px',
  },
  label: {
    fontFamily: 'var(--font-mono)',
    fontSize: '0.72rem',
    fontWeight: 600,
    letterSpacing: '0.14em',
    textTransform: 'uppercase',
    color: 'var(--text-faint)',
    margin: '0 0 8px',
  },
  chart: {
    display: 'flex',
    flexFlow: 'row nowrap',
    alignItems: 'flex-end',
    gap: '2px',
    height: '40px',
    overflow: 'hidden',
  },
  bar: {
    flex: '1 1 0',
    minWidth: '3px',
    maxWidth: '10px',
    backgroundColor: 'var(--green)',
    borderRadius: '2px 2px 0 0',
    opacity: 0.55,
    transition: 'opacity 0.15s',
    ':hover': {
      opacity: 1,
    },
  },
  empty: {
    color: 'var(--text-faint)',
    fontFamily: 'var(--font-mono)',
    fontSize: '0.72rem',
    margin: 0,
  },
});

export const root = stylex.props(s.root);
export const label = stylex.props(s.label);
export const chart = stylex.props(s.chart);
export const bar = stylex.props(s.bar);
export const empty = stylex.props(s.empty);
