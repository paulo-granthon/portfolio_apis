#!/bin/bash
PORT=3333
echo "Stopping api on port $PORT"

if ! lsof -i :"$PORT"
then
  echo "It seems that api isn't up on port $PORT"
  exit 0
fi

lsof -i :"$PORT" | awk 'NR!=1 {print $2}' | xargs kill
