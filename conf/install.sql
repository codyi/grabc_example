/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50717
 Source Host           : localhost
 Source Database       : goCronJob

 Target Server Type    : MySQL
 Target Server Version : 50717
 File Encoding         : utf-8

 Date: 11/28/2017 11:37:47 AM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `users`
-- ----------------------------
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `phone` varchar(50) NOT NULL COMMENT '用户手机号',
  `password` varchar(150) NOT NULL COMMENT '登录密码',
  `create_time` int(11) unsigned NOT NULL COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL COMMENT '更新时间',
  `is_lock` tinyint(2) unsigned NOT NULL COMMENT '用户状态 0正常  1锁定',
  `real_name` varchar(50) NOT NULL COMMENT '用户的真实姓名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `users`
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES ('1', '13311582115', '6a204bd89f3c8348afd5c77c717a097a', '11111', '1111', '0', 'icodyi');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
