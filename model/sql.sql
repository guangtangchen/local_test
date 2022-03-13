create table `test_table`
(
    `id`           bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `order_no`     varchar(64)     NOT NULL comment '字段1',
    `name`         varchar(64)     NOT NULL comment '字段1',
    `content`      text comment 'text',
    `num`          int(10)         not null default 0 comment 'int',
    `gmt_created`  datetime(3)     not null default current_timestamp(3) comment '创建时间',
    `gmt_modified` datetime(3)     not null default current_timestamp(3) on update current_timestamp(3) comment '更新时间',
    primary key (`id`),
    unique key `uniq_order_no` (`order_no`),
    KEY `idx_name` (`name`)
) engine = InnoDB
  auto_increment = 9000
  default charset = utf8mb4 comment ='测试表';


show tables;



