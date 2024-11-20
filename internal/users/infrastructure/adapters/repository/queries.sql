-- name: Save :one
insert into users (public_id, type, user_name, email, password, balance) values ($1, $2, $3, $4, $5, $6) returning id, created_at, updated_at;

-- name: GetByEmail :one
select id, public_id, type, user_name, email, balance, password, created_at, updated_at from users where email = $1;
