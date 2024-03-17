CREATE TABLE pal_skill_map
(
    id       bigint auto_increment,
    pal_id   bigint not null comment '帕鲁ID',
    skill_id bigint not null comment '技能ID',
    PRIMARY KEY (id)
) ENGINE = InnoDB
  COLLATE utf8mb4_general_ci COMMENT '帕鲁技能表';
