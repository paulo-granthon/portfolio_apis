import { useEffect, useState } from "react";
import { UserSchema } from "../schemas/user";
import { getProfilePicture } from "../services/github";
import { userStyles } from "../styles";

interface UserProps {
  user: UserSchema;
}

export default function User({ user }: UserProps) {
  const [userGithubProfileUrl, setUserGithubProfileUrl] = useState<string>("");

  const userInitialSemester = user.semesterMatriculated.toString();
  const userCurrentSemester = user.semesterMatriculated
    .currentSemester()
    .toString();

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
      <img
        src={userGithubProfileUrl}
        alt="GitHub Profile"
        {...userStyles.profilePicture}
      />
      <p>{user.summary}</p>
      <div>
        <p>{userInitialSemester}</p>
      </div>
      <p>Semestre Atual: {userCurrentSemester}</p>
    </div>
  );
}
