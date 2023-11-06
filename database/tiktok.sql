/*
Navicat MySQL Data Transfer

Source Server         : tiktok
Source Server Version : 50744
Source Host           : localhost:3306
Source Database       : tiktok

Target Server Type    : MYSQL
Target Server Version : 50744
File Encoding         : 65001

Date: 2023-11-06 16:49:08
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `comment`
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '评论记录 id',
  `user_id` int(11) NOT NULL COMMENT '发布评论的用户 id',
  `video_id` int(11) NOT NULL COMMENT '评论视频的 id',
  `content` varchar(255) NOT NULL COMMENT '评论的内容',
  `action_type` int(11) NOT NULL COMMENT '评论行为，1 表示已发布评论，2 表示删除评论',
  `created_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '评论发布时间',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '评论更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_comment_user` (`user_id`) USING BTREE,
  KEY `fk_comment_vide` (`video_id`) USING BTREE,
  CONSTRAINT `fk_comment_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_comment_vide` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of comment
-- ----------------------------

-- ----------------------------
-- Table structure for `like`
-- ----------------------------
DROP TABLE IF EXISTS `like`;
CREATE TABLE `like` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '点赞记录 id',
  `user_id` int(11) NOT NULL COMMENT '点赞用户 id',
  `video_id` int(11) NOT NULL COMMENT '点赞视频 id',
  `liked` int(11) NOT NULL COMMENT '默认 1 表示已点赞，0 表示未点赞',
  `created_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '点赞记录创建时间',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '点赞记录更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_like_user` (`user_id`) USING BTREE,
  KEY `fk_like_video` (`video_id`) USING BTREE,
  CONSTRAINT `fk_like_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_like_video` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of like
-- ----------------------------

-- ----------------------------
-- Table structure for `message`
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'message 记录 id',
  `user_id` int(11) NOT NULL COMMENT '发送 message 的 user_id',
  `receiver_id` int(11) NOT NULL COMMENT '接收 message 的 user_id',
  `msg_content` varchar(255) NOT NULL COMMENT '消息内容',
  `action_type` int(11) NOT NULL COMMENT '消息行为，1 表示发送/ 2 表示撤回',
  `created_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '消息创建时间',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '消息更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_message_user_1` (`user_id`) USING BTREE,
  KEY `fk_message_user_2` (`receiver_id`) USING BTREE,
  CONSTRAINT `fk_message_user_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_message_user_2` FOREIGN KEY (`receiver_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of message
-- ----------------------------

-- ----------------------------
-- Table structure for `relation`
-- ----------------------------
DROP TABLE IF EXISTS `relation`;
CREATE TABLE `relation` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'follow 关系 id',
  `user_id` int(11) NOT NULL COMMENT '用户 id',
  `following_id` int(11) NOT NULL COMMENT 'user_id 关注的用户 id',
  `followed` int(11) NOT NULL COMMENT '默认 0 表示未关注，1 表示已关注',
  `created_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_relation_user_1` (`user_id`) USING BTREE,
  KEY `fk_relation_user_2` (`following_id`) USING BTREE,
  CONSTRAINT `fk_relation_user_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `fk_relation_user_2` FOREIGN KEY (`following_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of relation
-- ----------------------------

-- ----------------------------
-- Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户 id，自增主键',
  `name` varchar(255) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '用户密码',
  `created_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE COMMENT '用户名保证唯一'
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'admin', '47dfbc37d2577197c6db50e5e52693a27dee2f3725671e2bb917f9b38fd44795', '2023-10-31 20:46:07', '2023-10-31 20:46:07');

-- ----------------------------
-- Table structure for `video`
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '视频 id',
  `author_id` int(11) NOT NULL COMMENT '视频作者 id',
  `title` varchar(255) NOT NULL COMMENT '视频标题',
  `play_url` varchar(255) NOT NULL COMMENT '视频播放地址',
  `cover_url` varchar(255) NOT NULL COMMENT '视频封面地址',
  `created_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '视频上传时间',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '视频更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_video_user` (`author_id`) USING BTREE,
  CONSTRAINT `fk_video_user` FOREIGN KEY (`author_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of video
-- ----------------------------
