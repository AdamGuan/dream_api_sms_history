-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.1.73 - Source distribution
-- 服务器操作系统:                      redhat-linux-gnu
-- HeidiSQL 版本:                  9.1.0.4867
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- 导出 dream_api_sms_history 的数据库结构
DROP DATABASE IF EXISTS `dream_api_sms_history`;
CREATE DATABASE IF NOT EXISTS `dream_api_sms_history` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `dream_api_sms_history`;


-- 导出  表 dream_api_sms_history.t_config_response 结构
DROP TABLE IF EXISTS `t_config_response`;
CREATE TABLE IF NOT EXISTS `t_config_response` (
  `F_response_no` smallint(5) NOT NULL COMMENT '响应code',
  `F_response_msg` varchar(50) NOT NULL COMMENT '响应信息'
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='api的响应配置';

-- 正在导出表  dream_api_sms_history.t_config_response 的数据：2 rows
DELETE FROM `t_config_response`;
/*!40000 ALTER TABLE `t_config_response` DISABLE KEYS */;
INSERT INTO `t_config_response` (`F_response_no`, `F_response_msg`) VALUES
	(0, '成功'),
	(1, '失败');
/*!40000 ALTER TABLE `t_config_response` ENABLE KEYS */;


-- 导出  表 dream_api_sms_history.t_sms_send_history 结构
DROP TABLE IF EXISTS `t_sms_send_history`;
CREATE TABLE IF NOT EXISTS `t_sms_send_history` (
  `F_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `F_pkg` varchar(50) NOT NULL COMMENT '包名',
  `F_debug` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1debug,0正式',
  `F_create_datetime` datetime NOT NULL COMMENT '时间',
  PRIMARY KEY (`F_id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='短信发送记录';

-- 正在导出表  dream_api_sms_history.t_sms_send_history 的数据：0 rows
DELETE FROM `t_sms_send_history`;
/*!40000 ALTER TABLE `t_sms_send_history` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_sms_send_history` ENABLE KEYS */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
