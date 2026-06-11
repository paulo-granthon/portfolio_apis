import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import ContributionList from './contributionList';

describe('ContributionList', () => {
  it('renders each contribution title, content and skills', () => {
    render(
      <ContributionList
        contributions={[
          { id: 1, title: 'Dockerization', content: 'did docker', skills: ['Docker', 'Bash'] },
        ]}
      />,
    );

    expect(screen.getByText('Dockerization')).toBeInTheDocument();
    expect(screen.getByText('did docker')).toBeInTheDocument();
    expect(screen.getByText('Docker')).toBeInTheDocument();
    expect(screen.getByText('Bash')).toBeInTheDocument();
  });

  it('shows the empty state when there are no contributions', () => {
    render(<ContributionList contributions={[]} />);
    expect(
      screen.getByText('Nenhuma contribuição especificada'),
    ).toBeInTheDocument();
  });
});
