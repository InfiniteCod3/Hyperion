package proxy

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var proxyRegex = regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}):(\d+)`)

func LoadFromFile(protocol ProxyProtocol, path string, manager *ProxyManager) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		match := proxyRegex.FindStringSubmatch(line)
		if len(match) != 3 {
			return err
		}

		ip := match[1]
		port := match[2]

		proxy := &Proxy{
			Ip:       ip,
			Port:     port,
			Protocol: protocol,
		}

		manager.Add(proxy)
	}

	return scanner.Err()
}
