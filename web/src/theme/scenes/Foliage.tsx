import { CSSProperties } from 'react';

// sem 6 (api6) + api1 — a dense, dark, muted foliage mass: many overlapping
// leaves read as combined vegetation outlines rather than countable leaves,
// laced with a few vines. The whole mass is one group with a single opacity, so
// overlaps composite solidly (no see-through). Pink appears only as a few very
// shy dew highlights.
function spikyPath(y: number, a: number) {
  const step = 110;
  let d = `M-100,${y}`;
  for (let i = 0; -100 + i * step < 1100; i++) {
    const cx = -100 + (i + 0.5) * step;
    const cy = y + (i % 2 === 0 ? -1 : 1) * a * (0.6 + Math.random() * 0.8);
    d += ` Q${cx},${cy} ${-100 + (i + 1) * step},${y}`;
  }
  return d;
}

const SPIKES = Array.from({ length: 6 }, (_, i) => {
  const y = 80 + i * 165 + (Math.random() - 0.5) * 60;
  const a = 65 + Math.random() * 55;
  return {
    y, a,
    w: 100 + Math.random() * 80,
    o: 0.10 + Math.random() * 0.08,
    d: 28 + Math.random() * 16,
    x: (Math.random() - 0.5) * 14,
    path: spikyPath(y, a),
  };
});

const span = () => -120 + Math.random() * 1240;
const FILLED = Array.from({ length: 60 }, () => ({ x: span(), y: span(), s: 0.5 + Math.random() * 1.3, rot: Math.random() * 360 }));
const OUTLINE = Array.from({ length: 26 }, () => ({ x: span(), y: span(), s: 0.55 + Math.random() * 1.1, rot: Math.random() * 360 }));
const DEW = Array.from({ length: 5 }, () => ({ x: Math.random() * 1000, y: Math.random() * 1000, rr: 1.4 + Math.random() * 2, d: 2.6 + Math.random() * 2.4 }));

const VINES = [
  'M-60,260 C 180,200 240,420 460,360 S 760,180 1080,300',
  'M-40,640 C 220,720 360,520 560,640 S 880,800 1080,660',
  'M120,-40 C 60,220 300,320 220,560 S 360,860 260,1080',
];

export function Foliage() {
  return (
    <svg className="scene" viewBox="0 0 1000 1000" preserveAspectRatio="xMidYMid slice">
      <defs>
        <path id="leaf2" d="M0,0 C24,-30 64,-28 100,0 C64,28 24,30 0,0 Z" />
        <radialGradient id="fol-dark" cx="50%" cy="40%" r="65%">
          <stop offset="0%" stopColor="var(--bg)" stopOpacity={0} />
          <stop offset="100%" stopColor="var(--bg)" stopOpacity={0.7} />
        </radialGradient>
      </defs>

      {SPIKES.map((s, i) => (
        <g
          key={i}
          className="sc-current"
          style={{ '--x': `${s.x}px`, '--d': `${s.d}s`, animationDelay: `${-i * 3.2}s` } as CSSProperties}
        >
          <path
            d={s.path}
            fill="none"
            stroke="var(--green)"
            strokeOpacity={s.o}
            strokeWidth={s.w}
            strokeLinecap="butt"
          />
        </g>
      ))}

      <g className="sc-current" style={{ '--x': '18px', '--y': '-12px', '--d': '34s' } as CSSProperties}>
        <g opacity={0.4}>
          {FILLED.map((l, i) => (
            <use key={i} href="#leaf2" transform={`translate(${l.x} ${l.y}) rotate(${l.rot}) scale(${l.s})`} fill="var(--green)" />
          ))}
          {VINES.map((d, i) => (
            <path key={i} d={d} fill="none" stroke="var(--green)" strokeWidth={3} strokeLinecap="round" />
          ))}
          {OUTLINE.map((l, i) => (
            <use
              key={i}
              href="#leaf2"
              transform={`translate(${l.x} ${l.y}) rotate(${l.rot}) scale(${l.s})`}
              fill="none"
              stroke="var(--green-soft)"
              strokeWidth={1.4}
            />
          ))}
        </g>
      </g>

      <rect x={0} y={0} width={1000} height={1000} fill="url(#fol-dark)" />

      {DEW.map((d, i) => (
        <circle
          key={i}
          className="sc-twinkle"
          style={{ '--o0': '0.15', '--o1': '0.7', '--d': `${d.d}s`, animationDelay: `${-i * 0.7}s` } as CSSProperties}
          cx={d.x}
          cy={d.y}
          r={d.rr}
          fill="var(--pink-soft)"
        />
      ))}
    </svg>
  );
}
