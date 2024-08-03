drop table if exists role_menu_permissions;
drop table if exists menu_permissions;
drop table if exists menus;
drop table if exists permissions;
drop table if exists roles;
drop table if exists users;

create table roles (
    id serial4 not null primary key,
    name varchar(20),
    created_at timestamp DEFAULT now() null,
    updated_at timestamp null,
    deleted_at timestamp null
);

CREATE TABLE menus (
	id serial4 NOT NULL primary key,
	name varchar(20) NULL,
	slug varchar(50) NOT NULL,
	created_at timestamp DEFAULT now() NULL,
	updated_at timestamp NULL,
	deleted_at timestamp NULL,
	CONSTRAINT menus_unique UNIQUE (slug)
);

create table permissions (
    id serial4 not null primary key,
    name varchar(20),
    created_at timestamp DEFAULT now() null,
    updated_at timestamp null,
    deleted_at timestamp null
);

CREATE TABLE users (
	id serial4 NOT NULL,
	"name" varchar(50) NULL,
	email varchar(50) NULL,
	"password" varchar(255) NULL,
	last_authenticated_at timestamp NULL,
	role_id int2 NULL,
	created_at timestamp DEFAULT now() NULL,
	updated_at timestamp NULL,
	deleted_at timestamp NULL,
	is_logged_in bool NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id),
	CONSTRAINT users_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.roles(id)
);

CREATE TABLE menu_permissions (
	id serial4 NOT NULL,
	menu_id int2 NULL,
	permission_id int2 NULL,
	created_at timestamp DEFAULT now() NULL,
	updated_at timestamp NULL,
	deleted_at timestamp NULL,
	CONSTRAINT role_permissions_pkey PRIMARY KEY (id),
	CONSTRAINT role_permissions_menus_fk FOREIGN KEY (menu_id) REFERENCES public.menus(id),
	CONSTRAINT role_permissions_permission_id_fkey FOREIGN KEY (permission_id) REFERENCES public.permissions(id)
);

CREATE TABLE role_menu_permissions (
	id serial4 NOT NULL,
	role_id int2 NULL,
	menu_permission_id int2 NULL,
	created_at timestamp DEFAULT now() NULL,
	updated_at timestamp NULL,
	deleted_at timestamp NULL,
	CONSTRAINT role_menu_permissions_pk PRIMARY KEY (id),
	CONSTRAINT role_menu_permissions_menu_permissions_fk FOREIGN KEY (menu_permission_id) REFERENCES public.menu_permissions(id),
	CONSTRAINT role_menu_permissions_roles_fk FOREIGN KEY (role_id) REFERENCES public.roles(id)
);
