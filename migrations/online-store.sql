drop table if exists users;
drop table if exists roles;
drop table if exists permissions;
drop table if exists role_permissions;

create table roles (
    id serial4 not null primary key,
    name varchar(20),
    created_at timestamp DEFAULT now() null,
    updated_at timestamp null,
    deleted_at timestamp null
);

create table permissions (
    id serial4 not null primary key,
    name varchar(20),
    created_at timestamp DEFAULT now() null,
    updated_at timestamp null,
    deleted_at timestamp null
);

create table role_permissions (
    id serial4 not null primary key,
    role_id int2,
    permission_id int2,
    created_at timestamp DEFAULT now() null,
    updated_at timestamp null,
    deleted_at timestamp null,

    foreign key (role_id) references roles (id),
    foreign key (permission_id) references permissions (id)
);

create table users (
    id serial4 not null primary key,
    name varchar(50),
    email varchar(50),
    password varchar(255),
    last_authenticated_at timestamp null,
    role_id int2,
    created_at timestamp DEFAULT now() null,
    updated_at timestamp null,
    deleted_at timestamp null,
    foreign key (role_id) references roles (id)
);
