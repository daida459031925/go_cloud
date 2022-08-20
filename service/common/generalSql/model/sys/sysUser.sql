CREATE TABLE `sys_user`
(
    `id`             bigint UNSIGNED  NOT NULL COMMENT '唯一主键id',
    `deleted`        tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '标识是否删除1存在 2删除',
    `account`        varchar(255)     NOT NULL COMMENT '账号',
    `password`       varchar(255)     NOT NULL COMMENT '密码',
    `salt`           varchar(255)     NOT NULL COMMENT '盐值，用作生成密码加密',
    `dict_id`        bigint UNSIGNED  NOT NULL COMMENT '性别（男|女|未知）',
    `create_user_id` bigint UNSIGNED  NOT NULL COMMENT '创建人',
    `create_time`    datetime(0)      NOT NULL COMMENT '创建时间',
    `update_user_id` bigint UNSIGNED  NULL COMMENT '最后更新人',
    `update_time`    datetime(0)      NULL COMMENT '最后更新时间',
    `secret`         varchar(255)     NOT NULL COMMENT '用户随机生成secret，用作token生成',
    `prev_secret`    varchar(255)     NOT NULL COMMENT '用户私有secret',
    `token_expire`   int UNSIGNED     NOT NULL COMMENT 'token时间 单位秒',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `IDX_sys_user_secret` (`secret`) USING BTREE,
    UNIQUE INDEX `IDX_sys_user_prev_secret` (`prev_secret`) USING BTREE,
    UNIQUE INDEX `IDX_sys_user_account` (`account`) USING BTREE
) CHARACTER SET = utf8mb4 COMMENT = '用户基础表尽量不要删除用户';