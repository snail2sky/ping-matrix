package main

import (
	"log"
	inter "pingmatrix/interface"
)

func uploadHostInfo(hostInfo inter.HostInfo) {
	log.Printf("DB: uploadHostInfo insert %#v\n", hostInfo)
	sql := `INSERT INTO host (host, comment) VALUE (?, ?)  ON DUPLICATE KEY UPDATE comment=(?); `
	row, err := db.Prepare(sql)
	if err != nil {
		log.Println(err.Error())
	}
	_, err = row.Exec(hostInfo.IP, hostInfo.Comment, hostInfo.Comment)
	if err != nil {
		log.Println(err.Error())
	}
}

func getHostsInfo() []string {
	log.Println("DB: getHostsInfo query host info")
	var host string
	var hostArray = make([]string, 0)
	var sql = `SELECT host FROM host`
	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&host)
		if err != nil {
			log.Println(err.Error())
		}
		hostArray = append(hostArray, host)
	}
	return hostArray
}

func uploadFPingInfo(fPingInfoArr []inter.FPingInfo) {
	log.Println("DB: uploadFPingInfo insert data")
	for _, singlePing := range fPingInfoArr {
		var sql = `INSERT valu (tss, src, dst, loss, rttmin, rttavg, rttmax) value (?, ?, ?, ?, ?, ?, ?)`
		row, err := db.Prepare(sql)
		if err != nil {
			log.Println(err.Error())
		}
		log.Printf("Will instert data are %#v\n", singlePing)
		_, err = row.Exec(singlePing.Tss, singlePing.Src, singlePing.Dst, singlePing.Loss, singlePing.Min, singlePing.Avg, singlePing.Max)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
