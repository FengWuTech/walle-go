package produce

import (
	"errors"
	"net"
	"os"
	"strconv"
	"time"
)

// ProduceLogID  产生logID
// Deprecated  alsosee NewLogID
func ProduceLogID() string {
	return NewLogID()
}

// NewLogID 创建LogID
func NewLogID() string {
	return strconv.FormatInt(NewLogIDInt(), 10)
}

// NewLogIDInt 创建LogID
func NewLogIDInt() int64 {
	usec := uint64(time.Now().UnixNano())
	return int64(usec&0x7FFFFFFF | 0x80000000)

}

// LocalIP 获取本地IP
func LocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return os.Hostname()
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("fail to get local ip")
}
