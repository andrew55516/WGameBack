CREATE TABLE "users"
(
    "user_id"       bigint PRIMARY KEY,
    "user_name" varchar NOT NULL,
    "password"  decimal NOT NULL,
    "email" varchar NOT NULL UNIQUE,
    "user_type" varchar NOT NULL,
    "token" varchar NOT NULL,
    "refresh_token" varchar NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL,
);