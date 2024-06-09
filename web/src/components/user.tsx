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
    <div {...userStyles.card}>
      <div {...userStyles.cardLeft}>
        <img
          alt="GitHub Profile Picture"
          src={userGithubProfileUrl}
          {...userStyles.cardLeftPicture}
        />
      </div>
      <div {...userStyles.cardRight}>
        <div {...userStyles.cardRightHeader}>
          <h2>{user.name}</h2>
          <div {...userStyles.cardRightSemester}>
            <p>Semestre Atual: {userCurrentSemester}</p>
          </div>
        </div>
        <p {...userStyles.cardRightSummary}>{user.summary}</p>
      </div>
    </div>
  );
}
