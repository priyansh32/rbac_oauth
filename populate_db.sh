#!/bin/bash

# SQLite database file
DB_FILE="foo.db"

# Check if the database file exists
if [ -e "$DB_FILE" ]; then
    echo "Database file already exists. Skipping population."
else
    # Execute SQL script using sqlite3
    sqlite3 "$DB_FILE" < populate.sql

    echo "Database populated successfully."
fi
