CREATE DATABASE mconfig_admin;

create table m_app
(
    id          int auto_increment
        primary key,
    app_name    varchar(255) not null,
    app_key     varchar(255) not null,
    description varchar(255) null,
    create_user int          null,
    create_time bigint       not null,
    update_time bigint       not null,
    update_user int          null,
    constraint m_app_app_key_uindex
        unique (app_key)
);

create table m_cluster
(
    id          int auto_increment
        primary key,
    namespace   varchar(255) not null,
    register    varchar(255) not null,
    description varchar(255) null,
    create_user int          null,
    create_time bigint       not null,
    update_time bigint       not null,
    update_user int          null
);

create table m_config
(
    id            int auto_increment
        primary key,
    app_id        int          not null,
    env_id        int          not null,
    config_name   varchar(255) not null,
    config_key    varchar(255) not null,
    config_value  varchar(255) null,
    config_schema varchar(255) null,
    description   varchar(255) null,
    create_user   int          null,
    create_time   bigint       not null,
    update_time   bigint       not null,
    update_user   int          null,
    constraint m_config_key_uindex
        unique (config_key)
);

create table m_env
(
    id          int auto_increment
        primary key,
    app_id      int          not null,
    filter      int          not null,
    env_name    varchar(255) not null,
    env_key     varchar(255) not null,
    description varchar(255) null,
    create_user int          null,
    create_time bigint       not null,
    update_time bigint       not null,
    update_user int          null,
    constraint m_app_app_key_uindex
        unique (env_key)
);

create table m_filter
(
    id          int auto_increment
        primary key,
    type        int    null comment '0 -- lua
1 -- js',
    filter      text   null,
    create_time bigint null,
    update_time bigint null,
    create_user int    null,
    update_user int    null
);

create table m_log
(
    id          int auto_increment
        primary key,
    module      int          null,
    action      varchar(255) null,
    user        int          null,
    create_time bigint       not null
)
    charset = latin1;



