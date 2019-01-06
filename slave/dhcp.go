package main

import (
	"net"
	"time"

	dhcp "github.com/krolaw/dhcp4"
)

type DHCPHandler struct {
	options       dhcp.Options
	ip            net.IP
	leaseDuration time.Duration
}

func (h *DHCPHandler) ServeDHCP(p dhcp.Packet, msgType dhcp.MessageType, options dhcp.Options) dhcp.Packet {
	//nic := p.CHAddr().String()
	ip := parseIP4("192.168.10.2")
	switch msgType {
	case dhcp.Discover:
		return dhcp.ReplyPacket(p, dhcp.Offer, h.ip, ip, h.leaseDuration,
			h.options.SelectOrderOrAll(options[dhcp.OptionParameterRequestList]))
	case dhcp.Request:
		return dhcp.ReplyPacket(p, dhcp.ACK, h.ip, ip, h.leaseDuration,
			h.options.SelectOrderOrAll(options[dhcp.OptionParameterRequestList]))
	}
	return nil
}

func RunDHCPServer() {
	handler := &DHCPHandler{
		ip:            parseIP4("192.168.10.1"),
		leaseDuration: 24 * time.Hour,
		options: dhcp.Options{
			dhcp.OptionSubnetMask: parseIP4("255.255.255.0"),
		},
	}
	for {
		dhcp.ListenAndServeIf("usb0", handler)
	}
}

func parseIP4(raw string) net.IP {
	ip := net.ParseIP(raw)
	if ip == nil {
		return nil
	}
	return ip[12:16]
}
