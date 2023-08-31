CREATE TABLE segments
(
    id serial primary key,
    slug varchar(255) not null unique
);

CREATE TABLE users_segments
(
    id serial primary key,
    user_id serial not null,
    segment_id serial not null,
    foreign key (segment_id) references segments (id) on delete cascade,
    unique (user_id, segment_id)
);

CREATE TYPE segment_operation AS ENUM('add', 'delete');

CREATE TABLE users_segments_history
(
    id serial primary key,
    user_id serial not null,
    segment_slug varchar(255) not null,
    operation segment_operation not null,
    updated_at timestamp not null
);

CREATE UNIQUE INDEX users_segments_index
    ON users_segments (user_id, segment_id)