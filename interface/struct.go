package _interface

// HostInfo 需要上传的主机信息
type HostInfo struct {
	IP      string
	Comment string
}

// FPingInfo 单个FPing的结果
type FPingInfo struct {
	Tss            int64
	Avg, Min, Max  float64
	Src, Dst, Loss string
}

// RpcResponse rpc response信息
type RpcResponse struct {
	Status        int
	AddrArray     []string
	PingInfoArray []*FPingInfo
}

// PingMatrixService rpc调用的服务接口
type PingMatrixService interface {
	// UploadHostInfo 上报自身信息到rpc server
	UploadHostInfo(hostInfo HostInfo, response *RpcResponse) error
	// GetHostsInfo 从 rpc server端查询所有主机信息
	GetHostsInfo(request string, response *RpcResponse) error
	// UploadFPingInfo 上报FPing的结果
	UploadFPingInfo(fPingInfo []FPingInfo, response *RpcResponse) error
}
