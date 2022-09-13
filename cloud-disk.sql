/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : 192.168.0.132:3306
 Source Schema         : cloud-disk

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 13/09/2022 13:54:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for repository_pool
-- ----------------------------
DROP TABLE IF EXISTS `repository_pool`;
CREATE TABLE `repository_pool`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `hash` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '文件唯一标识',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `ext` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '文件扩展名',
  `size` int(0) NULL DEFAULT NULL COMMENT '文件大小',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '文件路径',
  `create_at` timestamp(0) NULL DEFAULT NULL,
  `update_at` timestamp(0) NULL DEFAULT NULL,
  `delete_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of repository_pool
-- ----------------------------
INSERT INTO `repository_pool` VALUES (3, '7285274e-a435-4e3a-b7e0-bc3ce8d13522', '466a540e5aa15130e90a65f03fd3f49e', 'bg_s.jpg', '.jpg', 216823, 'https://1-1258379676.cos.ap-chengdu.myqcloud.com/cloud-disk/d46b333d-4f64-414b-8d4a-a8300eefa227.jpg', '2022-09-10 23:46:19', '2022-09-10 23:46:19', NULL);

-- ----------------------------
-- Table structure for share_basic
-- ----------------------------
DROP TABLE IF EXISTS `share_basic`;
CREATE TABLE `share_basic`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `user_identity` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `repository_identity` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '公共仓库唯一标识',
  `user_repository_identity` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '用户仓库池子',
  `click_num` int(0) NULL DEFAULT NULL COMMENT '点击次数',
  `expire_time` int(0) NULL DEFAULT NULL,
  `create_at` timestamp(0) NULL DEFAULT NULL,
  `update_at` timestamp(0) NULL DEFAULT NULL,
  `delete_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of share_basic
-- ----------------------------
INSERT INTO `share_basic` VALUES (3, '0c669f3f-c205-4ea9-8ad9-31d720946b13', 'USER_1', '7285274e-a435-4e3a-b7e0-bc3ce8d13522', '7db50e6b-712a-429d-8308-24b8874b7ca5', 5, 100, '2022-09-12 14:16:38', '2022-09-12 14:16:38', NULL);

-- ----------------------------
-- Table structure for user_basic
-- ----------------------------
DROP TABLE IF EXISTS `user_basic`;
CREATE TABLE `user_basic`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `create_at` timestamp(0) NULL DEFAULT NULL,
  `update_at` timestamp(0) NULL DEFAULT NULL,
  `delete_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_basic
-- ----------------------------
INSERT INTO `user_basic` VALUES (1, 'USER_1', 'zhangsan', '$2a$14$pVFyRg1s.hD6ZBJdA2CvceX1idc4KVxQfhhiYkeAgkeWW7YirD2vC', 'test@qq.com', '2022-09-10 09:05:03', '2022-09-10 09:05:05', NULL);
INSERT INTO `user_basic` VALUES (3, 'd7fd34ef-3722-416e-bbd7-690fd9c7a7fa', 'test', '$2a$14$LwFEL/tvHaPNJqDP65DyVOZgDqnU9GPjNoDn/DNMLRVECYVlQwAsi', 'xuwuruoshui@gmail.com', NULL, NULL, NULL);
INSERT INTO `user_basic` VALUES (4, '79b4f792-ffb0-4bfc-bddb-061413b0106c', 'test1', '$2a$14$g6G96F4fK6jMUE1X1guF3uOMfpX/UeVRSZpDwYZnkeppws5zRp21q', '2864758326@qq.com', NULL, NULL, NULL);

-- ----------------------------
-- Table structure for user_repository
-- ----------------------------
DROP TABLE IF EXISTS `user_repository`;
CREATE TABLE `user_repository`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `user_identity` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `parent_id` int(0) NULL DEFAULT NULL,
  `repository_identity` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `ext` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '文件或文件夹类型',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `create_at` timestamp(0) NULL DEFAULT NULL,
  `update_at` timestamp(0) NULL DEFAULT NULL,
  `delete_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_repository
-- ----------------------------
INSERT INTO `user_repository` VALUES (4, '7db50e6b-712a-429d-8308-24b8874b7ca5', 'USER_1', 6, '7285274e-a435-4e3a-b7e0-bc3ce8d13522', '.jpg', 'bgrrrf.jpg', '2022-09-11 14:26:16', '2022-09-12 10:50:52', NULL);
INSERT INTO `user_repository` VALUES (5, 'e420a5d4-f5ca-4e73-9473-9756beac5f2a', 'USER_1', 0, '', '', 'test', '2022-09-12 00:07:50', '2022-09-12 00:07:50', NULL);
INSERT INTO `user_repository` VALUES (6, '5c72f9c7-6f06-4783-8d17-e57c2fbb80cf', 'USER_1', 0, '', '', '音乐', '2022-09-12 00:08:51', '2022-09-12 00:08:51', NULL);
INSERT INTO `user_repository` VALUES (8, '0b933533-4ee2-4674-bf5c-7f032bd05cd4', 'USER_1', 5, '7285274e-a435-4e3a-b7e0-bc3ce8d13522', '.jpg', 'bg_s.jpg', '2022-09-12 16:29:38', '2022-09-12 16:29:38', NULL);

SET FOREIGN_KEY_CHECKS = 1;
