const API_URL = "https://api.github.com";

export async function getProfilePicture(
  username: string,
): Promise<string | undefined> {
  if (!username || !username.length) {
    throw new Error("username validation failed");
  }

  const githubToken = import.meta.env.VITE_GITHUB_TOKEN;

  if (!githubToken || !githubToken.length) {
    throw new Error("undefined github access token");
  }

  const user = await fetch(API_URL + "/users/" + username).then((response) =>
    response.json(),
  );

  if (!user) {
    throw new Error("couldn't determine user from generated url");
  }

  return user.avatar_url && user.avatar_url.length
    ? user.avatar_url
    : undefined;
}
