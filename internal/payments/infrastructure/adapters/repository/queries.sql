-- name: CreateTransaction :one
insert into transactions (
	public_id,
	status,
	method,
	description,
	receiver_id,
	sender_id,
	amount
) values ( $1, $2, $3, $4, $5, $6, $7 ) returning id, created_at, updated_at;

-- name: GetTransaction :one
select
    t.id,
    t.public_id,
    t.status,
    t.method,
    t.description,
    t.receiver_id,
    t.sender_id,
    t.amount,
    t.created_at,
    t.updated_at,
    u.public_id as receiver_public_id,
    ua.public_id as sender_public_id
from
    transactions t
inner join
    users u on t.receiver_id = u.id
inner join
    users ua on t.sender_id = ua.id
where
    t.public_id = $1;

-- name: GetReceiverTransactions :many
select
    t.id,
    t.public_id,
    t.status,
    t.method,
    t.description,
    t.receiver_id,
    t.sender_id,
    t.amount,
    t.created_at,
    t.updated_at,
    u.public_id as receiver_public_id,
    ua.public_id as sender_public_id
from
    transactions t
inner join
    users u on t.receiver_id = u.id
inner join
    users ua on t.sender_id = ua.id
where
    u.public_id = $1;

-- name: GetSenderTransactions :many
select
    t.id,
    t.public_id,
    t.status,
    t.method,
    t.description,
    t.receiver_id,
    t.sender_id,
    t.amount,
    t.created_at,
    t.updated_at,
    u.public_id as receiver_public_id,
    ua.public_id as sender_public_id
from
    transactions t
inner join
    users u on t.receiver_id = u.id
inner join
    users ua on t.sender_id = ua.id
where
    u.public_id = $1;

-- name: UpdateTransaction :one
update transactions
  set status = $1
  where public_id = $2
returning
  id,
  status,
	public_id,
	status,
	method,
	description,
	receiver_id,
	sender_id,
	amount,
  created_at,
  updated_at;
