#!/bin/bash

# function to imediately execute a query on connect
function watch_query() {
    query=${1:-"show databases"}
    watch -etn 1 "mysql -h$DB_HOST -u$DB_USER -p$DB_PASSWORD $DB_NAME -t -e '$query'"
    clear
}

# Function to execute a query once
function execute_query_once() {
    query=${1:-"show databases"}
    mysql -h"${DB_HOST}" -u"${DB_USER}" -p"${DB_PASSWORD}" "${DB_NAME}" -t -e "${query}"
}

function execute_sql_script_file() {
    script_file=${1}

    if [ ! -f "${script_file}" ]; then
        echo "File ${script_file} does not exist."
        echo -e "Usage: $0 <db_index> --script <script_file>"
        exit 1
    fi

    mysql -h"${DB_HOST}" -u"${DB_USER}" -p"${DB_PASSWORD}" "${DB_NAME}" < "${script_file}"
}

# Load the environment variables
# shellcheck source=./.env
. ./.env

case "$1" in
    "--watch" | "-w")
        query=${2:-"show databases"}
        watch_query "${query}"
        ;;
    "--query" | "-q")
        query=${2:-"show databases"}
        execute_query_once "${query}"
        ;;
    "--script" | "-s")
        execute_sql_script_file "${2}"
        ;;
    *)
        # If no specific flag is provided, execute pgcli command
        pgcli postgresql://"${DB_USER}":"${DB_PASSWORD}"@"${DB_HOST}":"${DB_PORT}"
        ;;
esac

