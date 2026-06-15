# Two-Section Sticky Layout + Contribution Timeline Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Transform the single-column portfolio page into a two-column layout with a sticky "me" sidebar on the left and a scrollable project list on the right, where each project's banner header sticks to the top of the right pane until the next project displaces it; then add per-project GitHub commit-history graphs.

**Architecture:** The right column becomes the sole scroll container (`overflow-y: auto; height: 100vh`); the left sidebar is `position: sticky; top: 0; height: 100vh`. Each project `<article>` uses `overflow: clip` (not `hidden`) so its inner banner div can be `position: sticky; top: 0` relative to the right scroll pane. `useScrollScenes` accepts an optional `HTMLElement | null` container so it can listen to that element's scroll event instead of `window`.

**Tech Stack:** React 18, Vite, StyleX (`@stylexjs/stylex`), TypeScript, NX monorepo (`yarn nx typecheck web`, `yarn nx test web`)

---

## File Map

| Action | File | Responsibility |
|--------|------|---------------|
| Modify | `web/src/theme/useScrollTheme.ts` | Accept `container?: HTMLElement \| null`; swap scroll listener target |
| Modify | `web/src/pages/portfolio.tsx` | Two-column layout wrapper; pass `containerRef` to hook |
| Modify | `web/src/styles/portfolio.ts` | Add `layout`, `sidebar`, `scrollPane` styles |
| Modify | `web/src/components/project.tsx` | Promote banner to sticky header, fix article `overflow` |
| Modify | `web/src/styles/project.ts` | Add `bannerSticky`; change `overflow` on article wrapper |
| Modify | `web/src/components/projectList.tsx` | Remove outer horizontal padding (now owned by scroll pane) |
| Modify | `web/src/styles/user.ts` | Sidebar-friendly card: column direction, remove max-width centering |
| Modify | `web/src/components/user.tsx` | Accept optional `compact` prop to pick sidebar vs full-width style |
| Modify | `web/src/services/github.ts` | Add `getProjectCommits(repoUrl, author)` |
| Create | `web/src/components/contributionTimeline.tsx` | Commit-count bar chart per project |
| Create | `web/src/styles/contributionTimeline.ts` | Timeline chart styles |
| Modify | `web/src/components/project.tsx` | Render `<ContributionTimeline>` below participation block |
| Modify | `web/src/schemas/portfolio.ts` | Pass `githubUsername` through to `Project` (or use context) |

---

## Phase 1 — Sticky Layout

### Task 1: Teach `useScrollScenes` about a scroll container

**Files:**
- Modify: `web/src/theme/useScrollTheme.ts`
- Test: `web/src/theme/useScrollTheme.test.ts` (if it exists; skip if not)

The hook currently calls `window.addEventListener('scroll', ...)`. After this task it accepts an optional `container` and listens on it instead.

- [ ] **Step 1.1: Change the hook signature**

In `web/src/theme/useScrollTheme.ts`, update `useScrollScenes`:

```ts
export function useScrollScenes(
  dep: unknown,
  container?: HTMLElement | null,
): SceneState {
```

The rest of the function body is unchanged at first — we'll wire the listener next.

- [ ] **Step 1.2: Swap the scroll / resize listeners**

Replace the two `window.addEventListener` calls and their cleanup:

```ts
// before
window.addEventListener('scroll', onScroll, { passive: true });
window.addEventListener('resize', onScroll);
return () => {
  window.removeEventListener('scroll', onScroll);
  window.removeEventListener('resize', onScroll);
  if (raf) cancelAnimationFrame(raf);
};
```

with:

```ts
const target: EventTarget = container ?? window;
target.addEventListener('scroll', onScroll, { passive: true });
window.addEventListener('resize', onScroll); // resize always on window
return () => {
  target.removeEventListener('scroll', onScroll);
  window.removeEventListener('resize', onScroll);
  if (raf) cancelAnimationFrame(raf);
};
```

Resize stays on `window` because the container's size changes come from the window resizing.

- [ ] **Step 1.3: Re-run when the container changes**

The effect must re-subscribe if the container ref flips from `null` to the element. Add `container` to the dependency array:

```ts
}, [dep, container]);
```

- [ ] **Step 1.4: Typecheck**

```bash
yarn nx typecheck web
```

Expected: no errors.

- [ ] **Step 1.5: Commit**

```bash
git add web/src/theme/useScrollTheme.ts
git commit -m "feat(web): accept container element in useScrollScenes scroll listener

Switch from window scroll to an optional container element.
The effect re-subscribes when the container ref resolves (null -> element).
getBoundingClientRect() is always viewport-relative so no geometry change is needed."
```

---

### Task 2: Two-column layout in `portfolio.tsx` and `portfolio.ts`

**Files:**
- Modify: `web/src/pages/portfolio.tsx`
- Modify: `web/src/styles/portfolio.ts`

The page becomes `display: flex; height: 100vh; overflow: hidden`. The left column is a fixed-width sticky sidebar; the right column is the scroll pane.

- [ ] **Step 2.1: Add layout styles to `portfolio.ts`**

Add to the `stylex.create({...})` object (keep all existing rules):

```ts
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
  },
```

And add exports at the bottom:

```ts
export const layout = stylex.props(styles.layout);
export const sidebar = stylex.props(styles.sidebar);
export const scrollPane = stylex.props(styles.scrollPane);
```

- [ ] **Step 2.2: Rewrite the `portfolio.tsx` render**

The current render returns a fragment with `<BackgroundFx>` and a `<div>`. Replace with:

```tsx
import { useRef } from 'react';
// ... existing imports unchanged

export default function Portfolio() {
  const [portfolio, setPortfolio] = useState<PortfolioSchema | undefined>();
  const params = useParams<{ userId: string }>();
  const scrollRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const id = params.userId ? parseInt(params.userId, 10) : undefined;
    if (!id) return;
    getPortfolio(id).then(portfolio => setPortfolio(portfolio));
  }, [params.userId]);

  const { active, current, next, pct } = useScrollScenes(
    portfolio?.projects.length,
    scrollRef.current,
  );
  useEffect(() => applyPalette(active), [active]);

  if (!portfolio) {
    return <p>Usuário não encontrado</p>;
  }

  return (
    <>
      <BackgroundFx current={current} next={next} pct={pct} />
      <div {...styles.layout}>
        <aside {...styles.sidebar}>
          <header {...styles.portfolio}>
            <p {...styles.kicker}>// portfólio · banco de dados · fatec-sjc</p>
            <div {...styles.headerRow}>
              <h1 {...styles.headerTitle}>Portfólio</h1>
              <a
                {...styles.downloadButton}
                href={portfolioMarkdownUrl(portfolio.user.id)}
                download
              >
                › baixar.md
              </a>
            </div>
          </header>
          <User user={portfolio.user} compact />
        </aside>
        <main {...styles.scrollPane} ref={scrollRef}>
          <ProjectList projects={portfolio.projects} />
        </main>
      </div>
    </>
  );
}
```

Note: `scrollRef.current` is `null` on the first render. The effect in `useScrollScenes` re-runs when `scrollRef.current` changes (because it is passed as `container` and in the dep array) — but only if `scrollRef.current` itself appears as a dep. The cleanest approach to trigger re-subscription is to use a state-based ref callback instead of `useRef`. Replace `const scrollRef = useRef<HTMLDivElement>(null)` with:

```tsx
const [scrollEl, setScrollEl] = useState<HTMLDivElement | null>(null);
```

And on the `<main>`:

```tsx
<main {...styles.scrollPane} ref={setScrollEl}>
```

Then:

```tsx
const { active, current, next, pct } = useScrollScenes(
  portfolio?.projects.length,
  scrollEl,
);
```

This way, when the `<main>` mounts, `setScrollEl` fires, `scrollEl` becomes the element, and `useScrollScenes` re-runs its effect with the real container.

- [ ] **Step 2.3: Adjust header styles in `portfolio.ts`**

The header is now inside the narrow sidebar (320px). Remove `maxWidth` and `margin: 0 auto` from the `portfolio` rule, and reduce padding:

```ts
  portfolio: {
    width: '100%',
    padding: 0,
  },
```

The sidebar itself provides the padding.

- [ ] **Step 2.4: Typecheck**

```bash
yarn nx typecheck web
```

Expected: no errors (the `compact` prop on `User` will error until Task 3).

- [ ] **Step 2.5: Commit** (after Task 3 is done and typechecks pass)

---

### Task 3: Sticky banner header + `User` sidebar variant

**Files:**
- Modify: `web/src/components/user.tsx`
- Modify: `web/src/styles/user.ts`
- Modify: `web/src/components/project.tsx`
- Modify: `web/src/styles/project.ts`
- Modify: `web/src/components/projectList.tsx`
- Modify: `web/src/styles/portfolio.ts` (projects / projectsHeader tweaks)

#### 3a: User — compact sidebar variant

- [ ] **Step 3a.1: Add `sidebarCard` style to `user.ts`**

Add to the `stylex.create({...})`:

```ts
  sidebarCard: {
    maxWidth: 'none',
    margin: 0,
    padding: '0',
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
  sidebarCardRight: {
    flex: '1 1 auto',
    minWidth: 0,
  },
  sidebarName: {
    margin: '0 0 6px',
    fontSize: '1.1rem',
    fontWeight: 700,
  },
```

Export:

```ts
export const sidebarCard = stylex.props(styles.sidebarCard);
export const sidebarCardRight = stylex.props(styles.sidebarCardRight);
export const sidebarName = stylex.props(styles.sidebarName);
```

- [ ] **Step 3a.2: Update `user.tsx` to accept `compact` prop**

```tsx
interface UserProps {
  user: UserSchema;
  compact?: boolean;
}

export default function User({ user, compact }: UserProps) {
  // ... unchanged up to return

  return (
    <div {...userStyles.card} {...(compact ? userStyles.sidebarCard : {})}>
      <div {...userStyles.cardLeft}>
        <img {...userStyles.cardLeftPicture} alt={user.name} src={avatarUrl} />
      </div>
      <div {...userStyles.cardRight} {...(compact ? userStyles.sidebarCardRight : {})}>
        <div {...userStyles.cardRightHeader}>
          <h2 {...(compact ? userStyles.sidebarName : {})}>{user.name}</h2>
          <span {...userStyles.cardRightSemester}>
            semestre atual · {userCurrentSemester}
          </span>
        </div>
        <p {...userStyles.cardRightSummary}>{user.summary}</p>
      </div>
    </div>
  );
}
```

The StyleX spread pattern: `{...stylex.props(styles.card, compact && styles.sidebarCard)}` is idiomatic StyleX and avoids double-spreading. Rewrite user.tsx to use merged props:

```tsx
export default function User({ user, compact }: UserProps) {
  const userCurrentSemester = user.semesterMatriculed.currentSemester().toString();
  const avatarUrl = user.githubUsername
    ? `https://github.com/${user.githubUsername}.png?size=240`
    : undefined;

  return (
    <div {...stylex.props(s.card, compact && s.sidebarCard)}>
      <div {...stylex.props(s.cardLeft)}>
        <img {...stylex.props(s.cardLeftPicture)} alt={user.name} src={avatarUrl} />
      </div>
      <div {...stylex.props(s.cardRight, compact && s.sidebarCardRight)}>
        <div {...stylex.props(s.cardRightHeader)}>
          <h2 {...stylex.props(compact && s.sidebarName)}>{user.name}</h2>
          <span {...stylex.props(s.cardRightSemester)}>
            semestre atual · {userCurrentSemester}
          </span>
        </div>
        <p {...stylex.props(s.cardRightSummary)}>{user.summary}</p>
      </div>
    </div>
  );
}
```

Where `s` = the renamed private `styles` constant (rename at top of file: `const s = stylex.create({...})`), and add `import * as stylex from '@stylexjs/stylex'` if not already present. Remove the named exports at the bottom of `user.ts` — the component now accesses styles directly. Instead the `.ts` file exports only the raw `stylex.create` object if needed by tests; or simply make the styles file internal-only.

Actually, StyleX convention in this codebase is: style file exports `stylex.props(...)` objects which are spread onto elements. Merging two such objects requires using `stylex.props(a, b)` not `{...a, ...b}` (StyleX deduplicates atomic class names). To use `stylex.props(s.card, compact && s.sidebarCard)`, the component needs to import both the styles object AND `stylex`. 

Simpler approach consistent with existing patterns: keep the exported props pattern, but have two variants exported from `user.ts`:

```ts
// user.ts: export both
export const card = stylex.props(styles.card);
export const cardCompact = stylex.props(styles.card, styles.sidebarCard);
```

Then in `user.tsx`:
```tsx
<div {...(compact ? userStyles.cardCompact : userStyles.card)}>
```

This is cleaner. Use this approach.

- [ ] **Step 3a.3: Full revised `user.ts` additions**

Append to `stylex.create`:

```ts
  sidebarCard: {
    maxWidth: 'none',
    margin: 0,
    padding: '0',
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
  sidebarH2: {
    margin: '0 0 4px',
    fontSize: '1.05rem',
  },
```

Add exports:

```ts
export const cardCompact = stylex.props(styles.card, styles.sidebarCard);
export const cardRightCompact = stylex.props(styles.cardRight, styles.sidebarRight);
export const cardRightHeaderCompact = stylex.props(styles.cardRightHeader, styles.sidebarH2);
```

Then `user.tsx`:

```tsx
export default function User({ user, compact }: UserProps) {
  const userCurrentSemester = user.semesterMatriculed.currentSemester().toString();
  const avatarUrl = user.githubUsername
    ? `https://github.com/${user.githubUsername}.png?size=240`
    : undefined;

  return (
    <div {...(compact ? userStyles.cardCompact : userStyles.card)}>
      <div {...userStyles.cardLeft}>
        <img {...userStyles.cardLeftPicture} alt={user.name} src={avatarUrl} />
      </div>
      <div {...(compact ? userStyles.cardRightCompact : userStyles.cardRight)}>
        <div {...userStyles.cardRightHeader}>
          <h2>{user.name}</h2>
          <span {...userStyles.cardRightSemester}>
            semestre atual · {userCurrentSemester}
          </span>
        </div>
        <p {...userStyles.cardRightSummary}>{user.summary}</p>
      </div>
    </div>
  );
}
```

#### 3b: Project — sticky banner + `overflow: clip`

- [ ] **Step 3b.1: Add `bannerSticky` and update `project` style in `project.ts`**

Change `overflow: 'hidden'` to `overflow: 'clip'` in the `project` rule. Add new style:

```ts
  project: {
    // ... keep all existing rules, but change:
    overflow: 'clip',          // was 'hidden' — clip preserves sticky children
    // remove the :hover transform — the card itself no longer lifts; body does
  },
  bannerSticky: {
    position: 'sticky',
    top: 0,
    zIndex: 2,
    height: '240px',
    overflow: 'hidden',
    borderTopLeftRadius: 'var(--radius)',
    borderTopRightRadius: 'var(--radius)',
    borderBottomWidth: '1px',
    borderBottomStyle: 'solid',
    borderBottomColor: 'var(--border)',
    backgroundColor: 'rgba(0,0,0,0.35)',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
  },
```

Export:

```ts
export const bannerSticky = stylex.props(styles.bannerSticky);
```

Note: remove `projectImageContainer` export (replaced by `bannerSticky`), or keep it if used elsewhere.

- [ ] **Step 3b.2: Update `project.tsx` to use sticky banner**

```tsx
export default function Project({ project }: ProjectProps) {
  const [cover, setCover] = useState(true);

  return (
    <article {...styles.project} data-project={project.name}>
      {project.image && (
        <div {...styles.bannerSticky}>
          <img
            {...(cover ? styles.projectImageCover : styles.projectImageContain)}
            src={project.image}
            alt={`${project.name} banner`}
            onLoad={e => {
              const img = e.currentTarget;
              setCover(img.naturalWidth / img.naturalHeight >= 1.9);
            }}
          />
        </div>
      )}
      <div {...styles.projectBody}>
        {/* unchanged body */}
      </div>
    </article>
  );
}
```

The banner now has `position: sticky; top: 0` (via `bannerSticky`). The article has `overflow: clip`. The right scroll pane is the sticky ancestor.

- [ ] **Step 3b.3: Update `projectList.tsx` — remove extra padding**

The scroll pane already provides padding for the right column. Remove horizontal padding from the `projects` rule in `portfolio.ts` or zero it out in `projectList.tsx`:

In `portfolio.ts`, update the `projects` rule:

```ts
  projects: {
    display: 'flex',
    flexFlow: 'column',
    width: '100%',
    padding: '20px 0 96px',   // remove left/right padding — scroll pane owns that
  },
```

And `scrollPane` already has no padding, so add padding to `scrollPane` or to `projects`. Cleaner: add `padding: '0 24px'` to `scrollPane` so all content inside is inset consistently.

Update `scrollPane` in `portfolio.ts`:

```ts
  scrollPane: {
    flex: '1 1 0',
    height: '100vh',
    overflowY: 'auto',
    overflowX: 'hidden',
    position: 'relative',
    padding: '0 24px',
    boxSizing: 'border-box',
  },
```

**Important:** When the sticky banner is inside `<article>` inside `<ProjectList>` inside `<main class="scrollPane">` with `padding: 0 24px`, the banner will only be as wide as the padded content area (not full-bleed to the scroll pane edge). If you want banners to be full-bleed (edge-to-edge), remove the padding from `scrollPane` and instead apply it only to `projectBody` (not the banner). In `project.ts`:

```ts
  projectBody: {
    position: 'relative',
    padding: '26px 30px 30px',   // unchanged — body has its own padding
  },
```

And the `article` has no horizontal padding, `bannerSticky` fills 100% width of the article, which is 100% width of the scroll pane (no horizontal padding there). This gives full-bleed banners.

Remove `padding` from `scrollPane` and instead add `paddingTop: '24px'` only, plus `paddingBottom: '96px'` on the projects container. Final `scrollPane`:

```ts
  scrollPane: {
    flex: '1 1 0',
    height: '100vh',
    overflowY: 'auto',
    overflowX: 'hidden',
    position: 'relative',
    paddingTop: '24px',
    boxSizing: 'border-box',
  },
```

- [ ] **Step 3b.4: Typecheck and visual check**

```bash
yarn nx typecheck web
```

Open `http://localhost:3333` and confirm:
1. Left sidebar is fixed — User card and header visible while scrolling right column
2. Project banners stick to top of right pane as you scroll through each project
3. Next project's banner naturally replaces the previous one (no overlap artifact)
4. Background scene crossfades still trigger as before
5. Palette still morphs based on which project is in view

- [ ] **Step 3b.5: Commit**

```bash
git add \
  web/src/theme/useScrollTheme.ts \
  web/src/pages/portfolio.tsx \
  web/src/styles/portfolio.ts \
  web/src/components/user.tsx \
  web/src/styles/user.ts \
  web/src/components/project.tsx \
  web/src/styles/project.ts \
  web/src/components/projectList.tsx

git commit -m "feat(web): two-column sticky layout with per-project banner headers

Left column is a sticky sidebar holding the user card and page header.
Right column is the sole scroll container; useScrollScenes now listens to it.
Project banners use position:sticky + overflow:clip (not hidden) on the article,
so each banner stays pinned at the top of the right pane until the next one takes over."
```

---

## Phase 2 — Contribution Timeline

### Task 4: Fetch commit history from GitHub API

**Files:**
- Modify: `web/src/services/github.ts`

**Data shape:** `project.url` is a string like `github.com/API-6-Semestre/api-6`. We parse the owner and repo from it, then call the GitHub REST API for commits filtered by the portfolio user's github username.

- [ ] **Step 4.1: Add `CommitEntry` type and `getProjectCommits` to `github.ts`**

```ts
export interface CommitEntry {
  date: string; // ISO-8601, e.g. "2024-03-15T10:22:00Z"
  sha: string;
}

export async function getProjectCommits(
  repoUrl: string,
  author: string,
): Promise<CommitEntry[]> {
  // repoUrl shape: "github.com/owner/repo" — strip protocol prefix
  const path = repoUrl.replace(/^https?:\/\//, '').replace(/^github\.com\//, '');
  const [owner, repo] = path.split('/');
  if (!owner || !repo) return [];

  const githubToken = import.meta.env.VITE_GITHUB_TOKEN;
  const headers: HeadersInit = githubToken
    ? { Authorization: `Bearer ${githubToken}` }
    : {};

  let page = 1;
  const commits: CommitEntry[] = [];

  while (true) {
    const url =
      `${API_URL}/repos/${owner}/${repo}/commits` +
      `?author=${encodeURIComponent(author)}&per_page=100&page=${page}`;

    const res = await fetch(url, { headers });
    if (!res.ok) break;

    const data = await res.json();
    if (!Array.isArray(data) || data.length === 0) break;

    for (const c of data) {
      const date = c.commit?.author?.date ?? c.commit?.committer?.date;
      if (date) commits.push({ date, sha: c.sha });
    }

    if (data.length < 100) break;
    page++;
  }

  return commits;
}
```

- [ ] **Step 4.2: Typecheck**

```bash
yarn nx typecheck web
```

Expected: no errors.

- [ ] **Step 4.3: Commit**

```bash
git add web/src/services/github.ts
git commit -m "feat(web): add getProjectCommits to GitHub service

Fetches all commits by a given author for a project's GitHub repo.
Paginates through 100-per-page until exhausted. Falls back gracefully on non-ok responses."
```

---

### Task 5: `ContributionTimeline` component

**Files:**
- Create: `web/src/components/contributionTimeline.tsx`
- Create: `web/src/styles/contributionTimeline.ts`
- Modify: `web/src/components/project.tsx` (add `githubUsername` prop + render timeline)
- Modify: `web/src/components/projectList.tsx` (pass `githubUsername`)
- Modify: `web/src/pages/portfolio.tsx` (pass `portfolio.user.githubUsername`)

**Approach:** Fetch commits once per project on mount. Group by ISO week (year + week number). Render as a mini horizontal bar chart — each week is one bar column, height proportional to commit count in that week. Total width: a rolling 52-week window (1 year) or the project's actual lifetime, whichever is shorter.

- [ ] **Step 5.1: Create `contributionTimeline.ts` styles**

```ts
// web/src/styles/contributionTimeline.ts
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
```

- [ ] **Step 5.2: Create `contributionTimeline.tsx`**

```tsx
// web/src/components/contributionTimeline.tsx
import { useEffect, useState } from 'react';
import { getProjectCommits, CommitEntry } from '../services/github';
import * as s from '../styles/contributionTimeline';

interface Props {
  repoUrl: string;
  author: string;
}

function isoWeek(date: Date): string {
  // YYYY-Www (ISO 8601 week)
  const d = new Date(Date.UTC(date.getFullYear(), date.getMonth(), date.getDate()));
  const day = d.getUTCDay() || 7;
  d.setUTCDate(d.getUTCDate() + 4 - day);
  const year = d.getUTCFullYear();
  const week = Math.ceil(((+d - +new Date(Date.UTC(year, 0, 1))) / 86400000 + 1) / 7);
  return `${year}-W${String(week).padStart(2, '0')}`;
}

function bucketByWeek(commits: CommitEntry[]): { week: string; count: number }[] {
  const map = new Map<string, number>();
  for (const c of commits) {
    const w = isoWeek(new Date(c.date));
    map.set(w, (map.get(w) ?? 0) + 1);
  }
  if (map.size === 0) return [];

  const sorted = [...map.keys()].sort();
  const first = sorted[0];
  const last = sorted[sorted.length - 1];

  // fill in zero-count weeks between first and last
  const result: { week: string; count: number }[] = [];
  const cur = new Date(first + '-1'); // approximate, good enough for display
  const end = new Date(last + '-1');
  // iterate week by week
  let w = first;
  while (w <= last) {
    result.push({ week: w, count: map.get(w) ?? 0 });
    // advance by 7 days
    const d = new Date(cur);
    d.setUTCDate(d.getUTCDate() + 7);
    cur.setUTCDate(cur.getUTCDate() + 7);
    w = isoWeek(cur);
    if (w === result[result.length - 1].week) break; // safety
  }
  return result;
}

export function ContributionTimeline({ repoUrl, author }: Props) {
  const [buckets, setBuckets] = useState<{ week: string; count: number }[] | null>(null);

  useEffect(() => {
    let cancelled = false;
    getProjectCommits(repoUrl, author).then(commits => {
      if (!cancelled) setBuckets(bucketByWeek(commits));
    });
    return () => { cancelled = true; };
  }, [repoUrl, author]);

  if (buckets === null) return null; // loading — render nothing
  if (buckets.length === 0) return <p {...s.empty}>sem commits registrados</p>;

  const max = Math.max(...buckets.map(b => b.count), 1);

  return (
    <div {...s.root}>
      <p {...s.label}>contribuições por semana</p>
      <div {...s.chart}>
        {buckets.map(b => (
          <div
            key={b.week}
            {...s.bar}
            style={{ height: `${Math.max(4, (b.count / max) * 40)}px` }}
            title={`${b.week}: ${b.count} commit${b.count !== 1 ? 's' : ''}`}
          />
        ))}
      </div>
    </div>
  );
}
```

- [ ] **Step 5.3: Thread `githubUsername` through to `Project`**

In `projectList.tsx`:

```tsx
interface ProjectListProps {
  projects: PortfolioProjectSchema[];
  githubUsername?: string;
}

export default function ProjectList({ projects, githubUsername }: ProjectListProps) {
  return (
    <div {...styles.projects}>
      <h2 {...styles.projectsHeader}>
        projetos · {projects.length} semestres
      </h2>
      <div>
        {projects.map(project => (
          <Project key={project.name} project={project} githubUsername={githubUsername} />
        ))}
      </div>
    </div>
  );
}
```

In `project.tsx`, add to `ProjectProps`:

```ts
interface ProjectProps {
  project: PortfolioProjectSchema;
  githubUsername?: string;
}
```

And render the timeline after `<ContributionList>`:

```tsx
{githubUsername && project.url && (
  <ContributionTimeline
    repoUrl={project.url}
    author={githubUsername}
  />
)}
```

In `portfolio.tsx`, pass username to `ProjectList`:

```tsx
<ProjectList
  projects={portfolio.projects}
  githubUsername={portfolio.user.githubUsername ?? undefined}
/>
```

- [ ] **Step 5.4: Typecheck**

```bash
yarn nx typecheck web
```

Expected: no errors.

- [ ] **Step 5.5: Visual test**

Open `http://localhost:3333/portfolio/1` (or your actual route). Confirm:
1. Below each project's contribution list, a mini bar chart appears (or nothing if no commits)
2. Bars are proportional to commit count per week
3. Hover on a bar shows the week and count in a tooltip
4. No console errors about GitHub API

- [ ] **Step 5.6: Commit**

```bash
git add \
  web/src/components/contributionTimeline.tsx \
  web/src/styles/contributionTimeline.ts \
  web/src/components/project.tsx \
  web/src/components/projectList.tsx \
  web/src/pages/portfolio.tsx

git commit -m "feat(web): per-project GitHub commit timeline chart

Fetches commits per project via GitHub API filtered by the portfolio user's
github username. Groups by ISO week and renders as a proportional bar chart.
Renders nothing while loading; renders a fallback message if no commits found."
```

---

## Self-Review

### Spec coverage

| Requirement | Task |
|-------------|------|
| Sticky "me" card on left | Task 2 (sidebar) + Task 3 (User compact) |
| Scrolling projects on right | Task 2 (scrollPane) |
| Scroll scene system still works | Task 1 (container listener) |
| Sticky per-project banner | Task 3b |
| Per-project contribution graph | Task 4 + Task 5 |

### Known gaps / decisions deferred

- **`project.url` format** — `getProjectCommits` assumes `github.com/owner/repo`. If any project URL is not a GitHub URL, commits returns `[]` silently (safe fallback).
- **`UserSchema.githubUsername`** — used as the GitHub author filter. If null/undefined, the timeline simply does not render (conditional guard in `project.tsx`).
- **API rate limits** — 6 projects × 1+ paginated fetches = up to 12 GitHub API calls on page load. With `VITE_GITHUB_TOKEN`, the limit is 5000/hour. Without a token, 60/hour unauthenticated. The service already reads `VITE_GITHUB_TOKEN`.
- **isoWeek date iteration** — the `bucketByWeek` loop has a safety break but the week-filling logic is approximate. For a display chart this is fine; it won't break on DST or leap years in a user-visible way.
- **`overflow: clip` browser support** — Chrome 90+, Firefox 81+, Safari 16+ (2022+). Safe for a modern portfolio site.
