CREATE TABLE `sys_user`
(
    `id`             bigint NOT NULL AUTO_INCREMENT COMMENT '唯一主键id',
    `deleted`        tinyint                                                       DEFAULT '1' COMMENT '标识是否删除1存在 2删除',
    `type`           tinyint                                                       DEFAULT NULL COMMENT '类型，用于区分用户加入来源：1.手动添加，2批量添加，3微信注册，4pc注册 5，同步',
    `account`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '账号',
    `password`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '密码',
    `salt`           varchar(255)                                                  DEFAULT NULL COMMENT '盐值，用作生成密码加密',
    `gender`         tinyint                                                       DEFAULT '0' COMMENT '性别（未知|男|女）',
    `createTime`     datetime                                                      DEFAULT NULL COMMENT '创建时间',
    `updateTime`     datetime                                                      DEFAULT NULL COMMENT '最后更新时间',
    `createBy`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '创建人',
    `updateBy`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '最后更新人',
    `createByUserId` bigint                                                        DEFAULT NULL COMMENT '为了查询用户，保存id，空间换时间的备用',
    `updateByUserId` bigint                                                        DEFAULT NULL COMMENT '为了查询用户，保存id，空间换时间的备用',
    `secret`         varchar(255)                                                  DEFAULT NULL COMMENT '用户随机生成secret，用作token生成',
    `prevSecret`     varchar(255)                                                  DEFAULT NULL COMMENT '用户私有secret',
    `tokenExpire`    int                                                           DEFAULT NULL COMMENT 'token时间 单位秒',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


