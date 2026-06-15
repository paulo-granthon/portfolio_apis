#!/bin/bash
# Serves the presentation slides (docs/slides.html) locally and opens them in the browser

PORT=${1:-8080}
SLIDES_DIR="$(dirname "$0")/docs"
URL="http://localhost:${PORT}/slides.html"

cd "$SLIDES_DIR" || exit 1

echo "Serving slides at ${URL} (Ctrl+C to stop)"
xdg-open "$URL" >/dev/null 2>&1 &

python3 -m http.server "$PORT"
