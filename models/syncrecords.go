package models

import (
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"log"
	"net"
	"path"
	"strings"
)

var machines []string

func getipByName(name string) []net.IP {
	ns, err := net.LookupIP(name)
	if err != nil {
		fmt.Println("no ips for the name")
		return ns
	} else {
		fmt.Println("get ips for " + name)
		return ns
	}
}

func getData(key string) string {
	client := etcd.NewClient(machines)
	resp, err := client.Get(key, false, true)
	if err != nil {
		return err.Error()
	} else {
		if resp.Node.Value != "" {
			return "key:" + resp.Node.Key + "value:" + resp.Node.Value + "expiration:" + resp.Node.Expiration.String()
		} else {
			return ""
		}
	}
}

func setData(key string, value string, ttl uint64) error {
	client := etcd.NewClient(machines)
	_, err := client.Set(key, value, ttl)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func deleteData(key string) error {
	client := etcd.NewClient(machines)
	_, err := client.Delete(key, false)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func getService(name string) string {
	tmpList := strings.Split(name, ".")
	for i, j := 0, len(tmpList)-1; i < j; i, j = i+1, j-1 {
		tmpList[i], tmpList[j] = tmpList[j], tmpList[i]
	}
	key := path.Join(append([]string{"/skydns/"}, tmpList...)...)
	return getData(key)
}

func setService(name string, ip string, ttl uint64) error {
	tmpList := strings.Split(name, ".")
	for i, j := 0, len(tmpList)-1; i < j; i, j = i+1, j-1 {
		tmpList[i], tmpList[j] = tmpList[j], tmpList[i]
	}
	key := path.Join(append([]string{"/skydns/"}, tmpList...)...)
	value := "{\"host\":\"" + ip + "\"}"
	fmt.Println("insert key: " + key)
	fmt.Println("insert value: " + value)
	return setData(key, value, ttl)
}

func deleteService(name string) error {
	tmpList := strings.Split(name, ".")
	for i, j := 0, len(tmpList)-1; i < j; i, j = i+1, j-1 {
		tmpList[i], tmpList[j] = tmpList[j], tmpList[i]
	}
	key := path.Join(append([]string{"/skydns/"}, tmpList...)...)
	return deleteData(key)
}

func init() {
	machines = []string{"http://etcd.product.sdp.nd:2379"} //set default
}

func Sync(rec Records) error {
	O.Read(&rec)
	if rec.Disabled == 1 { //停用
		return deleteService(rec.Name)
	}
	if rec.Ttl == 0 {
		rec.Ttl = 360000000
	}
	flag := setService(rec.Name, rec.Content, uint64(rec.Ttl))
	log.Printf("SetService called with %v", flag)
	return flag
}

func SyncAll() (int, error) {
	records, err := Search("")
	if err != nil {
		return 0, fmt.Errorf("can't get data from database")
	}
	var successNum int
	var hasError string
	for _, rec := range records {
		if rec.Disabled == 1 {
			continue
		}
		err = Sync(rec)
		if err == nil {
			successNum++
		} else {
			hasError = fmt.Sprintf("%v %v", hasError, rec.Name)
		}
	}
	if hasError == "" {
		return successNum, nil
	} else {
		return successNum, fmt.Errorf(hasError)
	}
}
