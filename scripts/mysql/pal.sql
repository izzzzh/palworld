CREATE TABLE pal
(
    id            bigint auto_increment,
    number        varchar(10)  not null default '' comment '编号',
    name          varchar(10)  not null default '' comment '名字',
    en_name       varchar(100) not null default '' comment '英文名字',
    icon          varchar(255) not null default '' comment '头像',
    attribute_ids varchar(10)  not null comment '属性id',
    hp            int          not null comment '血量',
    energy        int          not null comment '攻击力',
    defensively   int          not null comment '防御力',
    abilities     json comment '工作能力',
    eat           int                   default 0 comment '消耗',
    passive_skill int comment '被动技能',
    level         int                   default 0 comment '级别（孵蛋大小）',
    UNIQUE name_index (name),
    PRIMARY KEY (id)
) ENGINE = InnoDB
  COLLATE utf8mb4_general_ci COMMENT '帕鲁';
