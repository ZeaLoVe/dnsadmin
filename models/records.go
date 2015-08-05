package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"time"
)

const SKYDNS_DOMAIN = ".sdp"

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

var O orm.Ormer

func init() {
	orm.RegisterDataBase("default", "mysql", "user_9g8j2stroh:Oqh16lsuRj@tcp(172.24.133.50:3306)/dev_mysql_clj0727?charset=utf8", 30)
	orm.RegisterModel(new(Records))
	orm.RunSyncdb("default", false, true)
	O = orm.NewOrm()
}

func Save(rec Records) error {
	if rec.Name == "" {
		return fmt.Errorf("key name is empty")
	}
	if !strings.HasSuffix(rec.Name, SKYDNS_DOMAIN) {
		return fmt.Errorf("Not domain end with %v", SKYDNS_DOMAIN)
	}
	var tmp Records
	tmp.Name = rec.Name
	err := O.Read(&tmp)
	rec.Change_date = int(time.Now().Unix())
	if err == nil {
		num, err := O.Update(&rec, "Content", "Ttl", "Change_date", "Modifier_ip")
		log.Printf("Updata %v lines in save err:%v\n", num, err)
		if err != nil {
			return err
		} else {
			return Sync(rec)
		}
	} else {
		_, err := O.Insert(&rec)
		if err != nil {
			return err
		} else {
			return Sync(rec)
		}
	}
}

func Enable(rec Records) error {
	rec.Disabled = 0
	rec.Change_date = int(time.Now().Unix())
	num, err := O.Update(&rec, "Disabled", "Change_date", "Modifier_ip")
	if num >= 1 {
		return Sync(rec)
	} else {
		return err
	}
}

func Disable(rec Records) error {
	rec.Disabled = 1
	rec.Change_date = int(time.Now().Unix())
	num, err := O.Update(&rec, "Disabled", "Change_date", "Modifier_ip")
	if num >= 1 {
		return Sync(rec)
	} else {
		return err
	}
}

func Search(keyword string) ([]Records, error) {
	var res []Records
	var result []Records
	_, err := O.Raw("select * from records").QueryRows(&res)
	if err != nil {
		return nil, err
	}
	for _, rec := range res {
		if strings.Contains(rec.Name, keyword) {
			result = append(result, rec)
		}
	}
	return result, nil
}
