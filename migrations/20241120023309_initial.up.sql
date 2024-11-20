create type user_type as enum('customer', 'merchant');

create table users (
    id bigserial primary key,
    public_id varchar(15) not null unique,
    user_name varchar(255) not null,
    email varchar(255) not null,
    type user_type not null,
    password varchar(255) not null,
    balance bigint not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

create type transaction_status as enum('requires_confirmation', 'succeeded', 'declined');
create type transaction_payment_method as enum('credit_card', 'debit_card', 'bank_transfer');

create table transactions (
  id bigserial primary key,
  public_id varchar(15) not null unique,
  sender_id bigint not null references users(id) on delete cascade,
  receiver_id bigint not null references users(id) on delete cascade,
  status transaction_status not null,
  currency varchar(3) not null,
  amount bigint not null,
  method transaction_payment_method not null,
  description varchar(255) not null default '',
  created_at timestamp not null default now(),
  updated_at timestamp not null default now()
)
