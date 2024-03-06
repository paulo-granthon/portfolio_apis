import User from "../components/user";
import ProjectList from "../components/projectList";
import { UserSchema } from "../schemas/user";
import * as styles from "../styles/portfolio";

interface PortfolioProps {
  user: UserSchema;
}

export default function Portfolio({user}: PortfolioProps) {
  return (
    <div {...styles.portfolio}>
      <div>
        <h1>Portfolio</h1>
      </div>
      <User user={user} />
      <ProjectList user={user} />
    </div>
  );
}
