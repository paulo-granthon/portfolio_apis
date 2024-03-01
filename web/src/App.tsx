import { useEffect, useState } from "react";
import { UserSchema } from "./schemas/user";
import { getUser } from "./services/user";
import Portfolio from "./pages/portfolio";

export default function App() {
  const [user, setUser] = useState<UserSchema | undefined>();

  useEffect(() => {
    getUser(1).then((user) => setUser(user));
  }, []);

  return <>{user ? <Portfolio {...user} /> : <p>Usuário não encontrado</p>}</>;
}
