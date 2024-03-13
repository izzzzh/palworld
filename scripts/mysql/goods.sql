CREATE TABLE goods
(
    id          bigint auto_increment,
    name        varchar(255) not null default '' comment '名称',
    en_name     varchar(255) not null default '' comment '英文名称',
    description varchar(255) comment '描述',
    image       varchar(255) comment '图片',
    materials   json comment '制造材料',
    quality     int comment '品质',
    workload    int comment '工作量',
    types       int          not null default 0 comment '类型',
    PRIMARY KEY (id)
) ENGINE = InnoDB
  COLLATE utf8mb4_general_ci COMMENT '物品';
