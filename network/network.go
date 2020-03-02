// +build linux darwin

package network

import (
	"os"

	"github.com/toyozaki/GOnetstat"
)

func getNetworkInfo() (networkInfo map[string]interface{}, err error) {
	networkInfo = make(map[string]interface{})

	macaddress, err := macAddress()
	if err != nil {
		return networkInfo, err
	}
	networkInfo["macaddress"] = macaddress

	ipAddress, err := externalIpAddress()
	if err != nil {
		return networkInfo, err
	}
	networkInfo["ipaddress"] = ipAddress

	ipAddressV6, err := externalIpv6Address()
	if err != nil {
		return networkInfo, err
	}
	// We append an IPv6 address to the payload only if IPv6 is enabled
	if ipAddressV6 != "" {
		networkInfo["ipaddressv6"] = ipAddressV6
	}

	// Only supported on linux
	if _, err := os.Stat("/proc/net"); !os.IsNotExist(err) {
		_portscan := make(map[string]interface{})
		_portscan["tcp"] = GOnetstat.Tcp()
		_portscan["tcp6"] = GOnetstat.Tcp6()
		_portscan["udp"] = GOnetstat.Udp()
		_portscan["udp6"] = GOnetstat.Udp6()
		networkInfo["portinfo"] = _portscan
	}
	return
}
