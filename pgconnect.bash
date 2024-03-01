#!/bin/bash

# function to imediately execute a query on connect
function watch_query() {
    query=${1:-"show databases"}
    DB_NAME_MAYBE=""
    if [ -n "${DB_NAME}" ]; then
        DB_NAME_MAYBE="-d ${DB_NAME}"
    fi
    PGPASSWORD="${DB_PASSWORD}" watch -etn 1 "psql -h${DB_HOST} -U${DB_USER} ${DB_NAME_MAYBE} -c '${query}'"
    clear
}

# Function to execute a query once
function execute_query_once() {
    query=${1:-"show databases"}
    DB_NAME_MAYBE=""
    if [ -n "${DB_NAME}" ]; then
        DB_NAME_MAYBE="-d ${DB_NAME}"
    fi
    PGPASSWORD="${DB_PASSWORD}" psql -h"${DB_HOST}" -U"${DB_USER}" "${DB_NAME_MAYBE}" -c "${query}"
}

function execute_sql_script_file() {
    script_file=${1}
    if [ ! -f "${script_file}" ]; then
        echo "File ${script_file} does not exist."
        echo -e "Usage: $0 <db_index> --script <script_file>"
        exit 1
    fi

    DB_NAME_MAYBE=""
    if [ -n "${DB_NAME}" ]; then
        DB_NAME_MAYBE="-d ${DB_NAME}"
    fi

    PGPASSWORD="${DB_PASSWORD}" psql -h"${DB_HOST}" -U"${DB_USER}" "${DB_NAME_MAYBE}" -f "${script_file}"
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
        pgcli postgresql://"${DB_USER}":"${DB_PASSWORD}"@"${DB_HOST}":"${DB_PORT}"/"${DB_NAME}"
        ;;
esac

