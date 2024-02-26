#!/bin/bash
PORT=3333

if ! lsof -i :"$PORT"; then exit 0; fi

echo "Stopping api on port $PORT"
lsof -i :"$PORT" | awk 'NR!=1 {print $2}' | xargs kill
