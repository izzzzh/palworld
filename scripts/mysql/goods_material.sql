CREATE TABLE goods_material
(
    id          bigint auto_increment,
    goods_id    bigint not null comment '物品ID',
    material_id bigint not null comment '材料ID',
    count       int    not null comment '数量',
    PRIMARY KEY (id)
) ENGINE = InnoDB
  COLLATE utf8mb4_general_ci COMMENT '物品材料表';
