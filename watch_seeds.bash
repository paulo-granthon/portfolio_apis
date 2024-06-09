#!/bin/bash
# Performs automatic database migration when the seeds file is modified

SEEDS_FILE="./api/seeds/seeds.go"
COMMAND="make dm"

GREEN='\033[0;32m'
NC='\033[0m'

while true; do
	inotifywait -e modify $SEEDS_FILE
	echo -e "${GREEN}Seeds file has been modified. Running \`$COMMAND\`${NC}"
	$COMMAND
	echo -e "${GREEN}Done${NC}"
done
