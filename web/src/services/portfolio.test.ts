import { describe, it, expect, vi, afterEach } from 'vitest';
import { getPortfolio, mapPortfolio, portfolioMarkdownUrl } from './portfolio';
import { YearSemester } from '../schemas/yearSemester';

const rawPortfolio = {
  user: {
    id: 1,
    name: 'Paulo',
    summary: 'summary',
    semesterMatriculed: { year: 2022, semester: 2 },
    githubUsername: 'pg',
  },
  projects: [
    {
      id: 10,
      name: 'Khali',
      image: 'img',
      semester: 1,
      company: 'FATEC',
      summary: 'sum',
      description: 'desc',
      url: 'github.com/x',
      contributions: [{ id: 5, title: 't', content: 'c', skills: ['Go'] }],
    },
  ],
};

describe('portfolio service', () => {
  afterEach(() => {
    vi.restoreAllMocks();
  });

  it('maps the payload, reconstructing YearSemester and contributions', () => {
    const portfolio = mapPortfolio(rawPortfolio)!;
    expect(portfolio).toBeDefined();
    expect(portfolio.user.semesterMatriculed).toBeInstanceOf(YearSemester);
    expect(portfolio.projects).toHaveLength(1);
    expect(portfolio.projects[0].contributions[0].title).toBe('t');
    expect(portfolio.projects[0].contributions[0].skills).toEqual(['Go']);
  });

  it('returns undefined for error / empty payloads', () => {
    expect(mapPortfolio({ Error: 'boom' })).toBeUndefined();
    expect(mapPortfolio(null)).toBeUndefined();
  });

  it('getPortfolio calls the portfolio endpoint', async () => {
    const fetchMock = vi
      .fn()
      .mockResolvedValue({ json: async () => rawPortfolio });
    vi.stubGlobal('fetch', fetchMock);

    const portfolio = await getPortfolio(1);

    expect(fetchMock.mock.calls[0][0]).toContain('/portfolio/1');
    expect(portfolio?.user.name).toBe('Paulo');
  });

  it('getPortfolio returns undefined for falsy id without fetching', async () => {
    const fetchMock = vi.fn();
    vi.stubGlobal('fetch', fetchMock);

    expect(await getPortfolio(0)).toBeUndefined();
    expect(fetchMock).not.toHaveBeenCalled();
  });

  it('builds the markdown download url', () => {
    expect(portfolioMarkdownUrl(3)).toContain('/portfolio/3/markdown');
  });
});
