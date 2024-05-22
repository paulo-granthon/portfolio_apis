import TurndownService from "turndown";

export default function generateMarkdown() {
  const rootElement = document.getElementById("root");
  if (rootElement) {
    const htmlString = rootElement.innerHTML;

    // Initialize Turndown service
    const turndownService = new TurndownService();

    // Convert HTML to Markdown
    const markdown = turndownService.turndown(htmlString);

    return markdown;
  }
  return "";
}
