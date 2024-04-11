CREATE TABLE user
(
    id          bigint auto_increment,
    name        varchar(100) not null comment '姓名',
    password    varchar(225) not null comment '密码',
    role        int          not null default 1 comment '角色',
    avatar      varchar(255) not null comment '头像',
    create_at datetime     not null default current_timestamp comment '创建时间',
    PRIMARY KEY (id),
    UNIQUE KEY (name)
) ENGINE = InnoDB
  COLLATE utf8mb4_general_ci COMMENT '用户';
