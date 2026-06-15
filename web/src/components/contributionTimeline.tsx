import { useEffect, useState } from 'react';
import { getProjectCommits, CommitEntry } from '../services/github';
import * as s from '../styles/contributionTimeline';

interface Props {
  repoUrl: string;
  author: string;
}

function isoWeek(date: Date): string {
  const d = new Date(Date.UTC(date.getUTCFullYear(), date.getUTCMonth(), date.getUTCDate()));
  const day = d.getUTCDay() || 7;
  d.setUTCDate(d.getUTCDate() + 4 - day); // shift to Thursday of the same ISO week
  const year = d.getUTCFullYear();
  const startOfYear = new Date(Date.UTC(year, 0, 1));
  const week = Math.ceil(((+d - +startOfYear) / 86400000 + 1) / 7);
  return `${year}-W${String(week).padStart(2, '0')}`;
}

function bucketByWeek(commits: CommitEntry[]): { week: string; count: number }[] {
  const map = new Map<string, number>();
  for (const c of commits) {
    const w = isoWeek(new Date(c.date));
    map.set(w, (map.get(w) ?? 0) + 1);
  }
  if (map.size === 0) return [];

  const result: { week: string; count: number }[] = [];
  const firstCommitDate = commits
    .map(c => new Date(c.date))
    .reduce((a, b) => (a < b ? a : b));
  const lastCommitDate = commits
    .map(c => new Date(c.date))
    .reduce((a, b) => (a > b ? a : b));

  const cur = new Date(
    Date.UTC(firstCommitDate.getUTCFullYear(), firstCommitDate.getUTCMonth(), firstCommitDate.getUTCDate()),
  );
  // align cur to Monday of its week
  const wd = cur.getUTCDay() || 7;
  cur.setUTCDate(cur.getUTCDate() - (wd - 1));

  const lastWeek = isoWeek(lastCommitDate);

  for (let iterations = 0; iterations < 1000; iterations++) {
    const w = isoWeek(cur);
    result.push({ week: w, count: map.get(w) ?? 0 });
    if (w >= lastWeek) break;
    cur.setUTCDate(cur.getUTCDate() + 7);
  }

  return result;
}

export function ContributionTimeline({ repoUrl, author }: Props) {
  const [buckets, setBuckets] = useState<{ week: string; count: number }[] | null>(null);

  useEffect(() => {
    let cancelled = false;
    getProjectCommits(repoUrl, author).then(commits => {
      if (!cancelled) setBuckets(bucketByWeek(commits));
    });
    return () => { cancelled = true; };
  }, [repoUrl, author]);

  if (buckets === null) return null;
  if (buckets.length === 0) return <p {...s.empty}>sem commits registrados</p>;

  const max = Math.max(...buckets.map(b => b.count), 1);

  return (
    <div {...s.root}>
      <p {...s.label}>contribuições por semana</p>
      <div {...s.chart}>
        {buckets.map(b => (
          <div
            key={b.week}
            {...s.bar}
            style={{ height: `${Math.max(4, (b.count / max) * 40)}px` }}
            title={`${b.week}: ${b.count} commit${b.count !== 1 ? 's' : ''}`}
          />
        ))}
      </div>
    </div>
  );
}
