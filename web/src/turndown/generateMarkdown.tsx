import TurndownService from "turndown";

const imageRegex = /!\[(.*)\]\((.*)\)/g;

export default function generateMarkdown() {
  const rootElement = document.getElementById("root");
  if (rootElement) {
    const htmlString = rootElement.innerHTML;

    // Initialize Turndown service
    const turndownService = new TurndownService();

    // Convert HTML to Markdown
    const markdown = turndownService.turndown(htmlString);

    return processMarkdown(markdown);
  }
  return "";
}

function processMarkdown(markdown: string) {
  markdown = processImages(markdown);

  return markdown;
}

// use html `<img>` tag instead of markdown `![alt](src)`
// parse the `alt` attribute to extract the image height, it will be in the format `x*px`, example: `128px`
// the `alt` will be the `alt` without the image height
// set the height of the image to the extracted height
function processImages(markdown: string): string {
  const imageMarkdown = markdown.match(imageRegex);

  if (!imageMarkdown) return markdown;

  imageMarkdown.forEach((image) => {
    const imageAlt = image.match(/!\[(.*)\]/);
    const imageSrc = image.match(/\((.*)\)/);
    if (imageAlt && imageSrc) {
      let imageAltText = imageAlt[1];

      const imageHeight = imageAltText.match(/(\d+)px/);
      imageAltText = imageAltText.replace(/\s*(\d+)px/, "");

      const imageSrcText = imageSrc[1];

      let imgTag = `<img alt="${imageAltText}" src="${imageSrcText}" `;
      if (imageHeight) {
        imgTag += `height="${imageHeight[1]}" `;
      }
      imgTag += `/>`;
      markdown = markdown.replace(image, imgTag);
    }
  });

  return markdown;
}
