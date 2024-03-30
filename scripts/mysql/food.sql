CREATE TABLE food
(
    id          bigint auto_increment,
    name        varchar(255) not null default '' comment '名称',
    description varchar(255) comment '描述',
    san         int          not null comment 'san值回复',
    nutrition   int          not null comment '营养价值',
    expire      int          not null comment '腐败时间',
    has_buff    bool                  default false comment '是否有buff加成',
    UNIQUE name_index (name),
    PRIMARY KEY (id)
) ENGINE = InnoDB
  COLLATE utf8mb4_general_ci COMMENT '消耗品';
