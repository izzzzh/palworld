CREATE TABLE attribute
(
    id            bigint auto_increment,
    name          varchar(255) not null default '' comment '名称',
    alias         varchar(10)  not null default '' comment '别名',
    restrain      varchar(10) comment '克制',
    be_restrained int comment '被克制',
    PRIMARY KEY (id)
) ENGINE = InnoDB
  COLLATE utf8mb4_general_ci COMMENT '属性';
