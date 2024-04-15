CREATE TABLE comments
(
    id                bigint auto_increment primary key,
    user_id           bigint                       not null comment '用户id',
    category          enum ('pal','goods','skill') not null comment '类型',
    object_id         bigint                       not null comment '对象id',
    content           TEXT                         not null comment '内容',
    parent_comment_id bigint                       null comment '父评论id',
    root_comment_id   bigint                       null comment '根评论id',
    created_at        datetime default now() comment '创建时间',
    updated_at        datetime default now() comment '更新时间',
    deleted_at        datetime comment '删除时间'
) ENGINE = InnoDB
  COLLATE utf8mb4_general_ci COMMENT '评论';
