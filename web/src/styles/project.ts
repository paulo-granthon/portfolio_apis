import * as stylex from '@stylexjs/stylex';
import * as base from './base';
export { base };

const styles = stylex.create({
  project: {
    position: 'relative',
    width: '100%',
    boxSizing: 'border-box',
    marginTop: 0,
    marginBottom: '32px',
    marginLeft: 0,
    marginRight: 0,
    backgroundColor: 'var(--bg-elev)',
    borderWidth: '1px',
    borderStyle: 'solid',
    borderColor: 'var(--border)',
    borderRadius: 'var(--radius)',
    // clip preserves rounded corners without blocking position:sticky in children
    overflow: 'clip',
    transition: 'border-color 0.22s ease, box-shadow 0.22s ease',
    ':hover': {
      borderColor: 'var(--pink)',
      boxShadow: '0 30px 70px -36px var(--pink-glow)',
    },
  },

  // Sticky 2-col header: info (left) + banner image (right)
  projectSticky: {
    position: 'sticky',
    top: 0,
    zIndex: 2,
    display: 'flex',
    flexFlow: 'row nowrap',
    minHeight: '240px',
    backgroundColor: 'var(--bg-elev)',
    borderBottomWidth: '1px',
    borderBottomStyle: 'solid',
    borderBottomColor: 'var(--border)',
  },
  projectStickyInfo: {
    flex: '1 1 0',
    minWidth: 0,
    padding: '22px 26px',
    display: 'flex',
    flexDirection: 'column',
    gap: '12px',
    position: 'relative',
    overflow: 'hidden',
  },
  projectStickyBanner: {
    flexShrink: 0,
    width: '42%',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    overflow: 'hidden',
    borderLeftWidth: '1px',
    borderLeftStyle: 'solid',
    borderLeftColor: 'var(--border)',
    backgroundColor: 'rgba(0, 0, 0, 0.35)',
  },
  bannerImg: {
    height: '100%',
    width: 'auto',
    display: 'block',
    flexShrink: 0,
  },

  // Scrolls below the sticky header
  projectScrollContent: {
    padding: '22px 26px 28px',
    display: 'flex',
    flexDirection: 'column',
    gap: '16px',
  },

  semesterGhost: {
    position: 'absolute',
    top: '-10px',
    right: '14px',
    fontFamily: 'var(--font-display)',
    fontWeight: 800,
    fontSize: '6.5rem',
    lineHeight: 1,
    color: 'var(--pink)',
    opacity: 0.07,
    pointerEvents: 'none',
    userSelect: 'none',
  },
  projectHeader: {
    position: 'relative',
    width: '100%',
    display: 'flex',
    flexFlow: 'row wrap',
    alignItems: 'center',
    justifyContent: 'space-between',
    gap: '10px',
  },
  projectHeaderTitle: {
    margin: 0,
    fontSize: '1.5rem',
  },
  projectHeaderExtra: {
    display: 'flex',
    flexFlow: 'row wrap',
    gap: '6px',
    justifyContent: 'flex-end',
  },
  projectHeaderExtraItem: {
    margin: 0,
    fontFamily: 'var(--font-mono)',
    fontSize: '0.72rem',
    fontWeight: 500,
    color: 'var(--pink-soft)',
    backgroundColor: 'rgba(255, 61, 129, 0.08)',
    borderWidth: '1px',
    borderStyle: 'solid',
    borderColor: 'rgba(255, 61, 129, 0.28)',
    borderRadius: '999px',
    padding: '4px 10px',
  },
  projectSubHeader: {
    display: 'flex',
    flexFlow: 'row wrap',
    alignItems: 'baseline',
    justifyContent: 'space-between',
    gap: '8px',
    margin: 0,
    padding: 0,
  },
  projectSummary: {
    margin: 0,
    color: 'var(--text)',
    fontWeight: 500,
    fontSize: '0.9rem',
  },
  projectRepo: {
    fontFamily: 'var(--font-mono)',
    fontSize: '0.78rem',
  },
  projectDescription: {
    margin: 0,
    color: 'var(--text-dim)',
  },
  participation: {
    marginTop: '2px',
    padding: '14px 18px',
    backgroundColor: 'rgba(255, 61, 129, 0.05)',
    borderLeftWidth: '3px',
    borderLeftStyle: 'solid',
    borderLeftColor: 'var(--pink)',
    borderRadius: '0 8px 8px 0',
  },
  participationTitle: {
    margin: '0 0 6px',
    fontFamily: 'var(--font-mono)',
    fontSize: '0.70rem',
    fontWeight: 600,
    letterSpacing: '0.14em',
    textTransform: 'uppercase',
    color: 'var(--pink)',
  },
  participationText: {
    margin: 0,
    color: 'var(--text)',
    fontSize: '0.88rem',
  },
});

export const project = stylex.props(styles.project);
export const projectSticky = stylex.props(styles.projectSticky);
export const projectStickyInfo = stylex.props(styles.projectStickyInfo);
export const projectStickyBanner = stylex.props(styles.projectStickyBanner);
export const bannerImg = stylex.props(styles.bannerImg);
export const projectScrollContent = stylex.props(styles.projectScrollContent);
export const semesterGhost = stylex.props(styles.semesterGhost);
export const projectHeader = stylex.props(styles.projectHeader);
export const projectHeaderTitle = stylex.props(styles.projectHeaderTitle);
export const projectHeaderExtra = stylex.props(styles.projectHeaderExtra);
export const projectHeaderExtraItem = stylex.props(styles.projectHeaderExtraItem);
export const projectSubHeader = stylex.props(styles.projectSubHeader);
export const projectSummary = stylex.props(styles.projectSummary);
export const projectRepo = stylex.props(styles.projectRepo);
export const projectDescription = stylex.props(styles.projectDescription);
export const participation = stylex.props(styles.participation);
export const participationTitle = stylex.props(styles.participationTitle);
export const participationText = stylex.props(styles.participationText);
