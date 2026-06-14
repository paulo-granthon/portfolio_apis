import { Fragment, ReactNode } from 'react';

// Renders plain text where `backtick`-delimited spans become inline <code>
// elements. This is purely a frontend presentation concern: the stored text and
// the generated markdown keep their literal backticks untouched.
export function renderRichText(text: string): ReactNode {
  if (!text) return text;
  // Even-indexed segments are plain text; odd-indexed are inside backticks.
  return text.split('`').map((segment, index) =>
    index % 2 === 1 ? (
      <code key={index}>{segment}</code>
    ) : (
      <Fragment key={index}>{segment}</Fragment>
    ),
  );
}
