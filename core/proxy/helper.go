package proxy

import (
	"bufio"
	"net"
	"os"
	"strconv"
)

func LoadFromFile(protocol ProxyProtocol, path string, manager *ProxyManager) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ip, port, err := net.SplitHostPort(line)
		if err != nil {
			return err
		}

		portInt, err := strconv.Atoi(port)
		if err != nil {
			return err
		}

		proxy := &Proxy{
			Ip:       ip,
			Port:     uint16(portInt),
			Protocol: protocol,
		}

		manager.Add(proxy)
	}

	return scanner.Err()
}
