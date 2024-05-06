create table if not exists category (
    id serial primary key,
    categoryName varchar(255) not null
);

create table if not exists product (
    id serial primary key,
    productName varchar(255) not null,
    productCategory int,
    productPrice float not null,
    foreign key (productCategory) references category (id) on delete cascade
);

