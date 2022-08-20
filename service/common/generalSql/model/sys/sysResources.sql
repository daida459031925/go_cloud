CREATE TABLE `sys_resources`
(
    `id`             bigint UNSIGNED NOT NULL COMMENT '唯一主键id',
    `deleted`        tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '标识是否删除1存在 2删除',
    `name`           varchar(255) NOT NULL COMMENT '资源名称',
    `type`           varchar(255) NOT NULL COMMENT '资源的类型',
    `url`            varchar(255) NULL COMMENT '资源路径',
    `model`          varchar(255) NOT NULL COMMENT '资源请求类型',
    `super_id`       bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级id可以构建树形',
    `describe`       varchar(500) NULL DEFAULT NULL COMMENT '请求方法内容描述',
    `order_by`       int UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
    `create_user_id` bigint UNSIGNED NOT NULL COMMENT '创建者',
    `create_time`    datetime(0) NOT NULL COMMENT '创建时间',
    `update_user_id` bigint UNSIGNED NULL COMMENT '更新者',
    `update_time`    datetime(0) NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX            `IDX_sys_res_super_id`(`super_id`) USING BTREE COMMENT '加快查询上下级关系'
) CHARACTER SET = utf8mb4 COMMENT = '系统基本资源';