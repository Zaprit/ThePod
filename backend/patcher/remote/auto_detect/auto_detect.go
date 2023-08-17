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
	log.Info("Scanning for PS3s on the local network")

	defer func(conn *net.UDPConn) {
		_ = conn.Close()
	}(conn)

	payload := make([]byte, 1024)

	log.Info("We're gonna do it")
	for {
		log.Info("looking for ps3s")
		for i := 0; i < 30; i++ {
			log.Info("Check tick")
			_, _, _, addr, err := conn.ReadMsgUDP(payload, nil)
			log.Info("Sudden death")
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
							log.Info("Found PS3: ", addr.IP.String())
							FoundPS3s = append(FoundPS3s, addr.IP.String())
						}
					}
				}
			}
		}
		log.Info("Done looking for ps3s")
		time.Sleep(5 * time.Second)
	}
}
