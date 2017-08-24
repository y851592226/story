CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `username` varchar(200)  CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT "" COMMENT '帐号',
  `password` varchar(200)  CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT "" COMMENT '密码',
  `email` varchar(200)  CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT "" COMMENT '邮箱',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;

