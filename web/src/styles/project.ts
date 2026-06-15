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
  bannerSticky: {
    position: 'sticky',
    top: 0,
    zIndex: 2,
    height: '240px',
    overflow: 'hidden',
    backgroundColor: 'rgba(0, 0, 0, 0.35)',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    borderBottomWidth: '1px',
    borderBottomStyle: 'solid',
    borderBottomColor: 'var(--border)',
  },
  // Every banner gets the same vertical space. Wide banners (most of them) use
  // `cover` so they fill the section edge to edge; a banner with an unusual
  // aspect ratio (api3 is square) falls back to `contain` so it isn't cropped,
  // shown centered with side breathing room. The choice is made from the
  // image's own aspect ratio once it loads (see Project).
  projectImageContain: {
    maxWidth: '100%',
    maxHeight: '100%',
    width: 'auto',
    height: 'auto',
    objectFit: 'contain',
    display: 'block',
  },
  projectImageCover: {
    width: '100%',
    height: '100%',
    objectFit: 'cover',
    display: 'block',
  },
  projectBody: {
    position: 'relative',
    padding: '26px 30px 30px',
  },
  semesterGhost: {
    position: 'absolute',
    top: '-14px',
    right: '18px',
    fontFamily: 'var(--font-display)',
    fontWeight: 800,
    fontSize: '7rem',
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
    gap: '14px',
  },
  projectHeaderTitle: {
    margin: 0,
    fontSize: '1.7rem',
  },
  projectHeaderExtra: {
    display: 'flex',
    flexFlow: 'row wrap',
    gap: '8px',
    justifyContent: 'flex-end',
  },
  projectHeaderExtraItem: {
    margin: 0,
    fontFamily: 'var(--font-mono)',
    fontSize: '0.74rem',
    fontWeight: 500,
    color: 'var(--pink-soft)',
    backgroundColor: 'rgba(255, 61, 129, 0.08)',
    borderWidth: '1px',
    borderStyle: 'solid',
    borderColor: 'rgba(255, 61, 129, 0.28)',
    borderRadius: '999px',
    padding: '5px 12px',
  },
  projectSubHeader: {
    margin: '14px 0 0',
    padding: 0,
    display: 'flex',
    flexFlow: 'row wrap',
    alignItems: 'baseline',
    justifyContent: 'space-between',
    gap: '10px',
  },
  projectSummary: {
    margin: 0,
    color: 'var(--text)',
    fontWeight: 500,
  },
  projectRepo: {
    fontFamily: 'var(--font-mono)',
    fontSize: '0.82rem',
  },
  projectDescription: {
    margin: '18px 0 0',
    color: 'var(--text-dim)',
  },
  participation: {
    margin: '22px 0 4px',
    padding: '18px 20px',
    backgroundColor: 'rgba(255, 61, 129, 0.05)',
    borderLeftWidth: '3px',
    borderLeftStyle: 'solid',
    borderLeftColor: 'var(--pink)',
    borderRadius: '0 10px 10px 0',
  },
  participationTitle: {
    margin: '0 0 8px',
    fontFamily: 'var(--font-mono)',
    fontSize: '0.74rem',
    fontWeight: 600,
    letterSpacing: '0.14em',
    textTransform: 'uppercase',
    color: 'var(--pink)',
  },
  participationText: {
    margin: 0,
    color: 'var(--text)',
  },
});

export const project = stylex.props(styles.project);
export const bannerSticky = stylex.props(styles.bannerSticky);
export const projectImageContain = stylex.props(styles.projectImageContain);
export const projectImageCover = stylex.props(styles.projectImageCover);
export const projectBody = stylex.props(styles.projectBody);
export const semesterGhost = stylex.props(styles.semesterGhost);
export const projectHeader = stylex.props(styles.projectHeader);
export const projectHeaderTitle = stylex.props(styles.projectHeaderTitle);
export const projectHeaderExtra = stylex.props(styles.projectHeaderExtra);
export const projectHeaderExtraItem = stylex.props(
  styles.projectHeaderExtraItem,
);
export const projectSubHeader = stylex.props(styles.projectSubHeader);
export const projectSummary = stylex.props(styles.projectSummary);
export const projectRepo = stylex.props(styles.projectRepo);
export const projectDescription = stylex.props(styles.projectDescription);
export const participation = stylex.props(styles.participation);
export const participationTitle = stylex.props(styles.participationTitle);
export const participationText = stylex.props(styles.participationText);
