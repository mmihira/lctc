#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER lctc;
    CREATE DATABASE lctc;
    GRANT ALL PRIVILEGES ON DATABASE lctc TO lctc;
EOSQL
