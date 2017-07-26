package models

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

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

type UersInfoResponse struct {
	UserInfo `json:"user,omitempty"`
}

type UserInfo struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Cname string `json:"cnname,omitempty"`
	Role  int    `json:"role,omitempty"`
}

type OutputRecords struct {
	Name     string `json:"name"`
	Content  string `json:"content"`
	Ttl      int    `json:"ttl"`
	Auth     string `json:"auth"`
	Disabled int    `json:"disabled"`
}

type DomiansResponse struct {
	Total int       `json:"total"`
	Items []Records `json:"rows"`
}

func GetOutput(list []Records) []OutputRecords {
	var outputlist []OutputRecords
	var tmp = OutputRecords{
		Ttl:      0,
		Disabled: 0,
	}
	for _, rec := range list {
		tmp.Auth = rec.Auth
		tmp.Content = rec.Content
		tmp.Disabled = rec.Disabled
		tmp.Name = rec.Name
		tmp.Ttl = rec.Ttl
		outputlist = append(outputlist, tmp)
	}
	return outputlist
}

var O orm.Ormer

func init() {
	connStr := beego.AppConfig.String("connstr")
	if connStr == "" {
		connStr = "user_9g8j2stroh:Oqh16lsuRj@tcp(172.24.133.50:3306)/dev_mysql_clj0727?charset=utf8"
	}
	orm.RegisterDataBase("default", "mysql", connStr, 30)
	orm.RegisterModel(new(Records))
	orm.RunSyncdb("default", false, true)
	O = orm.NewOrm()
}

func Save(rec Records) error {
	if rec.Name == "" {
		return fmt.Errorf("key name is empty")
	}
	var tmp Records
	tmp.Name = rec.Name
	err := O.Read(&tmp)
	rec.Change_date = int(time.Now().Unix())
	if err == nil {
		if err = SetRecords(rec.Name, rec.Content, uint64(rec.Ttl)); err == nil {
			num, err := O.Update(&rec, "Content", "Ttl", "Change_date", "Modifier_ip")
			log.Printf("Updata %v lines in save err:%v\n", num, err)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		if err = SetRecords(rec.Name, rec.Content, uint64(rec.Ttl)); err == nil {
			_, err := O.Insert(&rec)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func Delete(rec Records) error {
	if rec.Name == "" {
		return fmt.Errorf("key name is empty")
	}

	if err := UnSetRecords(rec.Name); err == nil {
		_, err := O.Delete(&rec)
		if err != nil {
			return err
		}
	} else {
		msg := err.Error()
		if strings.Contains(msg, "Key not found") {
			_, err := O.Delete(&rec)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func Enable(rec Records) error {
	err := O.Read(&rec)
	if err != nil {
		return err
	}
	if err := SetRecords(rec.Name, rec.Content, uint64(rec.Ttl)); err == nil {
		rec.Disabled = 0
		rec.Change_date = int(time.Now().Unix())
		num, err := O.Update(&rec, "Disabled", "Change_date", "Modifier_ip")
		if num < 1 {
			return err
		}
	} else {
		return err
	}
	return nil
}

func Disable(rec Records) error {
	if err := UnSetRecords(rec.Name); err == nil {
		rec.Disabled = 1
		rec.Change_date = int(time.Now().Unix())
		num, err := O.Update(&rec, "Disabled", "Change_date", "Modifier_ip")
		if num < 1 {
			return err
		}
	} else {
		return err
	}
	return nil
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
