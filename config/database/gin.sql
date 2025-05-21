/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : 127.0.0.1:3306
 Source Schema         : gin

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 21/05/2025 14:25:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '用户id',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '内容',
  `category_id` int(11) NOT NULL DEFAULT 0 COMMENT '分类id',
  `data_source` tinyint(3) UNSIGNED NOT NULL DEFAULT 2 COMMENT '数据来源 1=文章库 2=自建',
  `is_publish` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否发布 1=已发布 2=未发布',
  `tag` json NULL COMMENT '标签',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '文章表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES (1, 1, '标题1', '内容1', 1, 2, 1, '[\"测试标签1\", \"测试标签2\"]', '2023-09-19 11:43:58', '2024-09-18 16:35:25', NULL);
INSERT INTO `article` VALUES (13, 1, '测试77', '测试内容77', 0, 2, 1, '[\"标签3\", \"标签4\"]', '2024-07-22 11:21:18', '2024-07-22 11:21:18', NULL);

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '分类名称',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '分类表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (1, '分类名称', '2023-09-19 11:43:43', '2023-09-19 11:43:43', NULL);

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `pid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级id',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '路由名称',
  `path` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '路由路径',
  `redirect` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '重定向',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
  `component_alias` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '组件路径',
  `is_link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '链接地址',
  `is_hide` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否隐藏 1=显示 2=隐藏',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态 1=启用 2=停用',
  `sort` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_pid`(`pid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES (1, 0, '首页', 'home', '/home', '', 'iconfont icon-shouye', 'home/index', '', 1, 1, 0, '2025-05-14 17:04:02', '2025-05-14 17:04:05', NULL);
INSERT INTO `menu` VALUES (2, 0, '系统设置', 'system', '/system', '/system/menu', 'iconfont icon-xitongshezhi', 'layout/routerView/parent', '', 1, 1, 0, '2025-05-21 10:35:36', '2025-05-21 10:35:38', NULL);
INSERT INTO `menu` VALUES (3, 2, '菜单管理', 'systemMenu', '/system/menu', '', 'iconfont icon-caidan', 'system/menu/index', '', 1, 1, 0, '2025-05-21 10:38:17', '2025-05-21 10:38:17', NULL);

-- ----------------------------
-- Table structure for menu_action
-- ----------------------------
DROP TABLE IF EXISTS `menu_action`;
CREATE TABLE `menu_action`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `menu_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '菜单id',
  `type` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '类型 1=header 2=operation',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '功能名称',
  `is_link` tinyint(3) UNSIGNED NOT NULL DEFAULT 2 COMMENT '是否为链接 1=是 2=否',
  `sort` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_menu_id`(`menu_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '菜单功能表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menu_action
-- ----------------------------
INSERT INTO `menu_action` VALUES (1, 3, 1, '新增菜单', 2, 0, '2025-05-21 10:24:14', '2025-05-21 10:24:14', NULL);
INSERT INTO `menu_action` VALUES (2, 3, 2, '编辑', 2, 0, '2025-05-21 10:30:24', '2025-05-21 10:30:24', NULL);
INSERT INTO `menu_action` VALUES (3, 3, 2, '功能', 2, 0, '2025-05-21 10:30:37', '2025-05-21 10:30:37', NULL);
INSERT INTO `menu_action` VALUES (4, 3, 2, '删除', 2, 0, '2025-05-21 10:30:49', '2025-05-21 10:30:49', NULL);

-- ----------------------------
-- Table structure for system_config
-- ----------------------------
DROP TABLE IF EXISTS `system_config`;
CREATE TABLE `system_config`  (
  `id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `cn_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '中文名称',
  `en_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '英文名称',
  `default_value` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '默认值',
  `option_value` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '可选值',
  `type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '配置类型 1=输入框 2=单选 3=复选 4=下拉菜单 5=文本域 6=附件',
  `category` tinyint(1) NOT NULL DEFAULT 1 COMMENT '配置分类 1=基本信息 2=联系方式 3=seo设置 ',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_en_name`(`en_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '系统配置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of system_config
-- ----------------------------
INSERT INTO `system_config` VALUES (1, '网站域名', 'web_domain', 'www.a.com', '', 1, 1, NULL, NULL, NULL);
INSERT INTO `system_config` VALUES (2, '关闭站点', 'is_open_site', '开启', '关闭,开启', 2, 1, NULL, NULL, NULL);
INSERT INTO `system_config` VALUES (3, '网站Logo', 'site_logo', '', '', 6, 1, NULL, NULL, NULL);
INSERT INTO `system_config` VALUES (4, '邮件端口', 'email_port', '465', '', 1, 2, NULL, NULL, NULL);
INSERT INTO `system_config` VALUES (5, '邮件标题', 'email_title', '【xxx】验证码', '', 1, 2, NULL, NULL, NULL);
INSERT INTO `system_config` VALUES (6, '发件人信息', 'send_user_info', '【管理员】', '', 1, 2, NULL, NULL, NULL);
INSERT INTO `system_config` VALUES (7, '发送内容', 'email_content', '【xxx】你的验证码是：', '', 5, 2, NULL, NULL, NULL);
INSERT INTO `system_config` VALUES (8, '关键词', 'web_keyword', '关键词...', '', 5, 3, NULL, NULL, NULL);
INSERT INTO `system_config` VALUES (9, '邮箱账号', 'email', 'xxx@email.com', '', 1, 2, NULL, NULL, NULL);
INSERT INTO `system_config` VALUES (10, '备案编号', 'record_number', 'Copyright© 2014-2019 | Powered by ***1.1 | 粤ICP备****号', '', 1, 1, NULL, NULL, NULL);
INSERT INTO `system_config` VALUES (11, '网站描述', 'web_description', 'web', '', 1, 1, NULL, NULL, NULL);
INSERT INTO `system_config` VALUES (12, '下拉选项', 'select', '下拉3', '下拉1,下拉2,下拉3', 4, 1, NULL, NULL, NULL);
INSERT INTO `system_config` VALUES (13, '复选框', 'checkbox', 'HTML,CSS', 'AJAX,HTML,JS,CSS', 3, 1, NULL, NULL, NULL);
INSERT INTO `system_config` VALUES (14, '文本域', 'textarea', '文本域', '0', 5, 1, NULL, NULL, NULL);
INSERT INTO `system_config` VALUES (15, '默认头像', 'default_head_img', '', '', 6, 1, NULL, NULL, NULL);
INSERT INTO `system_config` VALUES (16, '描述', 'seo_description', '', '', 5, 3, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '头像',
  `username` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `full_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `gender` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '性别 1=男 2=女',
  `age` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '年龄',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态 1=启用 2=停用',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uniq_idx_username`(`username`) USING BTREE,
  UNIQUE INDEX `uniq_idx_email`(`email`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 38 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'https://cdn.qitx.net/local/myblog/user_header_image/20230517/577a53d123bc4c4f19db0cb2c6c980a8.jpg', 'admin', '超级管理员', 'dsx.emil@qq.com', '$2a$10$OcSkSCBe8D5tGL2ulmJhTe0Xboy/fzwS1H7AdmkJjpQZfeGUHr5S6', '大师兄', 1, 31, 1, '2023-09-05 17:29:36', '2023-09-12 14:47:48', NULL);
INSERT INTO `user` VALUES (2, '', 'test2', '李四', 'ls@qq.com', '$2a$10$kycb2DM8CnubeoWABNPA1O2b0MrQQDqGsEZg8EuqK4G0a63EYDr.2', '昵称', 1, 1, 1, '2023-09-06 11:38:50', '2023-09-13 09:29:27', NULL);
INSERT INTO `user` VALUES (4, '', 'test1', '测试1', 'test1@qq.com', '$2a$10$Ww.IvYhlDpNt6Uq07X5W0OswksocMpae9dmaE2TaHclINQoBUF3Fq', '昵称', 1, 22, 1, '2023-09-07 17:48:39', '2023-09-12 09:52:47', '2023-09-12 09:53:12');
INSERT INTO `user` VALUES (10, '', 'dsx', '大师兄', 'dsx@qq.com', '$2a$10$Y2FUvgUMpMlJ5h/oooH7OOdInCZgheFQaiVkKu0Wx6YcXhiylAT3a', '大师兄1', 1, 0, 1, '2024-07-22 17:34:36', '2024-07-22 17:34:36', NULL);
INSERT INTO `user` VALUES (11, '', 'admin1', '张三1', 'zs1@qq.com', '$2a$10$OcSkSCBe8D5tGL2ulmJhTe0Xboy/fzwS1H7AdmkJjpQZfeGUHr5S6', 'dsx', 1, 28, 1, '2023-09-05 17:29:36', '2023-09-12 14:47:48', NULL);
INSERT INTO `user` VALUES (12, '', 'test3', '李四1', 'ls3@qq.com', '$2a$10$kycb2DM8CnubeoWABNPA1O2b0MrQQDqGsEZg8EuqK4G0a63EYDr.2', '昵称', 1, 1, 1, '2023-09-06 11:38:50', '2023-09-13 09:29:27', NULL);
INSERT INTO `user` VALUES (13, '', 'test4', '测试2', 'test4@qq.com', '$2a$10$Ww.IvYhlDpNt6Uq07X5W0OswksocMpae9dmaE2TaHclINQoBUF3Fq', '昵称', 1, 22, 1, '2023-09-07 17:48:39', '2023-09-12 09:52:47', '2023-09-12 09:53:12');
INSERT INTO `user` VALUES (14, '', 'dsx1', '大师兄1', 'dsx1@qq.com', '$2a$10$Y2FUvgUMpMlJ5h/oooH7OOdInCZgheFQaiVkKu0Wx6YcXhiylAT3a', '大师兄1', 1, 0, 1, '2024-07-22 17:34:36', '2024-07-22 17:34:36', NULL);
INSERT INTO `user` VALUES (15, '', 'admin2', '张三2', 'zs2@qq.com', '$2a$10$OcSkSCBe8D5tGL2ulmJhTe0Xboy/fzwS1H7AdmkJjpQZfeGUHr5S6', 'dsx', 1, 28, 1, '2023-09-05 17:29:36', '2023-09-12 14:47:48', NULL);
INSERT INTO `user` VALUES (16, '', 'test5', '李四5', 'ls5@qq.com', '$2a$10$kycb2DM8CnubeoWABNPA1O2b0MrQQDqGsEZg8EuqK4G0a63EYDr.2', '昵称', 1, 1, 1, '2023-09-06 11:38:50', '2023-09-13 09:29:27', NULL);
INSERT INTO `user` VALUES (17, '', 'test6', '测试6', 'test6@qq.com', '$2a$10$Ww.IvYhlDpNt6Uq07X5W0OswksocMpae9dmaE2TaHclINQoBUF3Fq', '昵称', 1, 22, 1, '2023-09-07 17:48:39', '2023-09-12 09:52:47', '2023-09-12 09:53:12');
INSERT INTO `user` VALUES (18, '', 'dsx2', '大师兄2', 'dsx2@qq.com', '$2a$10$Y2FUvgUMpMlJ5h/oooH7OOdInCZgheFQaiVkKu0Wx6YcXhiylAT3a', '大师兄1', 1, 0, 1, '2024-07-22 17:34:36', '2024-07-22 17:34:36', NULL);
INSERT INTO `user` VALUES (19, '', 'admin3', '张三3', 'zs3@qq.com', '$2a$10$OcSkSCBe8D5tGL2ulmJhTe0Xboy/fzwS1H7AdmkJjpQZfeGUHr5S6', 'dsx', 1, 28, 1, '2023-09-05 17:29:36', '2023-09-12 14:47:48', NULL);
INSERT INTO `user` VALUES (20, '', 'test7', '李四7', 'ls7@qq.com', '$2a$10$kycb2DM8CnubeoWABNPA1O2b0MrQQDqGsEZg8EuqK4G0a63EYDr.2', '昵称', 1, 1, 1, '2023-09-06 11:38:50', '2023-09-13 09:29:27', NULL);
INSERT INTO `user` VALUES (21, '', 'test8', '测试8', 'test8@qq.com', '$2a$10$Ww.IvYhlDpNt6Uq07X5W0OswksocMpae9dmaE2TaHclINQoBUF3Fq', '昵称', 1, 22, 1, '2023-09-07 17:48:39', '2023-09-12 09:52:47', '2023-09-12 09:53:12');
INSERT INTO `user` VALUES (22, '', 'admin4', '张三4', 'zs4@qq.com', '$2a$10$OcSkSCBe8D5tGL2ulmJhTe0Xboy/fzwS1H7AdmkJjpQZfeGUHr5S6', 'dsx', 1, 28, 1, '2023-09-05 17:29:36', '2023-09-12 14:47:48', NULL);
INSERT INTO `user` VALUES (23, '', 'test9', '李四9', 'ls9@qq.com', '$2a$10$kycb2DM8CnubeoWABNPA1O2b0MrQQDqGsEZg8EuqK4G0a63EYDr.2', '昵称', 1, 1, 1, '2023-09-06 11:38:50', '2023-09-13 09:29:27', NULL);
INSERT INTO `user` VALUES (24, '', 'test10', '测试10', 'test10@qq.com', '$2a$10$Ww.IvYhlDpNt6Uq07X5W0OswksocMpae9dmaE2TaHclINQoBUF3Fq', '昵称', 1, 22, 1, '2023-09-07 17:48:39', '2023-09-12 09:52:47', '2023-09-12 09:53:12');
INSERT INTO `user` VALUES (25, '', 'dsx3', '大师兄3', 'dsx3@qq.com', '$2a$10$Y2FUvgUMpMlJ5h/oooH7OOdInCZgheFQaiVkKu0Wx6YcXhiylAT3a', '大师兄1', 1, 0, 1, '2024-07-22 17:34:36', '2024-07-22 17:34:36', NULL);
INSERT INTO `user` VALUES (26, '', 'admin5', '张三5', 'zs5@qq.com', '$2a$10$OcSkSCBe8D5tGL2ulmJhTe0Xboy/fzwS1H7AdmkJjpQZfeGUHr5S6', 'dsx', 1, 28, 1, '2023-09-05 17:29:36', '2023-09-12 14:47:48', NULL);
INSERT INTO `user` VALUES (27, '', 'test11', '李四11', 'ls11@qq.com', '$2a$10$kycb2DM8CnubeoWABNPA1O2b0MrQQDqGsEZg8EuqK4G0a63EYDr.2', '昵称', 1, 1, 1, '2023-09-06 11:38:50', '2023-09-13 09:29:27', NULL);
INSERT INTO `user` VALUES (28, '', 'test12', '测试12', 'test12@qq.com', '$2a$10$Ww.IvYhlDpNt6Uq07X5W0OswksocMpae9dmaE2TaHclINQoBUF3Fq', '昵称', 1, 22, 1, '2023-09-07 17:48:39', '2023-09-12 09:52:47', '2023-09-12 09:53:12');
INSERT INTO `user` VALUES (29, '', 'dsx4', '大师兄4', 'dsx4@qq.com', '$2a$10$Y2FUvgUMpMlJ5h/oooH7OOdInCZgheFQaiVkKu0Wx6YcXhiylAT3a', '大师兄1', 1, 0, 1, '2024-07-22 17:34:36', '2024-07-22 17:34:36', NULL);
INSERT INTO `user` VALUES (30, '', 'admin6', '张三6', 'zs6@qq.com', '$2a$10$OcSkSCBe8D5tGL2ulmJhTe0Xboy/fzwS1H7AdmkJjpQZfeGUHr5S6', 'dsx', 1, 28, 1, '2023-09-05 17:29:36', '2023-09-12 14:47:48', NULL);
INSERT INTO `user` VALUES (31, '', 'test13', '李四13', 'ls13@qq.com', '$2a$10$kycb2DM8CnubeoWABNPA1O2b0MrQQDqGsEZg8EuqK4G0a63EYDr.2', '昵称', 1, 1, 1, '2023-09-06 11:38:50', '2023-09-13 09:29:27', NULL);
INSERT INTO `user` VALUES (32, '', 'test14', '测试14', 'test14@qq.com', '$2a$10$Ww.IvYhlDpNt6Uq07X5W0OswksocMpae9dmaE2TaHclINQoBUF3Fq', '昵称', 1, 22, 1, '2023-09-07 17:48:39', '2023-09-12 09:52:47', '2023-09-12 09:53:12');
INSERT INTO `user` VALUES (33, '', 'dsx5', '大师兄5', 'dsx5@qq.com', '$2a$10$Y2FUvgUMpMlJ5h/oooH7OOdInCZgheFQaiVkKu0Wx6YcXhiylAT3a', '大师兄1', 1, 0, 1, '2024-07-22 17:34:36', '2024-07-22 17:34:36', NULL);
INSERT INTO `user` VALUES (34, '', 'admin7', '张三7', 'zs7@qq.com', '$2a$10$OcSkSCBe8D5tGL2ulmJhTe0Xboy/fzwS1H7AdmkJjpQZfeGUHr5S6', 'dsx', 1, 28, 1, '2023-09-05 17:29:36', '2023-09-12 14:47:48', NULL);
INSERT INTO `user` VALUES (35, '', 'test15', '李四15', 'ls15@qq.com', '$2a$10$kycb2DM8CnubeoWABNPA1O2b0MrQQDqGsEZg8EuqK4G0a63EYDr.2', '昵称22', 1, 1, 1, '2023-09-06 11:38:50', '2023-09-13 09:29:27', NULL);
INSERT INTO `user` VALUES (36, '', 'test111', '测试16', 'test16@qq.com', '$2a$10$Ww.IvYhlDpNt6Uq07X5W0OswksocMpae9dmaE2TaHclINQoBUF3Fq', '昵称111', 1, 22, 1, '2023-09-07 17:48:39', '2023-09-12 09:52:47', '2023-09-12 09:53:12');
INSERT INTO `user` VALUES (37, '', 'testlisi', '测试李四', 'testlisi@qq.com', '$2a$10$1ED8aAx2IyYXUczXNh/2h.lAWcZdjXZShKBse6/0UBRDmUG1Y8j5G', '测试李四', 1, 31, 1, '2025-05-16 10:43:32', '2025-05-16 16:34:07', '2025-05-16 16:38:16');

SET FOREIGN_KEY_CHECKS = 1;
