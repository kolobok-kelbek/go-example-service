-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS "public";

CREATE TABLE IF NOT EXISTS "public"."logins"
(
    "id"         bigserial   NOT NULL,
    "login"      text        NOT NULL UNIQUE,
    "password"   text        NOT NULL,
    "created_at" timestamptz NOT NULL default now(),
    "updated_at" timestamptz NOT NULL default now(),
    "deleted_at" timestamptz NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "public"."users"
(
    "id"         bigserial   NOT NULL,
    "login_id"   bigint      NOT NULL,
    "username"   text        NOT NULL UNIQUE,
    "email"      text        NOT NULL UNIQUE,
    "created_at" timestamptz NOT NULL default now(),
    "updated_at" timestamptz NOT NULL default now(),
    "deleted_at" timestamptz NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT fk_login
        FOREIGN KEY (login_id)
            REFERENCES "public"."logins" (id)
);

CREATE TYPE typeTime AS ENUM ('work', 'free');

CREATE TABLE IF NOT EXISTS "public"."tomatoes"
(
    "id"         bigserial   NOT NULL,
    "user_id"    bigint      NOT NULL,
    "typeTime"   typeTime    NOT NULL,
    "from"       timestamptz NOT NULL,
    "to"         timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL default now(),
    "updated_at" timestamptz NOT NULL default now(),
    "deleted_at" timestamptz NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
            REFERENCES "public"."users" (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "public"."logins";
DROP TABLE IF EXISTS "public"."periods";
DROP TABLE IF EXISTS "public"."tomatoes";
DROP TABLE IF EXISTS "public"."users";
DROP TYPE IF EXISTS typeTime;
-- +goose StatementEnd
