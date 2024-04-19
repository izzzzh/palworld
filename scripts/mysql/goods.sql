CREATE TABLE goods
(
    id          bigint auto_increment primary key,
    name        varchar(255) default ''                not null comment '名称',
    description varchar(255)                           null comment '描述',
    image       varchar(255)                           null comment '图片',
    quality     int          default 0                 null comment '品质',
    types       varchar(10)                            null comment '类型',
    created_at  datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    updated_at  datetime     default CURRENT_TIMESTAMP not null comment '更新时间',
    deleted_at  datetime                               null comment '删除时间'
) ENGINE = InnoDB
  COLLATE utf8mb4_general_ci COMMENT '物品';
