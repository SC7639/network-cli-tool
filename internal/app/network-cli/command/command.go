package command

import (
	"fmt"
	"net"
	"strings"
)

// NS gets name servers for a host
func NS(host string) (string, error) {
	ns, err := net.LookupNS(host)
	if err != nil {
		return "", err
	}

	var nss strings.Builder
	for i := 0; i < len(ns); i++ {
		nss.WriteString(fmt.Sprintln(ns[i].Host))
	}

	return nss.String(), nil

}

// IP gets an ip address for a host
func IP(host string) (string, error) {
	ip, err := net.LookupIP(host)
	if err != nil {
		return "", err
	}

	var ips strings.Builder
	for i := 0; i < len(ip); i++ {
		ips.WriteString(fmt.Sprintf("%v\n", ip[i]))
	}

	return ips.String(), nil
}

// CNAME gets the Canonical name record for a host
func CNAME(host string) (string, error) {
	cname, err := net.LookupCNAME(host)
	if err != nil {
		return "", err
	}

	return cname, nil
}

// MX gets the mx record for a host
func MX(host string) (string, error) {
	mx, err := net.LookupMX(host)
	if err != nil {
		return "", err
	}

	var mxs strings.Builder
	for i := 0; i < len(mx); i++ {
		fmt.Println(mx[i].Host, mx[i].Pref)
		mxs.WriteString(fmt.Sprintln(mx[i].Host, mx[i].Pref))
	}

	return mxs.String(), nil
}
