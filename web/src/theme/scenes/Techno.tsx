import { CSSProperties } from 'react';

// sem 4 (api4) — techno / EVA: three parallax grids at different depths moving
// at different speeds and directions, circuit nodes pulsing, and a scanning
// beam sweeping across. Saturated purple with green accents.
function gridPath(tile: number): string {
  const min = -tile;
  const max = 1000 + tile;
  let d = '';
  for (let x = min; x <= max; x += tile) d += `M${x},${min} L${x},${max} `;
  for (let y = min; y <= max; y += tile) d += `M${min},${y} L${max},${y} `;
  return d;
}

const NODES = Array.from({ length: 9 }, () => ({
  x: 80 + Math.random() * 840,
  y: 80 + Math.random() * 840,
  d: 2.2 + Math.random() * 1.6,
}));

export function Techno() {
  return (
    <svg className="scene" viewBox="0 0 1000 1000" preserveAspectRatio="xMidYMid slice">
      <defs>
        <linearGradient id="tek-beam" x1="0" y1="0" x2="1" y2="0">
          <stop offset="0%" stopColor="var(--pink)" stopOpacity={0} />
          <stop offset="50%" stopColor="var(--pink)" stopOpacity={0.16} />
          <stop offset="100%" stopColor="var(--pink)" stopOpacity={0} />
        </linearGradient>
        <radialGradient id="tek-glow" cx="50%" cy="20%" r="60%">
          <stop offset="0%" stopColor="var(--pink)" stopOpacity={0.28} />
          <stop offset="100%" stopColor="var(--pink)" stopOpacity={0} />
        </radialGradient>
      </defs>

      <rect x={0} y={0} width={1000} height={1000} fill="url(#tek-glow)" />

      {/* far grid — slow, downward */}
      <g className="sc-drift" style={{ '--dx': '0px', '--dy': '140px', '--d': '40s' } as CSSProperties}>
        <path d={gridPath(140)} stroke="var(--border)" strokeWidth={1} fill="none" />
      </g>
      {/* mid grid — medium, leftward */}
      <g className="sc-drift" style={{ '--dx': '-84px', '--dy': '0px', '--d': '28s' } as CSSProperties}>
        <path d={gridPath(84)} stroke="var(--green)" strokeOpacity={0.35} strokeWidth={1} fill="none" />
      </g>
      {/* near grid — faster, diagonal */}
      <g className="sc-drift" style={{ '--dx': '48px', '--dy': '48px', '--d': '18s' } as CSSProperties}>
        <path d={gridPath(48)} stroke="var(--pink)" strokeOpacity={0.28} strokeWidth={1} fill="none" />
      </g>

      {NODES.map((n, i) => (
        <g
          key={i}
          className="sc-twinkle"
          style={{ '--o0': '0.15', '--o1': '0.85', '--d': `${n.d}s`, animationDelay: `${-i * 0.5}s` } as CSSProperties}
        >
          <rect x={n.x - 5} y={n.y - 5} width={10} height={10} fill="var(--green)" />
          <rect x={n.x - 11} y={n.y - 11} width={22} height={22} fill="none" stroke="var(--pink)" strokeWidth={1} strokeOpacity={0.6} />
        </g>
      ))}

      <g className="sc-scan" style={{ '--d': '7s' } as CSSProperties}>
        <rect x={-160} y={0} width={160} height={1000} fill="url(#tek-beam)" />
      </g>
    </svg>
  );
}
