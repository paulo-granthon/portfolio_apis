import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import User from "../components/user";
import ProjectList from "../components/projectList";
import { UserSchema } from "../schemas/user";
import { getUser } from "../services/user";
import * as styles from "../styles/portfolio";

export default function Portfolio() {
  const [user, setUser] = useState<UserSchema | undefined>();

  const params = useParams<{ userId: string }>();

  useEffect(() => {
    const id = params.userId ? parseInt(params.userId, 10) : undefined;
    if (!id) return;

    getUser(id).then((user) => setUser(user));
  }, [params.userId]);

  return (
    <>
      {user ? (
        <div>
          <div {...styles.portfolio}>
            <h1>Portfolio</h1>
          </div>
          <User user={user} />
          <ProjectList user={user} />
        </div>
      ) : (
        <p>Usuário não encontrado</p>
      )}
    </>
  );
}
