package config

var (
	APPPATH = "c:/windows"
)

type ProcessInfo struct {
	Name 	string
	Arg 	string
	Config 	string
}

type MasterConfig struct {
	ProcessInfo
	TcpServerConfig
}

type TcpServerConfig struct {
	Host string
}

type GatewayConfig struct {
	ProcessInfo
	TcpServerConfig
}