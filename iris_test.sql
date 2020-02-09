SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `username` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，1正常，2禁用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台用户表';

INSERT INTO `admin` (`id`, `username`, `password`, `status`) VALUES
(1,	'admin',	'e10adc3949ba59abbe56e057f20f883e',	1);

DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '分类名称',
  `weight` int(11) NOT NULL DEFAULT '0' COMMENT '权重（升序）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资讯分类表';

INSERT INTO `category` (`id`, `name`, `weight`) VALUES
(1,	'女装',	1),
(2,	'美妆',	2),
(5,	'女鞋',	4);

DROP TABLE IF EXISTS `news`;
CREATE TABLE `news` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '分类id',
  `title` varchar(250) NOT NULL DEFAULT '' COMMENT '标题',
  `description` varchar(500) NOT NULL DEFAULT '' COMMENT '简介',
  `pic` varchar(200) NOT NULL DEFAULT '' COMMENT '图片',
  `content` text NOT NULL COMMENT '内容',
  `weight` int(11) NOT NULL DEFAULT '0' COMMENT '权重（升序）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资讯表';

INSERT INTO `news` (`id`, `category_id`, `title`, `description`, `pic`, `content`, `weight`) VALUES
(2,	5,	'fdsafdsa',	'fdsafdsa',	'http://localhost:9091/static/upload/2019/December/11/3.jpg',	'<p>fdsafdsa<img src=\"http://localhost:9091/static/upload/2019/December/11/6.jpg\" style=\"width: 310px;\"></p>',	0),
(3,	1,	'432432',	'fdsafdsa12',	'http://localhost:9091/static/upload/2019/December/11/1575376723_3c2723396ad66ba4decb5400c1b48ccc.jpg',	'<p>fdsafdsa12</p><p><img src=\"http://localhost:9091/static/upload/2019/December/11/5.jpg\" style=\"width: 310px;\"><br></p>',	312);

-- 2019-12-11 06:16:38
