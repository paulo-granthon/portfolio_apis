import { useEffect, useState } from "react";
import { UserSchema } from "../schemas/user";
import { getProfilePicture } from "../services/github";
import * as userStyles from "../styles/user";

interface UserProps {
  user: UserSchema;
}

export default function User({ user }: UserProps) {
  const [userGithubProfileUrl, setUserGithubProfileUrl] = useState<string>("");

  const userInitialSemester = user.semesterMatriculed.toString();
  const userCurrentSemester = user.semesterMatriculed
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
    <div {...userStyles.userCard}>
      <div {...userStyles.userCardLeft}>
        <img
          src={userGithubProfileUrl}
          alt="GitHub Profile"
          {...userStyles.profilePicture}
        />
      </div>
      <div {...userStyles.userCardRight}>
        <h1>{user.name}</h1>
        <div {...userStyles.userCardSemester}>
          <p>{userInitialSemester}</p>
          <p>Semestre Atual: {userCurrentSemester}</p>
        </div>
        <p>{user.summary}</p>
      </div>
    </div>
  );
}
