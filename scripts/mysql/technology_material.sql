CREATE TABLE technology_material
(
    id            bigint auto_increment,
    technology_id bigint not null comment '科技ID',
    material_id   bigint not null comment '材料ID',
    cnt         int    not null comment '数量',
    PRIMARY KEY (id)
) ENGINE = InnoDB
  COLLATE utf8mb4_general_ci COMMENT '科技材料表';
