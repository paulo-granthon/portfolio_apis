import { CSSProperties } from 'react';

// sem 1 (Khali) — soft pastel: two distant glows behind a denser drift of
// leaves/petals. The leaves are one group with a single opacity (no see-through
// on overlap). The group drifts along a dominant diagonal while each leaf also
// sways, so the motion reads organic rather than a straight slide.
const LEAVES = Array.from({ length: 34 }, () => ({
  x: -120 + Math.random() * 1240,
  y: -120 + Math.random() * 1240,
  s: 0.5 + Math.random() * 1.0,
  rot: Math.random() * 360,
  green: Math.random() > 0.5,
  d: 7 + Math.random() * 6,
  sway: 5 + Math.random() * 7,
}));

export function Nature() {
  return (
    <svg className="scene" viewBox="0 0 1000 1000" preserveAspectRatio="xMidYMid slice">
      <defs>
        <radialGradient id="nat-glow-a" cx="50%" cy="50%" r="50%">
          <stop offset="0%" stopColor="var(--green)" stopOpacity={0.2} />
          <stop offset="100%" stopColor="var(--green)" stopOpacity={0} />
        </radialGradient>
        <radialGradient id="nat-glow-b" cx="50%" cy="50%" r="50%">
          <stop offset="0%" stopColor="var(--pink)" stopOpacity={0.22} />
          <stop offset="100%" stopColor="var(--pink)" stopOpacity={0} />
        </radialGradient>
        <path id="petal2" d="M0,0 C20,-28 56,-26 84,0 C56,26 20,28 0,0 Z" />
      </defs>

      <g className="sc-breathe" style={{ '--d': '22s' } as CSSProperties}>
        <circle cx={250} cy={300} r={440} fill="url(#nat-glow-a)" />
      </g>
      <g className="sc-breathe" style={{ '--d': '28s', animationDelay: '-9s' } as CSSProperties}>
        <circle cx={800} cy={760} r={480} fill="url(#nat-glow-b)" />
      </g>

      <g className="sc-current" style={{ '--x': '-46px', '--y': '-32px', '--d': '30s' } as CSSProperties}>
        <g opacity={0.5}>
          {LEAVES.map((l, i) => (
            <g key={i} transform={`translate(${l.x} ${l.y}) rotate(${l.rot}) scale(${l.s})`}>
              <g className="sc-sway" style={{ '--d': `${l.d}s`, '--r': `${l.sway}deg` } as CSSProperties}>
                <use href="#petal2" fill={l.green ? 'var(--green-soft)' : 'var(--pink-soft)'} />
              </g>
            </g>
          ))}
        </g>
      </g>
    </svg>
  );
}
