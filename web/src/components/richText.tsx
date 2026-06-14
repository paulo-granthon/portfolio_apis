import { Fragment, ReactNode } from 'react';

// Matches, in priority order: a linked code reference [`Symbol`](url); a plain
// markdown link [text](url); or a bare `backtick` code span.
const TOKEN =
  /\[`([^`]+)`\]\(([^)]+)\)|\[([^\]`]+)\]\(([^)]+)\)|`([^`]+)`/g;

// Renders plain text with light markdown: `backtick` spans become inline
// <code>, [`Symbol`](url) becomes a <code> linking to its definition on GitHub,
// and [text](url) becomes a plain link (used for PR / repo mentions). All open
// in a new tab. This is presentation-only: the stored text and the generated
// markdown keep their literal markdown, which GitHub renders the same way.
export function renderRichText(text: string): ReactNode {
  if (!text) return text;

  const nodes: ReactNode[] = [];
  let lastIndex = 0;
  let key = 0;

  for (const match of text.matchAll(TOKEN)) {
    const index = match.index ?? 0;
    if (index > lastIndex) {
      nodes.push(
        <Fragment key={key++}>{text.slice(lastIndex, index)}</Fragment>,
      );
    }

    const [, codeLabel, codeHref, linkText, linkHref, plainCode] = match;
    if (codeLabel !== undefined) {
      nodes.push(
        <a
          key={key++}
          className="codelink"
          href={codeHref}
          target="_blank"
          rel="noreferrer"
        >
          <code>{codeLabel}</code>
        </a>,
      );
    } else if (linkText !== undefined) {
      nodes.push(
        <a key={key++} href={linkHref} target="_blank" rel="noreferrer">
          {linkText}
        </a>,
      );
    } else {
      nodes.push(<code key={key++}>{plainCode}</code>);
    }

    lastIndex = index + match[0].length;
  }

  if (lastIndex < text.length) {
    nodes.push(<Fragment key={key++}>{text.slice(lastIndex)}</Fragment>);
  }

  return nodes;
}
