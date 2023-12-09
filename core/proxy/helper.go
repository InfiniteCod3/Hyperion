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

	lines := make(chan string)
	errors := make(chan error)
	done := make(chan bool)

	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			errors <- err
		}
		close(lines)
	}()

	for i := 0; i < 10; i++ {
		go func() {
			for line := range lines {
				line = strings.TrimSpace(line)
				match := proxyRegex.FindStringSubmatch(line)
				if len(match) != 3 {
					errors <- fmt.Errorf("invalid proxy format: %s", line)
					continue
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
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	close(errors)
	for err := range errors {
		if err != nil {
			return err
		}
	}

	return nil
}
