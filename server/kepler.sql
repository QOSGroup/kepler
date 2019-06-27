CREATE TABLE `apply_qcp` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `qcp_chain_id` varchar(50) NOT NULL COMMENT '联盟链名称',
  `qos_chain_id` varchar(50) NOT NULL COMMENT '公链名称',
  `qcp_pub` varchar(100) NOT NULL COMMENT 'qcp公钥',
  `email` varchar(100) NOT NULL COMMENT '申请者邮箱',
  `phone` varchar(20) NOT NULL COMMENT '申请者手机号',
  `info` varchar(500) NOT NULL COMMENT '描述',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `status` int(1) NOT NULL DEFAULT '0' COMMENT '状态 0：申请 1：通过 2：不通过',
  `note` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `apply_qcp_qos_chain_id_IDX` (`qos_chain_id`,`qcp_chain_id`,`email`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='联盟链申请';

CREATE TABLE `apply_qsc` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `qsc_name` varchar(50) NOT NULL COMMENT '联盟币名称',
  `qos_chain_id` varchar(50) NOT NULL COMMENT '公链名称',
  `qsc_pub` varchar(100) NOT NULL COMMENT 'qsc公钥',
  `banker_pub` varchar(100) NOT NULL DEFAULT '' COMMENT 'banker公钥',
  `email` varchar(100) NOT NULL COMMENT '申请者邮箱',
  `phone` varchar(20) NOT NULL COMMENT '申请者手机号',
  `info` varchar(500) NOT NULL COMMENT '描述',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `status` int(1) NOT NULL DEFAULT '0' COMMENT '状态 0：申请 1：通过 2：不通过',
  `note` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `apply_qsc_qos_chain_id_IDX` (`qos_chain_id`,`qsc_name`,`email`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='联盟币申请';

CREATE TABLE `ca_qcp` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `qos_chain_id` varchar(50) NOT NULL COMMENT '公链名称',
  `qcp_chain_id` varchar(50) NOT NULL COMMENT '联盟链chainId',
  `csr` varchar(500) NOT NULL COMMENT 'qcp csr',
  `crt` varchar(1000) NOT NULL COMMENT 'qcp crt',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `apply_id` int(11) NOT NULL COMMENT '申请记录ID',
  `expire_time` datetime NOT NULL COMMENT '失效时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ca_qcp_qos_chain_id_IDX` (`qos_chain_id`,`qcp_chain_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='联盟链证书';

CREATE TABLE `ca_qsc` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `qos_chain_id` varchar(50) NOT NULL COMMENT '公链chainId',
  `name` varchar(50) NOT NULL COMMENT '名称',
  `csr` varchar(500) NOT NULL COMMENT 'qsc csr',
  `crt` varchar(1000) NOT NULL COMMENT 'qsc crt',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `apply_id` int(11) NOT NULL COMMENT '申请记录ID',
  `expire_time` datetime NOT NULL COMMENT '失效时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ca_qsc_qos_chain_id_IDX` (`qos_chain_id`,`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='联盟币证书';

CREATE TABLE `root_ca` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pub_key` varchar(100) NOT NULL COMMENT '公钥',
  `priv_key` varchar(150) NOT NULL COMMENT '私钥',
  `chain_id` varchar(50) NOT NULL COMMENT '公链ChainId',
  `type` int(1) NOT NULL DEFAULT '1' COMMENT '类型 1：根证书 2：QSC 根证书 3：QCP根证书',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `root_ca_chain_id_IDX` (`chain_id`,`type`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='根证书';