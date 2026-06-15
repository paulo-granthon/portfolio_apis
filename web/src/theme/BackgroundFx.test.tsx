import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/react';
import { BackgroundFx } from './BackgroundFx';

const layerOpacity = (el: HTMLElement) => el.style.getPropertyValue('--layer-opacity');

describe('BackgroundFx', () => {
  it('keeps the 2-buffer pool stable across a project boundary crossing', () => {
    const { container, rerender } = render(<BackgroundFx current="default" next="default" pct={0} />);
    let layers = container.querySelectorAll<HTMLElement>('.bg-fx-layer');
    expect(layers).toHaveLength(2);
    expect(layers[0].dataset.theme).toBe('default');
    expect(layers[1].dataset.theme).toBe('default');
    // pct=0 < 80: current=1, next=0
    expect(layerOpacity(layers[0])).toBe('1');
    expect(layerOpacity(layers[1])).toBe('0');

    // mid-zone: secondary slot preloads Khali while invisible
    rerender(<BackgroundFx current="default" next="Khali" pct={50} />);
    layers = container.querySelectorAll<HTMLElement>('.bg-fx-layer');
    expect(layers[0].dataset.theme).toBe('default');
    expect(layers[1].dataset.theme).toBe('Khali');
    expect(layerOpacity(layers[0])).toBe('1');
    expect(layerOpacity(layers[1])).toBe('0');

    // transition zone midpoint pct=90: both at 50%
    rerender(<BackgroundFx current="default" next="Khali" pct={90} />);
    expect(layerOpacity(container.querySelectorAll<HTMLElement>('.bg-fx-layer')[0])).toBe('0.5');
    expect(layerOpacity(container.querySelectorAll<HTMLElement>('.bg-fx-layer')[1])).toBe('0.5');

    // at boundary pct=100: fully transitioned (current=0, next=1)
    rerender(<BackgroundFx current="default" next="Khali" pct={100} />);
    expect(layerOpacity(container.querySelectorAll<HTMLElement>('.bg-fx-layer')[0])).toBe('0');
    expect(layerOpacity(container.querySelectorAll<HTMLElement>('.bg-fx-layer')[1])).toBe('1');

    // crossing into Khali: roles flip, no remount, transition complete immediately
    rerender(<BackgroundFx current="Khali" next="default" pct={0} />);
    layers = container.querySelectorAll<HTMLElement>('.bg-fx-layer');
    expect(layers[0].dataset.theme).toBe('default');
    expect(layers[1].dataset.theme).toBe('Khali');
    expect(layerOpacity(layers[0])).toBe('0');
    expect(layerOpacity(layers[1])).toBe('1');
  });
});
