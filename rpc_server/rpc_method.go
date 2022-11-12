package main

import (
	"log"
	inter "pingmatrix/rpc_interface"
)

type PingMatrix struct{}

func (p *PingMatrix) UploadHostInfo(hostInfo inter.HostInfo, response *inter.RpcResponse) error {
	log.Printf("Call UploadHostInfo method. host is %s\n", hostInfo.IP)
	response.Status = 0
	uploadHostInfo(hostInfo)
	return nil
}

func (p *PingMatrix) GetHostsInfo(request string, response *inter.RpcResponse) error {
	hostArray := getHostsInfo()
	log.Printf("Call GetHostsInfo method. current hosts are %s\n", hostArray)
	response.Status = 0
	response.AddrArray = hostArray
	return nil
}

func (p *PingMatrix) UploadFPingInfo(fPingInfo []inter.FPingInfo, response *inter.RpcResponse) error {
	log.Println("Call UploadFPingInfo method.")
	uploadFPingInfo(fPingInfo)
	response.Status = 0
	return nil
}
