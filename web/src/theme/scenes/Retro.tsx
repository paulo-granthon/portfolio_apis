import { CSSProperties } from 'react';

// sem 3 (api3) — retro sunset: a big soft sun low on the canvas, a few soft,
// turbulence-warped colour bands (strong pink + mustard), film grain, and a
// scatter of soft sparkles. Less-defined shapes, gentle motion, warm — not a
// pure-black void with hard stripes.
const BANDS = [
  { y: 300, a: 60, w: 120, fill: 'var(--green)', o: 0.16, x: 26, d: 30 },
  { y: 470, a: 44, w: 90, fill: 'var(--pink)', o: 0.14, x: -22, d: 36 },
  { y: 620, a: 70, w: 150, fill: 'var(--green)', o: 0.12, x: 18, d: 26 },
];

const SPARKS = Array.from({ length: 5 }, () => ({
  x: 80 + Math.random() * 840,
  y: 80 + Math.random() * 840,
  s: 0.8 + Math.random() * 0.8,
  d: 3 + Math.random() * 2.2,
}));

const band = (y: number, a: number) =>
  `M-160,${y} C 160,${y - a} 320,${y + a} 540,${y} C 760,${y - a} 900,${y + a} 1160,${y}`;

export function Retro() {
  return (
    <svg className="scene" viewBox="0 0 1000 1000" preserveAspectRatio="xMidYMid slice">
      <defs>
        <radialGradient id="retro-sun" cx="50%" cy="50%" r="50%">
          <stop offset="0%" stopColor="var(--green)" stopOpacity={0.5} />
          <stop offset="50%" stopColor="var(--pink)" stopOpacity={0.16} />
          <stop offset="100%" stopColor="var(--pink)" stopOpacity={0} />
        </radialGradient>
        <radialGradient id="retro-corner" cx="50%" cy="50%" r="50%">
          <stop offset="0%" stopColor="var(--pink)" stopOpacity={0.22} />
          <stop offset="100%" stopColor="var(--pink)" stopOpacity={0} />
        </radialGradient>
        <filter id="retro-warp" x="-25%" y="-25%" width="150%" height="150%">
          <feTurbulence type="fractalNoise" baseFrequency="0.003 0.006" numOctaves={1} seed={3} result="n" />
          <feDisplacementMap in="SourceGraphic" in2="n" scale={60} xChannelSelector="R" yChannelSelector="G" />
        </filter>
        <path id="spark" d="M0,-13 L2.5,-2.5 L13,0 L2.5,2.5 L0,13 L-2.5,2.5 L-13,0 L-2.5,-2.5 Z" />
      </defs>

      <circle cx={140} cy={150} r={420} fill="url(#retro-corner)" />
      <g className="sc-breathe" style={{ '--d': '18s' } as CSSProperties}>
        <circle cx={540} cy={820} r={560} fill="url(#retro-sun)" />
      </g>

      <g filter="url(#retro-warp)">
        {BANDS.map((b, i) => (
          <g
            key={i}
            className="sc-current"
            style={{ '--x': `${b.x}px`, '--d': `${b.d}s`, animationDelay: `${-i * 4}s` } as CSSProperties}
          >
            <path d={band(b.y, b.a)} fill="none" stroke={b.fill} strokeOpacity={b.o} strokeWidth={b.w} strokeLinecap="round" />
          </g>
        ))}
      </g>

      {SPARKS.map((s, i) => (
        <g
          key={i}
          className="sc-twinkle"
          style={{ '--o0': '0.1', '--o1': '0.8', '--d': `${s.d}s`, animationDelay: `${-i * 0.8}s` } as CSSProperties}
        >
          <use href="#spark" transform={`translate(${s.x} ${s.y}) scale(${s.s})`} fill="var(--green-soft)" />
        </g>
      ))}
    </svg>
  );
}
