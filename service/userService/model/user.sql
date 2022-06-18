CREATE TABLE `sys_user`  (
     `id` bigint NOT NULL AUTO_INCREMENT COMMENT '唯一主键id',
     `deleted` tinyint NULL DEFAULT 1 COMMENT '标识是否删除1存在 2删除',
     `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '账号',
     `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '密码',
     `salt` varchar(255) NULL COMMENT '盐值，用作生成密码加密',
     `gender` tinyint NULL DEFAULT NULL COMMENT '性别（男|女|未知）',
     `createTime` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
     `updateTime` datetime(0) NULL DEFAULT NULL COMMENT '最后更新时间',
     `createBy` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '创建人',
     `updateBy` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '最后更新人',
     `secret` varchar(255) NULL COMMENT '用户随机生成secret，用作token生成',
     `prevSecret` varchar(255) NULL COMMENT '用户私有secret',
     `tokenExpire` int NULL COMMENT 'token时间 单位秒',
     PRIMARY KEY (`id`)
);