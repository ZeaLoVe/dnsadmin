# dnsadmin
a simple system for skydns records manage, using beego

DNS data is store in etcd(format see SKYDNS project)

SKYDNS as a DNS server 

Use mysql as database (SQL)

CREATE TABLE records (

  name                 VARCHAR(255) NOT NULL,
  
  type                 VARCHAR(10) DEFAULT NULL,
  
  content              VARCHAR(1000) DEFAULT NULL,
  
  ttl                  INT DEFAULT NULL,
  
  prio                 INT DEFAULT NULL,
  
  change_date          INT DEFAULT NULL,
  
  disabled             TINYINT(1) DEFAULT 0,
  
  auth                   VARCHAR(100) DEFAULT NULL,
  
  modifier_ip            VARCHAR(100) DEFAULT NULL,
  
  PRIMARY KEY (name)
  
);

Go struct of DB

type Records struct {

	Name        string `orm:"size(255);pk"`
	
	Type        string `orm:"size(10)"`
	
	Content     string `orm:"size(1000)"`
	
	Ttl         int
	
	Prio        int
	
	Change_date int
	
	Disabled    int
	
	Auth        string `orm:"size(100)"`
	
	Modifier_ip string `orm:"size(100)"`
	
}
