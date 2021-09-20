#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE USER $APP_DB_USER WITH PASSWORD '$APP_DB_PASS';
  CREATE DATABASE $APP_DB_NAME;
  GRANT ALL PRIVILEGES ON DATABASE $APP_DB_NAME TO $APP_DB_USER;
  \connect $APP_DB_NAME $APP_DB_USER
  BEGIN;
        CREATE TABLE game_logs (
        id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
        created_at timestamp with time zone DEFAULT now(),
        updated_at timestamp with time zone DEFAULT now(),
        start timestamp with time zone,
        "end" timestamp with time zone,
        nb_snake integer,
        nb_turn integer,
        won boolean DEFAULT false,
        position integer,
        ruleset text,
        width integer,
        height integer,
        algorithm_version integer DEFAULT 0
    );
    CREATE INDEX game_logs_ruleset_idx ON game_logs(ruleset text_ops);
    CREATE INDEX game_logs_start_idx ON game_logs(start timestamptz_ops);
    CREATE INDEX game_logs_won_idx ON game_logs(won bool_ops);
    CREATE INDEX game_logs_algorithm_version_idx ON game_logs(algorithm_version int4_ops);
  COMMIT;
EOSQL