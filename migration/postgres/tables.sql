create table users(
    id uuid primary key ,
    first_name varchar(30),
    last_name varchar(30),
    email varchar(30),
    phone varchar(20)
);

create table orders(
    id uuid primary key ,
    amount int ,
    user_id uuid references users(id) on delete cascade,
    created_at timestamp default current_timestamp
);

create table products(
    id uuid primary key ,
    price int,
    name varchar(30)
);

create table order_products(
    id uuid primary key ,
    order_id uuid references orders(id) on delete cascade,
    product_id uuid references products(id) on delete cascade,
    quantity int,
    price int
);