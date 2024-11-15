/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50719
 Source Host           : localhost:3306
 Source Schema         : aidraw

 Target Server Type    : MySQL
 Target Server Version : 50719
 File Encoding         : 65001

 Date: 03/03/2022 10:27:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for works
-- ----------------------------
DROP TABLE IF EXISTS `works`;
CREATE TABLE `works`  (
                              `id` int(11) NOT NULL AUTO_INCREMENT,
                              `work_id` BIGINT UNSIGNED NOT NULL COMMENT '作品ID',
                              `url` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                              `prompt` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                              `category_id` BIGINT UNSIGNED NOT NULL COMMENT '分类ID',
                              `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              PRIMARY KEY (`id`) USING BTREE,
                              UNIQUE INDEX `idx_work_url`(`url`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of works
-- ----------------------------
# INSERT INTO `works` VALUES (1, 145038507964891137, '1', '3134ed3bdfd3c8429dd86c89baee2823a5', 1, 1, 'url', '2024-11-05 21:48:13', '2024-11-05 21:48:13');


-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
                          `id` int(11) NOT NULL AUTO_INCREMENT,
                          `user_id` BIGINT UNSIGNED NOT NULL COMMENT '作者的用户ID',
                          `category_id` BIGINT UNSIGNED NOT NULL COMMENT '分类ID',
                          `category_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                          `description` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                          `cover_url` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                          `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (5, 145282843017216001, 145285714488066049, '2', '2', 'url', '2024-11-07 14:43:54', '2024-11-07 14:43:54');
INSERT INTO `category` VALUES (6, 145282843017216001, 145285840686284801, '3', '3', 'url', '2024-11-07 14:45:11', '2024-11-07 14:45:11');
INSERT INTO `category` VALUES (7, 145282843017216001, 145286079342182401, '3', '3', 'url', '2024-11-07 14:47:33', '2024-11-07 14:47:33');
INSERT INTO `category` VALUES (8, 145282843017216001, 145286528434700289, '4', '4', 'url', '2024-11-07 14:52:05', '2024-11-07 14:52:05');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
                         `id` bigint(20) NOT NULL AUTO_INCREMENT,
                         `user_id` bigint(20) NOT NULL,
                         `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                         `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                         `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                         `gender` tinyint(4) NOT NULL DEFAULT 0,
                         `avatar` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                         `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`) USING BTREE,
                         UNIQUE INDEX `idx_username`(`username`) USING BTREE,
                         UNIQUE INDEX `idx_user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 145038507964891137, '1', '3134ed3bdfd3c8429dd86c89baee2823a5', 1, 1, 'url', '2024-11-05 21:48:13', '2024-11-05 21:48:13');
INSERT INTO `user` VALUES (2, 145282843017216001, '2', '3234ed3bdfd3c8429dd86c89baee2823a5', 2, 2, 'url', '2024-11-07 14:15:24', '2024-11-07 14:15:24');

SET FOREIGN_KEY_CHECKS = 1;
