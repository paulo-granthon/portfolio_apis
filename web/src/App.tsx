import Portfolio from "./pages/portfolio";
import { UserSchema } from "./schemas/user";
import { YearSemester } from "./schemas/user";

export default function App() {
  const user: UserSchema = {
    id: 0,
    name: "paulo",
    summary:
      "Backend developer intern at @gorilainvest | Database technologist student at FATEC | Self titled full-stack developer",
    semesterMatriculated: new YearSemester(2022, 2),
    githubUsername: "paulo-granthon",
  };

  return <Portfolio {...user} />;
}
