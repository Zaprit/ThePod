package auto_detect

import (
	"github.com/labstack/gommon/log"
	"net"
	"strings"
	"time"
)

var FoundPS3s []string

func AutoDetect() {
	conn, err := net.ListenMulticastUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4(239, 255, 255, 250),
		Port: 1900,
		Zone: "",
	})
	if err != nil {
		log.Error("could not listen for ssdp packets")
	}

	defer func(conn *net.UDPConn) {
		_ = conn.Close()
	}(conn)

	payload := make([]byte, 1024)

	timer := time.NewTicker(15 * time.Second)

	for true {
		for i := 0; i < 15; i++ {
			_, _, _, addr, err := conn.ReadMsgUDP(payload, nil)
			if err != nil {
				log.Warn("could not read udp message")
			}

			for _, s := range strings.Split(string(payload), "\n") {
				if strings.HasPrefix(s, "X-AV-Client-Info:") {
					if strings.Contains(s, "PLAYSTATION3") && strings.Contains(s, "Sony Computer Entertainment Inc.") {
						duplicate := false
						for _, ip := range FoundPS3s {
							if ip == addr.IP.String() {
								duplicate = true
							}
						}

						if !duplicate {
							FoundPS3s = append(FoundPS3s, addr.IP.String())
						}
					}
				}
			}
		}
		<-timer.C
	}
}
