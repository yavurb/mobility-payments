-- name: Save :one
insert into users (public_id, type, user_name, email, password, balance) values ($1, $2, $3, $4, $5, $6) returning id, created_at, updated_at;

-- name: GetByPublicID :one
select id, public_id, type, user_name, email, balance, password, created_at, updated_at from users where public_id = $1;

-- name: GetByEmail :one
select id, public_id, type, user_name, email, balance, password, created_at, updated_at from users where email = $1;

-- name: UpdateBalance :one
update users set balance = $1 where public_id = $2 returning balance;
