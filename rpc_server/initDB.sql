DROP TABLE IF EXISTS host_info;
CREATE TABLE host_info (
                           ip char(100) NOT NULL,
                           comment char(100),
                           PRIMARY KEY (ip) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS fping_result;
CREATE TABLE fping_result (
                              id int(11) NOT NULL AUTO_INCREMENT,
                              src char(100) DEFAULT NULL,
                              dst char(100) DEFAULT NULL,
                              loss char(100) DEFAULT NULL,
                              tss int(100) NOT NULL,
                              rttmin double(11,2) DEFAULT '0.00',
  rttavg double(11,2) DEFAULT '0.00',
  rttmax double(11,2) DEFAULT '0.00',
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=178150 DEFAULT CHARSET=utf8;