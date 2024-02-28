import Project from "./components/project";
import { ProjectSchema } from "./schemas/project";
import * as stylex from "@stylexjs/stylex";

const styles = stylex.create({
  base: {
    backgroundColor: "gray",
    fontSize: 40,
  },
});

export default function App() {
  const project: ProjectSchema = {
    id: 1,
    name: "Project 1",
    semester: 1,
    summary: "This is a summary",
    company: "Company 1",
    url: "https://example.com",
  };
  return (
    <div {...stylex.props(styles.base)}>
      <h1>Porfolio APIs</h1>
      <Project {...project} />
    </div>
  );
}
