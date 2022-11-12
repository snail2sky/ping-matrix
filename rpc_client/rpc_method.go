package main

import (
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	inter "pingmatrix/rpc_interface"
)

type rpcConn struct {
	*rpc.Client
}

func (r *rpcConn) connect(rpcAddr string) {
	conn, err := jsonrpc.Dial("tcp", rpcAddr)
	r.Client = conn
	if err != nil {
		log.Fatalln("dialling error:", err)
	}
}

// UploadHostInfo 上传本机信息到 rpc server
func (r *rpcConn) UploadHostInfo(hostInfo inter.HostInfo, response *inter.RpcResponse) error {
	log.Printf("Will upload host info is %#v\n", hostInfo)
	err := r.Call("PingMatrix.UploadHostInfo", hostInfo, response)
	log.Printf("Upload host info response is %#v\n", *response)
	if err != nil {
		log.Println(err.Error())
	}
	return nil
}

// GetHostsInfo 获取所有已经存在的机器信息
func (r *rpcConn) GetHostsInfo(request string, response *inter.RpcResponse) error {
	err := r.Call("PingMatrix.GetHostsInfo", request, response)
	log.Printf("Get all hosts info are %#v\n", response.AddrArray)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

// UploadFPingInfo 上传当前机器所有的FPing信息
func (r *rpcConn) UploadFPingInfo(fPingInfoArr []inter.FPingInfo, response *inter.RpcResponse) error {
	log.Printf("Will upload fpinginfo %#v\n", fPingInfoArr)
	err := r.Call("PingMatrix.UploadFPingInfo", fPingInfoArr, response)
	log.Printf("Upload fping info response is %#v\n", *response)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

// NewRpcConn 得到一个rpc连接对象
func NewRpcConn(addr string) *rpcConn {
	conn := &rpcConn{}
	conn.connect(addr)
	return conn
}
