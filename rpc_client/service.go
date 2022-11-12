package main

import (
	"fmt"
	"log"
	"math"
	"os/exec"
	inter "pingmatrix/interface"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func delSliceItemByIndex(index int, slice []string) []string {
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}

func fPing(ip, args string, hostAddr ...string) []inter.FPingInfo {
	var fPingInfoArr = make([]inter.FPingInfo, 0)
	var fPingInfo = new(inter.FPingInfo)
	var localIPIndex = math.MaxInt
	for i, addr := range hostAddr {
		if addr == ip {
			localIPIndex = i
			break
		}
	}
	if localIPIndex != math.MaxInt {
		hostAddr = delSliceItemByIndex(localIPIndex, hostAddr)
	}
	allHost := strings.Join(hostAddr, " ")
	allArgs := strings.Split(fmt.Sprintf("%s %s", args, allHost), " ")
	log.Println("Execute fping", strings.Join(allArgs, " "))
	cmd := exec.Command("fping", allArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("fping execute error are: %#v\n", err.Error())
	}
	re := regexp.MustCompile(`(.*) +: xmt/rcv/%loss = (.*), min/avg/max = (.*)`) //正则过滤下丢包率为100%的
	regexpMatch := re.FindAllStringSubmatch(string(output), -1)
	for _, element := range regexpMatch {
		fPingInfo.Src = ip
		fPingInfo.Tss = time.Now().Unix()
		fPingInfo.Dst = element[1]
		fPingInfo.Loss = strings.Split(element[2], "/")[2]
		fPingInfo.Min, _ = strconv.ParseFloat(strings.Split(element[3], "/")[0], 64)
		fPingInfo.Avg, _ = strconv.ParseFloat(strings.Split(element[3], "/")[1], 64)
		fPingInfo.Max, _ = strconv.ParseFloat(strings.Split(element[3], "/")[2], 64)
		fPingInfoArr = append(fPingInfoArr, *fPingInfo)
	}
	log.Printf("fping results are %#v\n", fPingInfoArr)
	log.Println("Execute fping stopped")
	return fPingInfoArr
}