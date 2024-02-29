import { useEffect, useState } from "react";
import { UserSchema } from "../schemas/user";
import { getProfilePicture } from "../services/github";

interface UserProps {
  user: UserSchema;
}

export default function User({ user }: UserProps) {
  const [userGithubProfileUrl, setUserGithubProfileUrl] = useState<string>("");

  const userCurrentSemester = "todo!()";

  useEffect(() => {
    getProfilePicture("paulo-granthon").then(
      (user_avatar_url: string | undefined) => {
        if (user_avatar_url && user_avatar_url.length) {
          setUserGithubProfileUrl(user_avatar_url);
        }
      },
    );
  }, []);

  return (
    <div>
      <h1>{user.name}</h1>
      <img src={userGithubProfileUrl} alt="GitHub Profile" />
      <p>{user.summary}</p>
      <div>
        <p>{user.semesterMatriculated.year}-</p>
        <p>{user.semesterMatriculated.semester}</p>
      </div>
      <p>Semestre Atual: {userCurrentSemester}</p>
    </div>
  );
}
