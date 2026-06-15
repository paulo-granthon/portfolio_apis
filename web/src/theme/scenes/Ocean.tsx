import { CSSProperties } from 'react';

// sem 5 (api5) — deep ocean: a vertical depth gradient (darker the deeper),
// many thick, soft, turbulence-warped current bands, a rusty sunken glow, and a
// dense field of drifting dust/star particles. Slow and heavy, not "beachy".
const BANDS = Array.from({ length: 20 }, (_, i) => ({
  y: 40 + i * 48 + (Math.random() - 0.5) * 20,
  a: 3 + Math.random() * 7,
  w: 20 + Math.random() * 45,
  o: 0.08 + Math.random() * 0.1,
  x: (Math.random() - 0.5) * 30,
  d: 40 + Math.random() * 24,
}));
const PARTICLES = Array.from({ length: 34 }, () => ({
  x: Math.random() * 1000,
  y: Math.random() * 1000,
  rr: 1.1 + Math.random() * 2.1,
  d: 2.8 + Math.random() * 2.6,
  bright: Math.random() > 0.7,
}));

const band = (y: number, a: number) =>
  `M-140,${y} C 130,${y - a} 300,${y + a} 500,${y} C 720,${y - a} 880,${y + a} 1140,${y}`;

export function Ocean() {
  return (
    <svg className="scene" viewBox="0 0 1000 1000" preserveAspectRatio="xMidYMid slice">
      <defs>
        <linearGradient id="ocean-depth" x1="0" y1="0" x2="0" y2="1">
          <stop offset="0%" stopColor="var(--green)" stopOpacity={0.07} />
          <stop offset="40%" stopColor="var(--bg)" stopOpacity={0.5} />
          <stop offset="100%" stopColor="#01030a" stopOpacity={0.96} />
        </linearGradient>
        <filter id="ocean-turb" x="-20%" y="-20%" width="140%" height="140%">
          <feTurbulence type="fractalNoise" baseFrequency="0.0035 0.006" numOctaves={1} seed={11} result="n" />
          <feDisplacementMap in="SourceGraphic" in2="n" scale={30} xChannelSelector="R" yChannelSelector="G" />
        </filter>
        <radialGradient id="ocean-sun" cx="50%" cy="50%" r="50%">
          <stop offset="0%" stopColor="var(--pink)" stopOpacity={0.18} />
          <stop offset="100%" stopColor="var(--pink)" stopOpacity={0} />
        </radialGradient>
      </defs>

      <rect x={0} y={0} width={1000} height={1000} fill="url(#ocean-depth)" />
      <circle cx={770} cy={910} r={460} fill="url(#ocean-sun)" />

      <g filter="url(#ocean-turb)">
        {BANDS.map((b, i) => (
          <g
            key={i}
            className="sc-current"
            style={{ '--x': `${b.x}px`, '--d': `${b.d}s`, animationDelay: `${-i * 2.5}s` } as CSSProperties}
          >
            <path
              d={band(b.y, b.a)}
              fill="none"
              stroke="var(--green)"
              strokeOpacity={b.o}
              strokeWidth={b.w}
              strokeLinecap="round"
            />
          </g>
        ))}
      </g>

      {PARTICLES.map((p, i) => (
        <circle
          key={i}
          className="sc-twinkle"
          style={{ '--o0': '0.08', '--o1': p.bright ? '0.8' : '0.5', '--d': `${p.d}s`, animationDelay: `${-i * 0.3}s` } as CSSProperties}
          cx={p.x}
          cy={p.y}
          r={p.rr}
          fill={p.bright ? 'var(--pink-soft)' : 'var(--green-soft)'}
        />
      ))}
    </svg>
  );
}
