const API_URL = 'https://api.github.com';

export async function getProfilePicture(
  username: string,
): Promise<string | undefined> {
  if (!username || !username.length) {
    throw new Error('username validation failed');
  }

  const githubToken = import.meta.env.VITE_GITHUB_TOKEN;

  if (!githubToken || !githubToken.length) {
    throw new Error('undefined github access token');
  }

  const user = await fetch(API_URL + '/users/' + username).then(response =>
    response.json(),
  );

  if (!user) {
    throw new Error("couldn't determine user from generated url");
  }

  return user.avatar_url && user.avatar_url.length
    ? user.avatar_url
    : undefined;
}

export interface CommitEntry {
  date: string;
  sha: string;
}

export async function getProjectCommits(
  repoUrl: string,
  author: string,
): Promise<CommitEntry[]> {
  const path = repoUrl
    .replace(/^https?:\/\//, '')
    .replace(/^github\.com\//, '');
  const [owner, repo] = path.split('/');
  if (!owner || !repo) return [];

  const githubToken = import.meta.env.VITE_GITHUB_TOKEN;
  const headers: HeadersInit = githubToken
    ? { Authorization: `Bearer ${githubToken}` }
    : {};

  let page = 1;
  const commits: CommitEntry[] = [];

  while (true) {
    const url =
      `${API_URL}/repos/${owner}/${repo}/commits` +
      `?author=${encodeURIComponent(author)}&per_page=100&page=${page}`;

    const res = await fetch(url, { headers });
    if (!res.ok) break;

    const data = await res.json();
    if (!Array.isArray(data) || data.length === 0) break;

    for (const c of data) {
      const date = c.commit?.author?.date ?? c.commit?.committer?.date;
      if (date) commits.push({ date, sha: c.sha });
    }

    if (data.length < 100) break;
    page++;
  }

  return commits;
}
