package overrides

import (
	"encoding/json"
	"net"
	"os"
)

type connConfig struct {
	TunAddr    string `json:"tunAddr"`
	DnsAddr    string `json:"dnsAddr"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	ServerAddr string `json:"serverAddr"`
	ServerPort string `json:"serverPort"`
	AllowedIp  string `json:"allowedIp"`
	Network    string `json:"network"` //TODO: make it useful:)
	Address    string `json:"address"`
}

func newConnConfig(file *os.File) (c *connConfig, err error) {
	buff := make([]byte, 1024, 1024)
	n, err := file.Read(buff)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buff[:n], c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

var (
	file, _   = os.Open("config.json")
	config, _ = newConnConfig(file)
	conns     map[int]*net.Conn
)
