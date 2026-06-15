// Per-project color palettes, applied to :root by the scroll observer so the
// whole page morphs into each project's aesthetic. Keyed by project name.
// Values map directly to the animatable CSS custom properties in index.css.

export type Palette = Record<string, string>;

export const PALETTE_VARS = [
  '--bg',
  '--bg-elev',
  '--bg-elev-2',
  '--border',
  '--border-soft',
  '--text',
  '--text-dim',
  '--text-faint',
  '--pink',
  '--pink-soft',
  '--green',
  '--green-soft',
  '--pink-glow',
  '--green-glow',
  '--glow-1',
  '--glow-2',
  '--vignette',
] as const;

export const PALETTES: Record<string, Palette> = {
  // Semester 1 — pastel soft pink + mature green over an almost-black green.
  Khali: {
    '--bg': '#0a0f0c',
    '--bg-elev': '#111a15',
    '--bg-elev-2': '#16211b',
    '--border': '#243029',
    '--border-soft': '#18221c',
    '--text': '#ecf4ee',
    '--text-dim': '#a4b8ac',
    '--text-faint': '#62756a',
    '--pink': '#f7a8c4',
    '--pink-soft': '#ffc8db',
    '--green': '#4fb393',
    '--green-soft': '#8fdcc2',
    '--pink-glow': 'rgba(247, 168, 196, 0.30)',
    '--green-glow': 'rgba(79, 179, 147, 0.26)',
    '--glow-1': 'rgba(247, 168, 196, 0.14)',
    '--glow-2': 'rgba(79, 179, 147, 0.12)',
    '--vignette': 'rgba(10, 15, 12, 0.86)',
  },
  // Semester 2 — partner branding: bright, fun shades of blue with white text.
  API2Semestre: {
    '--bg': '#0a0e16',
    '--bg-elev': '#111726',
    '--bg-elev-2': '#172234',
    '--border': '#243349',
    '--border-soft': '#18222f',
    '--text': '#eaf2fb',
    '--text-dim': '#9fb6d3',
    '--text-faint': '#5d7290',
    '--pink': '#3a9bff',
    '--pink-soft': '#7dbdff',
    '--green': '#58d6ff',
    '--green-soft': '#9fe7ff',
    '--pink-glow': 'rgba(58, 155, 255, 0.34)',
    '--green-glow': 'rgba(88, 214, 255, 0.30)',
    '--glow-1': 'rgba(58, 155, 255, 0.18)',
    '--glow-2': 'rgba(88, 214, 255, 0.13)',
    '--vignette': 'rgba(10, 14, 22, 0.86)',
  },
  // Semester 3 — retro high-contrast: black, strong pink, mustard yellow, white.
  api3: {
    '--bg': '#050505',
    '--bg-elev': '#0f0e0d',
    '--bg-elev-2': '#151311',
    '--border': '#2c241f',
    '--border-soft': '#1c1815',
    '--text': '#fdf6ea',
    '--text-dim': '#c0b1a0',
    '--text-faint': '#6f6256',
    '--pink': '#ff2e88',
    '--pink-soft': '#ff6fab',
    '--green': '#e8b62c',
    '--green-soft': '#f4d06a',
    '--pink-glow': 'rgba(255, 46, 136, 0.40)',
    '--green-glow': 'rgba(232, 182, 44, 0.32)',
    '--glow-1': 'rgba(255, 46, 136, 0.20)',
    '--glow-2': 'rgba(232, 182, 44, 0.14)',
    '--vignette': 'rgba(5, 5, 5, 0.90)',
  },
  // Semester 4 — saturated cartoon/techno purple + a bit of green (EVA-01).
  api4: {
    '--bg': '#0e0a1a',
    '--bg-elev': '#17102e',
    '--bg-elev-2': '#1e163b',
    '--border': '#352559',
    '--border-soft': '#1d1638',
    '--text': '#efe9ff',
    '--text-dim': '#b4a7d8',
    '--text-faint': '#70609f',
    '--pink': '#a64dff',
    '--pink-soft': '#c48cff',
    '--green': '#4dff9e',
    '--green-soft': '#93ffc6',
    '--pink-glow': 'rgba(166, 77, 255, 0.40)',
    '--green-glow': 'rgba(77, 255, 158, 0.28)',
    '--glow-1': 'rgba(166, 77, 255, 0.22)',
    '--glow-2': 'rgba(77, 255, 158, 0.12)',
    '--vignette': 'rgba(14, 10, 26, 0.85)',
  },
  // Semester 5 — oceanic: dark ocean-night blue + mid sunset orange.
  api5: {
    '--bg': '#07111c',
    '--bg-elev': '#0e1c2b',
    '--bg-elev-2': '#132536',
    '--border': '#213a53',
    '--border-soft': '#122130',
    '--text': '#e9f2f8',
    '--text-dim': '#99b4c7',
    '--text-faint': '#587583',
    '--pink': '#ff8a4c',
    '--pink-soft': '#ffb083',
    '--green': '#2f9fd4',
    '--green-soft': '#74c7e8',
    '--pink-glow': 'rgba(255, 138, 76, 0.32)',
    '--green-glow': 'rgba(47, 159, 212, 0.30)',
    '--glow-1': 'rgba(255, 138, 76, 0.16)',
    '--glow-2': 'rgba(47, 159, 212, 0.18)',
    '--vignette': 'rgba(7, 17, 28, 0.87)',
  },
  // Semester 6 — green + purple. This is the default/global palette.
  api6: {
    '--bg': '#0a0c0e',
    '--bg-elev': '#11151a',
    '--bg-elev-2': '#161b21',
    '--border': '#232a32',
    '--border-soft': '#1a2027',
    '--text': '#e8eef2',
    '--text-dim': '#9aa6b1',
    '--text-faint': '#5b6672',
    '--pink': '#ff3d81',
    '--pink-soft': '#ff7eae',
    '--green': '#2ee6a6',
    '--green-soft': '#7df0ca',
    '--pink-glow': 'rgba(255, 61, 129, 0.35)',
    '--green-glow': 'rgba(46, 230, 166, 0.30)',
    '--glow-1': 'rgba(255, 61, 129, 0.16)',
    '--glow-2': 'rgba(46, 230, 166, 0.12)',
    '--vignette': 'rgba(10, 12, 14, 0.86)',
  },
};

// Apply a project's palette to :root, or clear back to the default (the values
// declared in index.css) when name is null/unknown.
export function applyPalette(name: string | null): void {
  const root = document.documentElement;
  const palette = name ? PALETTES[name] : undefined;
  if (palette) {
    for (const [key, value] of Object.entries(palette)) {
      root.style.setProperty(key, value);
    }
  } else {
    for (const key of PALETTE_VARS) {
      root.style.removeProperty(key);
    }
  }
}
