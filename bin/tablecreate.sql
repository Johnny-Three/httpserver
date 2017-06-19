drop table student;
CREATE TABLE `student` (
  `sid` char(5) NOT NULL,
  `class` tinyint(3) NOT NULL,
  `score` tinyint(3) NOT NULL,
  PRIMARY KEY (`sid`)
) ENGINE=MyISAM AUTO_INCREMENT=16040 DEFAULT CHARSET=gbk;


drop table class;
CREATE TABLE `class` (
  `class` tinyint(3) NOT NULL,
  `teacher` varchar(20) NOT NULL,
  PRIMARY KEY (`class`,`teacher`)
) ENGINE=MyISAM AUTO_INCREMENT=16040 DEFAULT CHARSET=gbk;
