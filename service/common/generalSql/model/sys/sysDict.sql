CREATE TABLE `sys_dict`
(
    `id`             bigint UNSIGNED NOT NULL COMMENT '字典主键',
    `deleted`        tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '标识是否删除1存在 2删除',
    `status`         tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态（1正常 2停用）',
    `is_edit`        tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '表示当前内容是否可以供普通用户修改（1可以2不行）',
    `type`           varchar(255) NOT NULL COMMENT '字典类型',
    `name`           varchar(255) NOT NULL COMMENT '字典名称',
    `super_id`       bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级id可以构建树形',
    `describe`       varchar(255) NULL COMMENT '备注',
    `css_class`      varchar(255) NULL COMMENT '样式属性（其他样式扩展）',
    `list_class`     varchar(255) NULL COMMENT '表格回显样式',
    `create_user_id` bigint UNSIGNED NOT NULL COMMENT '创建者',
    `create_time`    datetime     NOT NULL COMMENT '创建时间',
    `update_user_id` bigint UNSIGNED NULL COMMENT '更新者',
    `update_time`    datetime NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX            `IDX_sys_dict_type`(`type`) USING BTREE,
    INDEX            `IDX_sys_dict_super_id`(`super_id`) USING BTREE COMMENT '加快查询上下级关系'
) CHARACTER SET = utf8mb4 COMMENT = '字典类型表,用作系统中常见可供用户选择修改的字段';
