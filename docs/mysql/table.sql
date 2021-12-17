-- user add
USE mysql;
CREATE USER 'imbiz'@'%' IDENTIFIED BY 'imbiz';
GRANT ALL PRIVILEGES ON *.* TO 'imbiz'@'%';


-- database add
CREATE SCHEMA `IMBiz` DEFAULT CHARACTER SET utf8 ;

-- table add
USE `IMBiz`;
CREATE TABLE `USER` (
  `USER_ID` varchar(45) NOT NULL,
  `PW` varchar(200) NOT NULL,
  `NAME` varchar(45) NOT NULL,
  `EMAIL` varchar(100) NOT NULL,
  `DEPARTMENT_ID` varchar(45) DEFAULT NULL,
  `POSITION` varchar(45) DEFAULT NULL,
  `LAST_ACCESS_DT` datetime DEFAULT NULL,
  `UPDATE_DT` datetime DEFAULT NULL,
  `CREATE_DT` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`USER_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `DEPARTMENT` (
  `ID` varchar(45) NOT NULL,
  `NAME` varchar(100) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `TEST` (
  `INDEX` int(11) NOT NULL AUTO_INCREMENT,
  `TEXT` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`INDEX`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
