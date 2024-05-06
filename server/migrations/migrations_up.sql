create table if not exists category (
    id serial primary key,
    categoryName varchar(255) not null
);

create table if not exists product (
    product_id serial primary key,
    product_name varchar(255) not null,
    product_category int,
    product_price float not null,
    foreign key (productCategory) references category (id) on delete cascade
);

