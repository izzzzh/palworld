CREATE TABLE pal_mate_map
(
    id         bigint auto_increment,
    parent_one bigint not null comment '父母一号',
    parent_two bigint not null comment '父母二号',
    result     bigint not null comment '后代',
    PRIMARY KEY (id)
) ENGINE = InnoDB
  COLLATE utf8mb4_general_ci COMMENT '帕鲁配对表';
