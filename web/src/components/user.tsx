import { UserSchema } from '../schemas/user';
import * as userStyles from '../styles/user';

interface UserProps {
  user: UserSchema;
}

export default function User({ user }: UserProps) {
  const userCurrentSemester = user.semesterMatriculed
    .currentSemester()
    .toString();

  const avatarUrl = user.githubUsername
    ? `https://github.com/${user.githubUsername}.png?size=240`
    : undefined;

  return (
    <div {...userStyles.card}>
      <div {...userStyles.cardLeft}>
        <img {...userStyles.cardLeftPicture} alt={user.name} src={avatarUrl} />
      </div>
      <div {...userStyles.cardRight}>
        <div {...userStyles.cardRightHeader}>
          <h2>{user.name}</h2>
          <span {...userStyles.cardRightSemester}>
            semestre atual · {userCurrentSemester}
          </span>
        </div>
        <p {...userStyles.cardRightSummary}>{user.summary}</p>
      </div>
    </div>
  );
}
