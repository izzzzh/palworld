CREATE TABLE skill
(
    id           bigint auto_increment,
    name         varchar(255) not null default '' comment '名称',
    description  varchar(255) not null comment '描述',
    ct           int          not null comment '冷却',
    power        int          not null comment '威力',
    attribute_id int          not null comment '属性ID',
    UNIQUE name_index (name),
    PRIMARY KEY (id)
) ENGINE = InnoDB
  COLLATE utf8mb4_general_ci COMMENT '技能';
