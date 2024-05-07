create table if not exists category (
    category_id serial primary key,
    category_name varchar(255) not null
);

create table if not exists product (
    product_id serial primary key,
    product_name varchar(255) not null,
    product_category int,
    product_price float not null,
    foreign key (product_category) references category (category_id) on delete cascade
);

