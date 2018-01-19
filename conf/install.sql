-- MySQL dump 10.13  Distrib 5.7.17, for osx10.12 (x86_64)
--
-- Host: localhost    Database: grabc_example
-- ------------------------------------------------------
-- Server version 5.7.17

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `rabc_assignment_permission`
--

DROP TABLE IF EXISTS `rabc_assignment_permission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rabc_assignment_permission` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `permission_id` int(11) unsigned NOT NULL COMMENT '权限ID',
  `role_id` int(11) unsigned NOT NULL COMMENT '角色ID',
  `create_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rabc_assignment_permission`
--

LOCK TABLES `rabc_assignment_permission` WRITE;
/*!40000 ALTER TABLE `rabc_assignment_permission` DISABLE KEYS */;
INSERT INTO `rabc_assignment_permission` VALUES (5,2,1,1512723764),(6,1,2,1512723802);
/*!40000 ALTER TABLE `rabc_assignment_permission` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rabc_assignment_role`
--

DROP TABLE IF EXISTS `rabc_assignment_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rabc_assignment_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户ID',
  `role_id` int(11) unsigned NOT NULL COMMENT '角色ID',
  `create_at` int(11) unsigned NOT NULL COMMENT '授权时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rabc_assignment_role`
--

LOCK TABLES `rabc_assignment_role` WRITE;
/*!40000 ALTER TABLE `rabc_assignment_role` DISABLE KEYS */;
INSERT INTO `rabc_assignment_role` VALUES (2,1,1,1512723609),(4,2,2,1512730276);
/*!40000 ALTER TABLE `rabc_assignment_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rabc_assignment_route`
--

DROP TABLE IF EXISTS `rabc_assignment_route`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rabc_assignment_route` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `route_id` int(11) unsigned NOT NULL COMMENT '路由id',
  `permission_id` int(11) unsigned NOT NULL COMMENT '权限ID',
  `create_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rabc_assignment_route`
--

LOCK TABLES `rabc_assignment_route` WRITE;
/*!40000 ALTER TABLE `rabc_assignment_route` DISABLE KEYS */;
INSERT INTO `rabc_assignment_route` VALUES (1,8,1,1512723066),(6,14,2,1512723747),(8,15,1,1512730259),(9,17,1,1512730398),(10,18,1,1512750255),(15,58,2,1513590036);
/*!40000 ALTER TABLE `rabc_assignment_route` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rabc_menu`
--

DROP TABLE IF EXISTS `rabc_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rabc_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(128) NOT NULL COMMENT '菜单名称',
  `parent` int(11) unsigned DEFAULT NULL COMMENT '父级菜单ID',
  `url` varchar(255) DEFAULT NULL COMMENT '菜单地址',
  `order` int(11) unsigned DEFAULT NULL COMMENT '菜单排序',
  `create_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `icon` varchar(245) DEFAULT NULL COMMENT '菜单前面的icon代码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rabc_menu`
--

LOCK TABLES `rabc_menu` WRITE;
/*!40000 ALTER TABLE `rabc_menu` DISABLE KEYS */;
INSERT INTO `rabc_menu` VALUES (15,'权限管理',16,'permission/index',0,1513242204,'fa-circle-o'),(16,'系统权限',0,'site/index',0,1513242218,'fa-cogs'),(17,'角色管理',16,'role/index',0,1513565748,'fa-circle-o'),(18,'用户授权',16,'assignment/index',0,1513565760,'fa-circle-o'),(19,'菜单管理',16,'menu/index',0,1513565785,'fa-circle-o'),(20,'路由管理',16,'route/index',1,1513567033,'fa-circle-o'),(21,'回到首页',0,'site/index',2,1516246400,'fa-dashboard');
/*!40000 ALTER TABLE `rabc_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rabc_permission`
--

DROP TABLE IF EXISTS `rabc_permission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rabc_permission` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(50) NOT NULL COMMENT '权限名称',
  `create_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `description` varchar(200) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rabc_permission`
--

LOCK TABLES `rabc_permission` WRITE;
/*!40000 ALTER TABLE `rabc_permission` DISABLE KEYS */;
INSERT INTO `rabc_permission` VALUES (1,'基本权限',1512721135,'用户都具备的基本权限'),(2,'超级管理员',1512723734,'拥有全部权限');
/*!40000 ALTER TABLE `rabc_permission` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rabc_role`
--

DROP TABLE IF EXISTS `rabc_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rabc_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(100) NOT NULL COMMENT '角色名称',
  `create_at` int(11) NOT NULL COMMENT '创建时间',
  `description` varchar(200) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rabc_role`
--

LOCK TABLES `rabc_role` WRITE;
/*!40000 ALTER TABLE `rabc_role` DISABLE KEYS */;
INSERT INTO `rabc_role` VALUES (1,'超级管理员',1512721066,'超级管理员'),(2,'基础用户',1512723795,'只拥有基本权限');
/*!40000 ALTER TABLE `rabc_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rabc_route`
--

DROP TABLE IF EXISTS `rabc_route`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `rabc_route` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `url` varchar(100) NOT NULL COMMENT '路由名称',
  `create_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=68 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rabc_route`
--

LOCK TABLES `rabc_route` WRITE;
/*!40000 ALTER TABLE `rabc_route` DISABLE KEYS */;
INSERT INTO `rabc_route` VALUES (14,'*/*',1512723712),(18,'site/*',1512750250),(19,'site/index',1513153290),(20,'site/login',1513153290),(21,'site/logout',1513153290),(22,'site/nopermission',1513153290),(24,'assignment/ajaxadd',1513565690),(25,'assignment/ajaxremove',1513565690),(26,'assignment/index',1513565690),(27,'assignment/user',1513565690),(29,'menu/index',1513565690),(30,'permission/*',1513565690),(31,'permission/add',1513565690),(32,'permission/ajaxaddroute',1513565690),(33,'permission/ajaxremoveroute',1513565690),(34,'permission/assignment',1513565690),(35,'permission/index',1513565690),(36,'role/*',1513565690),(37,'role/ajaxassignment',1513565690),(38,'role/ajaxunassignment',1513565690),(39,'role/assignment',1513565690),(40,'role/index',1513565690),(41,'route/*',1513565690),(42,'route/ajaxadd',1513565691),(43,'route/ajaxremove',1513565691),(44,'route/index',1513565691),(45,'site/get',1513565691),(46,'site/post',1513565691),(47,'user/*',1513565691),(48,'user/modifypassword',1513565691),(65,'assignment/*',1513594633),(66,'user/checkpermision',1516180714),(67,'menu/*',1516182098);
/*!40000 ALTER TABLE `rabc_route` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `phone` varchar(50) NOT NULL COMMENT '用户手机号',
  `password` varchar(150) NOT NULL COMMENT '登录密码',
  `real_name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'18888888888','e10adc3949ba59abbe56e057f20f883e','icodyi'),(2,'17777777777','e10adc3949ba59abbe56e057f20f883e','测试账户');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-01-19 12:56:32
