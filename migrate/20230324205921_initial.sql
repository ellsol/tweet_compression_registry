-- +goose Up
-- SQL in this section is executed when the migration is applied.

create schema tcr;

create table tcr.tweet
(
    id         uuid primary key         not null default gen_random_uuid(),
    original   varchar(255) unique      not null,
    checksum   varchar(255) unique      not null,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table tcr.tweet;
drop schema tcr;
