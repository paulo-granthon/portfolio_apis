import { useEffect, useState } from 'react';

export type SceneState = {
  // most-present project at the true viewport centre — drives the palette
  active: string | null;
  // theme assigned to the "current" background buffer
  current: string;
  // theme assigned to the "next" background buffer, ready to crossfade in
  next: string;
  // 0..100 position of the viewport top within the current project's zone
  pct: number;
};

export function useScrollScenes(
  dep: unknown,
  container?: HTMLElement | null,
): SceneState {
  const [state, setState] = useState<SceneState>({
    active: null,
    current: 'default',
    next: 'default',
    pct: 0,
  });

  useEffect(() => {
    const els = Array.from(
      document.querySelectorAll<HTMLElement>('[data-project]'),
    );
    if (els.length === 0) return;

    let raf = 0;
    const presenceAt = (probe: number, top: number, bottom: number, fade: number) => {
      const outside = Math.max(0, top - probe, probe - bottom);
      return Math.max(0, 1 - outside / fade);
    };

    const compute = () => {
      raf = 0;
      const h = window.innerHeight;
      const center = h / 2;
      const fade = h * 0.15;

      const rects = els.map(el => ({
        name: el.dataset.project ?? '',
        rect: el.getBoundingClientRect(),
      }));

      // palette: most-present project at the true centre (so colour leads)
      let active: string | null = null;
      let best = 0;
      for (const { name, rect } of rects) {
        const p = presenceAt(center, rect.top, rect.bottom, fade);
        if (p > best) {
          best = p;
          active = name;
        }
      }

      // background buffers: "current" is the project whose top has scrolled
      // to/past the viewport top (or 'default' for the header area above
      // project 1); pct is how far through that project's zone we are.
      // "next" is the adjacent zone whose crossfade we're approaching —
      // the previous one in the first half, the following one in the second.
      let idx = -1;
      for (let i = 0; i < rects.length; i++) {
        if (rects[i].rect.top <= 0) idx = i; else break;
      }

      let current: string;
      let zoneTop: number;
      let zoneBottom: number;
      if (idx === -1) {
        current = 'default';
        zoneBottom = rects[0].rect.top;
        zoneTop = zoneBottom - h;
      } else {
        current = rects[idx].name;
        zoneTop = rects[idx].rect.top;
        zoneBottom = rects[idx].rect.bottom;
      }
      const pct = Math.min(100, Math.max(0, (-zoneTop / (zoneBottom - zoneTop)) * 100));

      const all = ['default', ...rects.map(r => r.name)];
      const ci = all.indexOf(current);
      const next =
        pct < 50
          ? ci > 0 ? all[ci - 1] : current
          : ci < all.length - 1 ? all[ci + 1] : current;

      setState({ active: best > 0 ? active : null, current, next, pct });
    };

    const onScroll = () => {
      if (!raf) raf = requestAnimationFrame(compute);
    };

    compute();
    const target: EventTarget = container ?? window;
    target.addEventListener('scroll', onScroll, { passive: true });
    window.addEventListener('resize', onScroll);
    return () => {
      target.removeEventListener('scroll', onScroll);
      window.removeEventListener('resize', onScroll);
      if (raf) cancelAnimationFrame(raf);
    };
  }, [dep, container]);

  return state;
}
