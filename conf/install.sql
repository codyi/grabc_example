/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50717
 Source Host           : localhost
 Source Database       : grabc_example

 Target Server Type    : MySQL
 Target Server Version : 50717
 File Encoding         : utf-8

 Date: 12/08/2017 18:56:38 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `rabc_assignment_permission`
-- ----------------------------
CREATE TABLE `rabc_assignment_permission` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `permission_id` int(11) unsigned NOT NULL COMMENT '权限ID',
  `role_id` int(11) unsigned NOT NULL COMMENT '角色ID',
  `create_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `rabc_assignment_permission`
-- ----------------------------
BEGIN;
INSERT INTO `rabc_assignment_permission` VALUES ('5', '2', '1', '1512723764'), ('6', '1', '2', '1512723802');
COMMIT;

-- ----------------------------
--  Table structure for `rabc_assignment_role`
-- ----------------------------
CREATE TABLE `rabc_assignment_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户ID',
  `role_id` int(11) unsigned NOT NULL COMMENT '角色ID',
  `create_at` int(11) unsigned NOT NULL COMMENT '授权时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `rabc_assignment_role`
-- ----------------------------
BEGIN;
INSERT INTO `rabc_assignment_role` VALUES ('2', '1', '1', '1512723609'), ('4', '2', '2', '1512730276');
COMMIT;

-- ----------------------------
--  Table structure for `rabc_assignment_route`
-- ----------------------------
CREATE TABLE `rabc_assignment_route` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `route_id` int(11) unsigned NOT NULL COMMENT '路由id',
  `permission_id` int(11) unsigned NOT NULL COMMENT '权限ID',
  `create_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `rabc_assignment_route`
-- ----------------------------
BEGIN;
INSERT INTO `rabc_assignment_route` VALUES ('1', '8', '1', '1512723066'), ('6', '14', '2', '1512723747'), ('8', '15', '1', '1512730259'), ('9', '17', '1', '1512730398');
COMMIT;

-- ----------------------------
--  Table structure for `rabc_permission`
-- ----------------------------
CREATE TABLE `rabc_permission` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(50) NOT NULL COMMENT '权限名称',
  `create_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `description` varchar(200) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `rabc_permission`
-- ----------------------------
BEGIN;
INSERT INTO `rabc_permission` VALUES ('1', '基本权限', '1512721135', '用户都具备的基本权限'), ('2', '超级管理员', '1512723734', '拥有全部权限');
COMMIT;

-- ----------------------------
--  Table structure for `rabc_role`
-- ----------------------------
CREATE TABLE `rabc_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(100) NOT NULL COMMENT '角色名称',
  `create_at` int(11) NOT NULL COMMENT '创建时间',
  `description` varchar(200) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `rabc_role`
-- ----------------------------
BEGIN;
INSERT INTO `rabc_role` VALUES ('1', '超级管理员', '1512721066', '超级管理员'), ('2', '基础用户', '1512723795', '只拥有基本权限');
COMMIT;

-- ----------------------------
--  Table structure for `rabc_route`
-- ----------------------------
CREATE TABLE `rabc_route` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `route` varchar(100) NOT NULL COMMENT '路由名称',
  `create_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `rabc_route`
-- ----------------------------
BEGIN;
INSERT INTO `rabc_route` VALUES ('14', '*/*', '1512723712'), ('15', 'user/modifypassword', '1512727621'), ('16', 'site/index', '1512730251'), ('17', 'site/*', '1512730392');
COMMIT;

-- ----------------------------
--  Table structure for `users`
-- ----------------------------
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `phone` varchar(50) NOT NULL COMMENT '用户手机号',
  `password` varchar(150) NOT NULL COMMENT '登录密码',
  `real_name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `users`
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES ('1', '18888888888', 'e10adc3949ba59abbe56e057f20f883e', 'icodyi'), ('2', '17777777777', 'e10adc3949ba59abbe56e057f20f883e', '测试账户');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
