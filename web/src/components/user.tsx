import { UserSchema } from '../schemas/user';
import * as userStyles from '../styles/user';

interface UserProps {
  user: UserSchema;
  compact?: boolean;
}

export default function User({ user, compact }: UserProps) {
  const userCurrentSemester = user.semesterMatriculed
    .currentSemester()
    .toString();

  const avatarUrl = user.githubUsername
    ? `https://github.com/${user.githubUsername}.png?size=240`
    : undefined;

  return (
    <div {...(compact ? userStyles.cardCompact : userStyles.card)}>
      <div {...userStyles.cardLeft}>
        <img
          {...(compact ? userStyles.cardLeftPictureCompact : userStyles.cardLeftPicture)}
          alt={user.name}
          src={avatarUrl}
        />
      </div>
      <div {...(compact ? userStyles.cardRightCompact : userStyles.cardRight)}>
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
