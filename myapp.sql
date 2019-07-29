/*
 Navicat Premium Data Transfer

 Source Server         : mumuim
 Source Server Type    : MySQL
 Source Server Version : 50560
 Source Host           : localhost:3306
 Source Schema         : myapp

 Target Server Type    : MySQL
 Target Server Version : 50560
 File Encoding         : 65001

 Date: 29/07/2019 09:10:59
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for app_admin
-- ----------------------------
DROP TABLE IF EXISTS `app_admin`;
CREATE TABLE `app_admin`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '管理员',
  `username` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `password` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `create_time` int(11) NOT NULL DEFAULT 0,
  `role` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `parent` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `money` decimal(10, 2) NULL DEFAULT 0.00,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of app_admin
-- ----------------------------
INSERT INTO `app_admin` VALUES (1, 'qweqwe', 'efe6398127928f1b2e9ef3207fb82663', 1554872895, '超级管理员', '0', 0.00);
INSERT INTO `app_admin` VALUES (2, 'asdasd', 'efe6398127928f1b2e9ef3207fb82663', 1554872895, '管理员', '0', 0.00);
INSERT INTO `app_admin` VALUES (3, 'zxczxc', 'efe6398127928f1b2e9ef3207fb82663', 1554872895, '管理员', '0', 0.00);
INSERT INTO `app_admin` VALUES (4, 'qazqaz', 'efe6398127928f1b2e9ef3207fb82663', 1554872895, '管理员', '0', 0.00);
INSERT INTO `app_admin` VALUES (5, 'aaaaaa', 'efe6398127928f1b2e9ef3207fb82663', 1554872895, '管理员', '0', 0.00);
INSERT INTO `app_admin` VALUES (6, 'qqqqq', 'd41d8cd98f00b204e9800998ecf8427e', 1554872895, '财务', '0', 0.00);

-- ----------------------------
-- Table structure for app_menu
-- ----------------------------
DROP TABLE IF EXISTS `app_menu`;
CREATE TABLE `app_menu`  (
  `id` int(10) UNSIGNED NOT NULL COMMENT '菜单管理',
  `menu_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级菜单',
  `name` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `action` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '菜单路径',
  `parent_id` int(10) NOT NULL DEFAULT 0 COMMENT '祖父id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `menu_id`(`menu_id`) USING BTREE,
  INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of app_menu
-- ----------------------------
INSERT INTO `app_menu` VALUES (1, 0, '管理员管理', 'admin', 0);
INSERT INTO `app_menu` VALUES (2, 0, '用户管理', 'users', 0);
INSERT INTO `app_menu` VALUES (101, 1, '管理员列表', 'admin/lst', 0);
INSERT INTO `app_menu` VALUES (102, 1, '角色列表', 'role/lst', 0);
INSERT INTO `app_menu` VALUES (103, 1, '管理员日志', 'admin/log', 0);
INSERT INTO `app_menu` VALUES (201, 2, '用户列表', 'users/lst', 0);
INSERT INTO `app_menu` VALUES (10101, 101, '管理员添加', 'admin/add', 1);
INSERT INTO `app_menu` VALUES (10102, 101, '管理员修改', 'admin/edit', 1);
INSERT INTO `app_menu` VALUES (10103, 101, '管理员删除', 'admin/remove', 1);
INSERT INTO `app_menu` VALUES (10201, 102, '角色添加', 'role/add', 1);
INSERT INTO `app_menu` VALUES (10202, 102, '权限修改', 'role/edit', 1);
INSERT INTO `app_menu` VALUES (10203, 102, '角色删除', 'role/remove', 1);
INSERT INTO `app_menu` VALUES (20101, 201, '用户添加', 'users/add', 2);
INSERT INTO `app_menu` VALUES (20102, 201, '用户修改', 'users/edit', 2);
INSERT INTO `app_menu` VALUES (20103, 201, '用户删除', 'users/remove', 2);

-- ----------------------------
-- Table structure for app_role
-- ----------------------------
DROP TABLE IF EXISTS `app_role`;
CREATE TABLE `app_role`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色名称',
  `name` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `minu_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `create_time` int(10) NOT NULL DEFAULT 0,
  `create_admin` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of app_role
-- ----------------------------
INSERT INTO `app_role` VALUES (1, '管理员', '1,101,102,103,10101,10102,10103,10201,10202,10203,2,201,20101,20102,20103', 1554873202, 'qweqwe');
INSERT INTO `app_role` VALUES (2, '财务', '2,201,20101,20102,20103', 1554873202, 'qweqwe');

-- ----------------------------
-- Table structure for app_users
-- ----------------------------
DROP TABLE IF EXISTS `app_users`;
CREATE TABLE `app_users`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户名称',
  `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `submit_password` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '提现密码',
  `money` decimal(11, 5) NOT NULL DEFAULT 0.00000 COMMENT '金额',
  `total_money` decimal(11, 5) NOT NULL DEFAULT 0.00000 COMMENT '总金额记录',
  `reg_time` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `login_time` int(10) NULL DEFAULT 0,
  `login_ip` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `parent_id` mediumint(9) NOT NULL DEFAULT 0,
  `qq` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `nickname` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '用户昵称',
  `sex` tinyint(1) UNSIGNED NOT NULL DEFAULT 0,
  `phone` varchar(12) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `province` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '省',
  `headimgurl` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '头像',
  `city` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '所住城市',
  `user_status` tinyint(2) NULL DEFAULT 1 COMMENT '用户是否冻结：0否，1是',
  `frozen_time` int(10) NULL DEFAULT 0 COMMENT '冻结时间',
  `login_code` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '登录标识',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE,
  INDEX `parent_id`(`parent_id`) USING BTREE,
  INDEX `nickename`(`nickname`) USING BTREE,
  INDEX `login_code`(`login_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
