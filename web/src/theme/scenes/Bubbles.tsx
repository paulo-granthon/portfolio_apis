import { CSSProperties } from 'react';

// sem 2 (API2Semestre) — playful, flat, abstract: irregular rounded blobs of
// many sizes and blue tones (plus a shy soft pink), densely filling the canvas
// over a soft gradient wash. Flat lighting (no glossy highlights), and they
// wobble gently in place instead of sliding in a direction. One group opacity,
// so overlaps stay flat (no see-through).
const mid = (p: number[], q: number[]) => [(p[0] + q[0]) / 2, (p[1] + q[1]) / 2];

function blobPath(cx: number, cy: number, radius: number): string {
  const n = 7;
  const pts = Array.from({ length: n }, (_, i) => {
    const a = (i / n) * Math.PI * 2;
    const rad = radius * (0.7 + Math.random() * 0.55);
    return [cx + Math.cos(a) * rad, cy + Math.sin(a) * rad];
  });
  let d = `M ${mid(pts[n - 1], pts[0]).join(' ')} `;
  for (let i = 0; i < n; i++) {
    const m = mid(pts[i], pts[(i + 1) % n]);
    d += `Q ${pts[i][0]} ${pts[i][1]} ${m[0]} ${m[1]} `;
  }
  return d + 'Z';
}

const TONES = ['var(--green)', 'var(--pink)', 'var(--green-soft)', 'var(--pink)', 'rgba(255, 150, 200, 0.9)', 'var(--green)'];
const BLOBS = Array.from({ length: 17 }, () => {
  const cx = Math.random() * 1000;
  const cy = Math.random() * 1000;
  const radius = 90 + Math.random() * 190;
  return {
    cx, cy, radius,
    path: blobPath(cx, cy, radius),
    tone: TONES[Math.floor(Math.random() * TONES.length)],
    dx: (Math.random() - 0.5) * 26,
    dy: (Math.random() - 0.5) * 26,
    dur: 12 + Math.random() * 12,
  };
});

export function Bubbles() {
  return (
    <svg className="scene" viewBox="0 0 1000 1000" preserveAspectRatio="xMidYMid slice">
      <defs>
        <radialGradient id="bub-wash-a" cx="30%" cy="28%" r="60%">
          <stop offset="0%" stopColor="var(--green)" stopOpacity={0.22} />
          <stop offset="100%" stopColor="var(--green)" stopOpacity={0} />
        </radialGradient>
        <radialGradient id="bub-wash-b" cx="76%" cy="74%" r="60%">
          <stop offset="0%" stopColor="var(--pink)" stopOpacity={0.22} />
          <stop offset="100%" stopColor="var(--pink)" stopOpacity={0} />
        </radialGradient>
      </defs>

      <rect x={0} y={0} width={1000} height={1000} fill="url(#bub-wash-a)" />
      <rect x={0} y={0} width={1000} height={1000} fill="url(#bub-wash-b)" />

      <g opacity={0.55}>
        {BLOBS.map((b, i) => (
          <g
            key={i}
            className="sc-current"
            style={{ '--x': `${b.dx}px`, '--y': `${b.dy}px`, '--d': `${b.dur}s`, animationDelay: `${-i * 0.9}s` } as CSSProperties}
          >
            <path d={b.path} fill={b.tone} />
          </g>
        ))}
      </g>
    </svg>
  );
}
