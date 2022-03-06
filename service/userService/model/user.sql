CREATE TABLE `user`
(
    `id`         bigint NOT NULL AUTO_INCREMENT COMMENT '唯一主键id',
    `deleted`    tinyint      DEFAULT '1' COMMENT '标识是否删除1存在 2删除',
    `account`    varchar(255) DEFAULT NULL COMMENT '账号',
    `password`   varchar(255) DEFAULT NULL COMMENT '密码',
    `gender`     tinyint      DEFAULT NULL COMMENT '性别（男|女|未知）',
    `createTime` datetime     DEFAULT NULL COMMENT '创建时间',
    `updateTime` datetime     DEFAULT NULL COMMENT '最后更新时间',
    `createBy`   varchar(255) DEFAULT NULL COMMENT '创建人',
    `updateBy`   varchar(255) DEFAULT NULL COMMENT '最后更新人',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;