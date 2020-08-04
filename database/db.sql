CREATE  OR REPLACE FUNCTION next_id(OUT result bigint, seq text) AS $$
DECLARE
    epoch_time bigint := 1594958459980;
    seq_id bigint;
    time_now bigint;
    shard_id int := 5;
BEGIN
    SELECT nextval(seq) % 1024 INTO seq_id;
    SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp())) INTO time_now;
    result := (time_now - epoch_time)*1000 << 23;
    result := result | (shard_id << 10);
    result := result | (seq_id);
END;
    $$ LANGUAGE PLPGSQL;

-- Users table
DROP TABLE IF EXISTS users;
DROP SEQUENCE IF EXISTS users_id_seq;
CREATE SEQUENCE users_id_seq;
CREATE TABLE users(
    id bigint not null default next_id('users_id_seq') primary key,
    name text,
    username text unique,
    password text,
    email text,
    created_at timestamptz default now(),
    updated_at timestamptz default now(),
    deleted_at timestamptz
);

-- Tasks table
DROP TABLE IF EXISTS tasks;
DROP SEQUENCE IF EXISTS tasks_id_seq;
CREATE SEQUENCE tasks_id_seq;
CREATE TABLE tasks(
    id bigint not null default next_id('tasks_id_seq') primary key,
    name text,
    status smallint,
    task_type smallint,
    user_id bigint,
    created_at timestamptz default now(),
    updated_at timestamptz,
    foreign key (user_id) references users(id)
);