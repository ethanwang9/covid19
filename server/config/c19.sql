/*
 Navicat Premium Data Transfer

 Source Server         : c19
 Source Server Type    : MySQL
 Source Server Version : 50736 (5.7.36)
 Source Host           : localhost:20000
 Source Schema         : c19

 Target Server Type    : MySQL
 Target Server Version : 50736 (5.7.36)
 File Encoding         : 65001

 Date: 03/12/2022 17:22:23
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for c19_sys
-- ----------------------------
DROP TABLE IF EXISTS `c19_sys`;
CREATE TABLE `c19_sys`  (
  `id` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `copyright` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `gov_no1` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `gov_no2` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `mp_url` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `mp_img` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `mail` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `blog` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of c19_sys
-- ----------------------------
INSERT INTO `c19_sys` VALUES ('9', 'Ethan.Wang', '', '', 'https://mp.weixin.qq.com/mp/profile_ext?action=home&__biz=Mzg4Njg3NjcyMw==&scene=124#wechat_redirect', 'https://compeition-excute.oss-cn-beijing.aliyuncs.com/poss/69fef69c0de6dd40f45fe03fef87c0fd.jpg', 'work@ethan9.cn', 'https://www.ethan9.cn', '2022-11-25 17:26:26', '2022-11-27 11:22:09');

-- ----------------------------
-- Table structure for c19_user
-- ----------------------------
DROP TABLE IF EXISTS `c19_user`;
CREATE TABLE `c19_user`  (
  `uid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `wx_openid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `avatar` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `level` tinyint(2) NOT NULL,
  `location` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`uid`) USING BTREE,
  UNIQUE INDEX `wx_openid`(`wx_openid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of c19_user
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
