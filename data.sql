/*
 Navicat Premium Data Transfer

 Source Server         : my8.0
 Source Server Type    : MySQL
 Source Server Version : 80014
 Source Host           : localhost:3356
 Source Schema         : coinwallet

 Target Server Type    : MySQL
 Target Server Version : 80014
 File Encoding         : 65001

 Date: 20/08/2021 09:25:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_menu`;
CREATE TABLE `admin_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL DEFAULT '0',
  `name` varchar(30) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `url` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `icon` varchar(30) COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'fa-list',
  `is_show` tinyint(4) NOT NULL DEFAULT '1',
  `sort_id` int(11) NOT NULL DEFAULT '1000',
  `log_method` varchar(8) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '不记录',
  PRIMARY KEY (`id`),
  KEY `admin_menu_url` (`url`)
) ENGINE=InnoDB AUTO_INCREMENT=102 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of admin_menu
-- ----------------------------
BEGIN;
INSERT INTO `admin_menu` VALUES (1, 0, '后台首页', 'admin/index/index', 'fa-home', 1, 99, '不记录');
INSERT INTO `admin_menu` VALUES (2, 0, '系统管理', 'admin/sys', 'fa-desktop', 1, 1099, '不记录');
INSERT INTO `admin_menu` VALUES (3, 2, '用户管理', 'admin/admin_user/index', 'fa-user', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (4, 3, '添加用户界面', 'admin/admin_user/add', 'fa-plus', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (5, 3, '修改用户界面', 'admin/admin_user/edit', 'fa-edit', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (6, 3, '删除用户', 'admin/admin_user/del', 'fa-close', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (7, 2, '角色管理', 'admin/admin_role/index', 'fa-group', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (8, 7, '添加角色界面', 'admin/admin_role/add', 'fa-plus', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (9, 7, '修改角色界面', 'admin/admin_role/edit', 'fa-edit', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (10, 7, '删除角色', 'admin/admin_role/del', 'fa-close', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (11, 7, '角色授权界面', 'admin/admin_role/access', 'fa-key', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (12, 2, '菜单管理', 'admin/admin_menu/index', 'fa-align-justify', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (13, 12, '添加菜单界面', 'admin/admin_menu/add', 'fa-plus', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (14, 12, '修改菜单界面', 'admin/admin_menu/edit', 'fa-edit', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (15, 12, '删除菜单', 'admin/admin_menu/del', 'fa-close', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (16, 2, '操作日志', 'admin/admin_log/index', 'fa-keyboard-o', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (17, 16, '日志详情', 'admin/admin_log/view', 'fa-search-plus', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (18, 2, '个人资料', 'admin/admin_user/profile', 'fa-smile-o', 1, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (55, 3, '修改头像', 'admin/admin_user/update_avatar', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (56, 3, '添加用户', 'admin/admin_user/create', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (57, 3, '修改用户', 'admin/admin_user/update', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (58, 3, '用户启用', 'admin/admin_user/enable', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (59, 3, '用户禁用', 'admin/admin_user/disable', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (60, 3, '修改昵称', 'admin/admin_user/update_nickname', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (61, 3, '修改密码', 'admin/admin_user/update_password', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (62, 7, '创建角色', 'admin/admin_role/create', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (63, 7, '修改角色', 'admin/admin_role/update', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (64, 7, '启用角色', 'admin/admin_role/enable', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (65, 7, '禁用角色', 'admin/admin_role/disable', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (66, 7, '角色授权', 'admin/admin_role/access_operate', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (67, 12, '创建菜单', 'admin/admin_menu/create', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (68, 12, '修改菜单', 'admin/admin_menu/update', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (69, 0, '市场管理', 'admin/market', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (70, 69, '币种管理', 'admin/asset/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (71, 70, '添加-界面', 'admin/asset/add', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (72, 70, '编辑-界面', 'admin/asset/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (73, 70, '币种管理-创建', 'admin/asset/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (74, 70, '币种管理-更新', 'admin/asset/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (75, 70, '币种管理-删除', 'admin/asset/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (76, 69, '行情币管理', 'admin/market/asset/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (77, 76, '添加-界面', 'admin/market/asset/add', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (78, 76, '编辑-界面', 'admin/market/asset/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (79, 76, '行情币管理-创建', 'admin/market/asset/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (80, 76, '行情币管理-更新', 'admin/market/asset/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (81, 76, '行情币管理-删除', 'admin/market/asset/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (82, 69, '合约管理', 'admin/token/config/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (83, 82, '合约管理-添加', 'admin/token/config/add', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (84, 82, '合约管理-编辑', 'admin/token/config/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (85, 82, '合约管理-更新', 'admin/token/config/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (86, 82, '合约管理-删除', 'admin/token/config/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (87, 82, '合约管理-创建', 'admin/token/config/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (88, 0, '公告管理', 'admin/news', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (89, 88, '公告管理', 'admin/news/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (90, 89, '公告管理-添加', 'admin/news/add', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (91, 89, '公告管理-编辑', 'admin/news/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (92, 89, '公告管理-创建', 'admin/news/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (93, 89, '公告管理-更新', 'admin/news/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (94, 89, '公告管理-删除', 'admin/news/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (95, 69, '地址管理', 'admin/address/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (96, 95, '地址管理-编辑', 'admin/address/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (97, 95, '地址管理-更新', 'admin/address/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (98, 95, '地址管理-删除', 'admin/address/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (99, 69, '市场管理', 'admin/market/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (100, 69, '燃油管理', 'admin/gas/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (101, 69, '交易记录', 'admin/record/index', 'fa-list', 1, 1000, '不记录');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
