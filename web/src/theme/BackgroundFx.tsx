import React, { useEffect, useState } from 'react';
import { Scene } from './scenes';

// Crossfade weights from pct (0..100 position of the viewport top within the
// "current" project's zone). At pct=0/100 both buffers sit at 50%, so the
// handoff across a project boundary is seamless.
const currentOpacity = (pct: number) => (pct < 80 ? 1 : 1 - (pct - 80) / 20);
const nextOpacity = (pct: number) => (pct < 80 ? 0 : (pct - 80) / 20);

type Pool = { slots: [string, string]; primary: 0 | 1 };

// Two persistent scene buffers crossfade between the current project's scene
// and the adjacent one being transitioned to/from. When "current" becomes
// whatever was already loaded as "next", the buffers' roles are simply
// swapped (no remount) — a slot's theme only changes while that slot is
// invisible (opacity 0).
export function BackgroundFx({
  current,
  next,
  pct,
}: {
  current: string;
  next: string;
  pct: number;
}) {
  const [pool, setPool] = useState<Pool>({ slots: [current, next], primary: 0 });

  useEffect(() => {
    setPool(prev => {
      let { slots, primary } = prev;
      let secondary: 0 | 1 = primary === 0 ? 1 : 0;
      if (slots[primary] !== current) {
        if (slots[secondary] === current) {
          primary = secondary;
          secondary = primary === 0 ? 1 : 0;
        } else {
          slots = [...slots] as [string, string];
          slots[primary] = current;
        }
      }
      if (slots[secondary] !== next) {
        slots = slots === prev.slots ? ([...slots] as [string, string]) : slots;
        slots[secondary] = next;
      }
      return slots === prev.slots && primary === prev.primary ? prev : { slots, primary };
    });
  }, [current, next]);

  const opacities: [number, number] =
    pool.primary === 0
      ? [currentOpacity(pct), nextOpacity(pct)]
      : [nextOpacity(pct), currentOpacity(pct)];

  return (
    <div className="bg-fx" aria-hidden="true">
      {pool.slots.map((theme, i) => (
        <div key={i} className="bg-fx-layer" data-theme={theme} style={{ '--layer-opacity': opacities[i] } as React.CSSProperties}>
          <Scene theme={theme} />
        </div>
      ))}
    </div>
  );
}
