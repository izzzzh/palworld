CREATE TABLE technology_tree
(
    id          bigint auto_increment,
    name        varchar(255) not null default '' comment '名称',
    description varchar(255) comment '描述',
    cost        int          not null comment '消耗',
    icon        varchar(255) not null default '' comment '图标',
    ancient     bool                  default false comment '是否古代科技',
    level       int          not null comment '解锁等级',
    PRIMARY KEY (id)
) ENGINE = InnoDB
  COLLATE utf8mb4_general_ci COMMENT '科技树';
