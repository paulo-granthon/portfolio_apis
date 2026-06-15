import { CSSProperties } from 'react';

// default zone — a quiet geometric grid: fine guide-lines, a major grid at 5×
// spacing, and a few centred composition elements. All tones are near-neutral
// (rgba white at very low opacity) so the scene works over any palette's bg.
export function Grid() {
  return (
    <svg className="scene" viewBox="0 0 1000 1000" preserveAspectRatio="xMidYMid slice">
      <defs>
        <pattern id="grid-fine" width="50" height="50" patternUnits="userSpaceOnUse">
          <path d="M 50 0 L 0 0 0 50" fill="none" stroke="rgba(255,255,255,0.045)" strokeWidth="0.8" />
        </pattern>
        <pattern id="grid-major" width="250" height="250" patternUnits="userSpaceOnUse">
          <path d="M 250 0 L 0 0 0 250" fill="none" stroke="rgba(255,255,255,0.09)" strokeWidth="1" />
        </pattern>
      </defs>

      <rect width="1000" height="1000" fill="url(#grid-fine)" />
      <rect width="1000" height="1000" fill="url(#grid-major)" />

      <g
        className="sc-breathe"
        style={{ '--d': '24s' } as CSSProperties}
        stroke="rgba(255,255,255,0.07)"
        fill="none"
        strokeWidth="1"
      >
        <circle cx="500" cy="500" r="290" />
        <circle cx="500" cy="500" r="185" />
        <rect x="92" y="92" width="56" height="56" />
        <rect x="852" y="92" width="56" height="56" />
        <rect x="92" y="852" width="56" height="56" />
        <rect x="852" y="852" width="56" height="56" />
        <line x1="500" y1="415" x2="500" y2="585" stroke="rgba(255,255,255,0.09)" />
        <line x1="415" y1="500" x2="585" y2="500" stroke="rgba(255,255,255,0.09)" />
      </g>

      <circle cx="500" cy="500" r="3" fill="rgba(255,255,255,0.18)" />
    </svg>
  );
}
