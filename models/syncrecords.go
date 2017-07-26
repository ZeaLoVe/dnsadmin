package models

import (
	"fmt"
	"log"
	"path"
	"strings"
)

func setService(name string, ip string, ttl uint64) error {
	tmpList := strings.Split(name, ".")
	for i, j := 0, len(tmpList)-1; i < j; i, j = i+1, j-1 {
		tmpList[i], tmpList[j] = tmpList[j], tmpList[i]
	}
	key := path.Join(append([]string{"/skydns/"}, tmpList...)...)
	value := "{\"host\":\"" + ip + "\"}"
	fmt.Println("insert key: " + key)
	fmt.Println("insert value: " + value)
	return DefaultBackend.UpdateKV(key, value, ttl)
}

func deleteService(name string) error {
	tmpList := strings.Split(name, ".")
	for i, j := 0, len(tmpList)-1; i < j; i, j = i+1, j-1 {
		tmpList[i], tmpList[j] = tmpList[j], tmpList[i]
	}
	key := path.Join(append([]string{"/skydns/"}, tmpList...)...)
	_, err := DefaultBackend.DeleteKey(key)
	return err
}

func DeleteService(name string) error {
	return deleteService(name)
}

func SetRecords(name string, content string, ttl uint64) error {
	return setService(name, content, ttl)
}

func UnSetRecords(name string) error {
	return deleteService(name)
}

func Sync(rec Records) error {
	O.Read(&rec)
	if rec.Disabled == 1 { //停用
		return deleteService(rec.Name)
	}
	//	if rec.Ttl == 0 {
	//		rec.Ttl = 360000000
	//	}
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
