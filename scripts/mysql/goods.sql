CREATE TABLE goods
(
    id          bigint auto_increment primary key,
    name        varchar(255) default '' not null comment '名称',
    description varchar(255)            null comment '描述',
    image       varchar(255)            null comment '图片',
    materials   json                    null comment '制造材料',
    quality     int          default 0  null comment '品质',
    workload    int                     null comment '工作量',
    types       varchar(10)             null comment '类型'
) ENGINE = InnoDB
  COLLATE utf8mb4_general_ci COMMENT '物品';
