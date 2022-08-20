CREATE TABLE `sys_role`
(
    `id`             bigint UNSIGNED NOT NULL COMMENT '唯一主键id',
    `deleted`        tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '标识是否删除1存在 2删除',
    `name`           varchar(255) NOT NULL COMMENT '角色名字用作标识客户看到的名字，可根据不同业务来实现名字是否相同',
    `super_id`       bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级id可以构建树形',
    `create_user_id` bigint UNSIGNED NOT NULL COMMENT '创建者',
    `create_time`    datetime(0) NOT NULL COMMENT '创建时间',
    `update_user_id` bigint UNSIGNED NULL COMMENT '更新者',
    `update_time`    datetime(0) NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX            `IDX_sys_role_super_id`(`super_id`) USING BTREE COMMENT '加快查询上下级关系'
) CHARACTER SET = utf8mb4 COMMENT = '角色表用作确定系统中唯一角色存在';