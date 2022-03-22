-- In order to compare the performance of seata-golang, we copy some sql in
-- https://github.com/opentrx/seata-go-samples/tree/v2/gorm/scripts

CREATE database if NOT EXISTS `dtm_bench` default character set utf8mb4 collate utf8mb4_unicode_ci;
use `dtm_bench`;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for so_item
-- ----------------------------
DROP TABLE IF EXISTS `so_item`;
CREATE TABLE `so_item` (
  `sysno` bigint(20) NOT NULL AUTO_INCREMENT,
  `so_sysno` bigint(20) DEFAULT NULL,
  `product_sysno` bigint(20) DEFAULT NULL,
  `product_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '商品名称',
  `cost_price` decimal(16,6) DEFAULT NULL COMMENT '成本价',
  `original_price` decimal(16,6) DEFAULT NULL COMMENT '原价',
  `deal_price` decimal(16,6) DEFAULT NULL COMMENT '成交价',
  `quantity` int(11) DEFAULT NULL COMMENT '数量',
  PRIMARY KEY (`sysno`)
) ENGINE=InnoDB AUTO_INCREMENT=1269646673999564801 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='订单明细表';

-- ----------------------------
-- Table structure for so_master
-- ----------------------------
DROP TABLE IF EXISTS `so_master`;
CREATE TABLE `so_master` (
  `sysno` bigint(20) NOT NULL,
  `so_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `buyer_user_sysno` bigint(20) DEFAULT NULL COMMENT '下单用户号',
  `seller_company_code` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '卖家公司编号',
  `receive_division_sysno` bigint(20) NOT NULL,
  `receive_address` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `receive_zip` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `receive_contact` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `receive_contact_phone` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `stock_sysno` bigint(20) DEFAULT NULL,
  `payment_type` tinyint(4) DEFAULT NULL COMMENT '支付方式：1，支付宝，2，微信',
  `so_amt` decimal(16,6) DEFAULT NULL COMMENT '订单总额',
  `status` tinyint(4) DEFAULT NULL COMMENT '10,创建成功，待支付；30；支付成功，待发货；50；发货成功，待收货；70，确认收货，已完成；90，下单失败；100已作废',
  `order_date` timestamp NULL DEFAULT NULL COMMENT '下单时间',
  `payment_date` timestamp NULL DEFAULT NULL COMMENT '支付时间',
  `delivery_date` timestamp NULL DEFAULT NULL COMMENT '发货时间',
  `receive_date` timestamp NULL DEFAULT NULL COMMENT '发货时间',
  `appid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '订单来源',
  `memo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `create_user` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `gmt_create` timestamp NULL DEFAULT NULL,
  `modify_user` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `gmt_modified` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`sysno`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='订单表';

-- ----------------------------
-- Table structure for undo_log
-- ----------------------------
DROP TABLE IF EXISTS `undo_log`;
CREATE TABLE `undo_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `branch_id` bigint(20) NOT NULL,
  `xid` varchar(128) NOT NULL,
  `context` varchar(128) NOT NULL,
  `rollback_info` longblob NOT NULL,
  `log_status` int(11) NOT NULL,
  `log_created` datetime NOT NULL,
  `log_modified` datetime NOT NULL,
  `ext` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_unionkey` (`xid`,`branch_id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `inventory`;
CREATE TABLE `inventory` (
  `sysno` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `product_sysno` bigint(20) unsigned NOT NULL,
  `account_qty` int(11) DEFAULT NULL COMMENT '财务库存',
  `available_qty` int(11) DEFAULT NULL COMMENT '可用库存',
  `allocated_qty` int(11) DEFAULT NULL COMMENT '分配库存',
  `adjust_locked_qty` int(11) DEFAULT NULL COMMENT '调整锁定库存',
  PRIMARY KEY (`sysno`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品库存';

-- ----------------------------
-- Records of inventory
-- ----------------------------
BEGIN;
INSERT INTO `inventory` VALUES (1, 1, 1000000, 1000000, 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product` (
  `sysno` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `product_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '品名',
  `product_title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `product_desc` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '描述',
  `product_desc_long` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '描述',
  `default_image_src` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `c3_sysno` bigint(20) DEFAULT NULL,
  `barcode` varchar(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `length` int(11) DEFAULT NULL,
  `width` int(11) DEFAULT NULL,
  `height` int(11) DEFAULT NULL,
  `weight` float DEFAULT NULL,
  `merchant_sysno` bigint(20) DEFAULT NULL,
  `merchant_productid` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '1，待上架；2，上架；3，下架；4，售罄下架；5，违规下架',
  `gmt_create` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `create_user` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '创建人',
  `modify_user` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '修改人',
  `gmt_modified` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`sysno`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品SKU';

-- ----------------------------
-- Records of product
-- ----------------------------
BEGIN;
INSERT INTO `product` VALUES (1, '刺力王', '从小喝到大的刺力王', '好喝好喝好好喝', '', 'https://img10.360buyimg.com/mobilecms/s500x500_jfs/t1/61921/34/1166/131384/5cf60a94E411eee07/1ee010f4142236c3.jpg', 0, '', 15, 5, 5, 5, 1, '', 1, '2019-05-28 03:36:17', '', '', '2019-06-06 01:37:36');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
