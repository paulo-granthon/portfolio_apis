import { FC, memo } from 'react';
import { Foliage } from './Foliage';
import { Bubbles } from './Bubbles';
import { Retro } from './Retro';
import { Techno } from './Techno';
import { Ocean } from './Ocean';
import { Grid } from './Grid';

// One generative scene per project; unknown themes fall back to Foliage.
const SCENES: Record<string, FC> = {
  default: Grid,
  Khali: Foliage,
  API2Semestre: Bubbles,
  api3: Retro,
  api4: Techno,
  api5: Ocean,
  api6: Foliage,
};

// memo: theme is constant per layer, so the SVG never re-renders when the
// parent re-renders every scroll frame — only the wrapper opacity changes.
export const Scene = memo(function Scene({ theme }: { theme: string }) {
  const Component = SCENES[theme] ?? Foliage;
  return <Component />;
});
