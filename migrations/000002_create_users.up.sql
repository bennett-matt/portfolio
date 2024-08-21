create table if not exists users(
  id bigserial primary key,
  name text not null,
  email text not null,
  hashed_password text not null,
  created_at timestamp(0) with time zone not null default now(),
  updated_at timestamp(0) with time zone not null default now(),
  constraint users_uc_email unique (email)
);