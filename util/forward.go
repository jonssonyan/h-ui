package util

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

var (
	netManager string

	// Constants for command options
	Ipv4   = "iptables"
	Ipv6   = "ip6tables"
	Add    = "add"
	Delete = "delete"
)

func init() {
	// Check for nftables
	nft, err := exec.Command("command", "-v", "nft").Output()
	if err == nil && len(nft) > 0 {
		netManager = "nft"
		return
	}

	// Check for iptables
	iptables, err := exec.Command("command", "-v", "iptables").Output()
	if err == nil && len(iptables) > 0 {
		netManager = "iptables"
		return
	}

	// If neither found, set to empty
	netManager = ""
}

func PortForward(rules string, target string, option string, protocol string) error {
	switch netManager {
	case "nft":
		switch option {
		case Add, Delete:
			return NftForward(rules, target, option)
		default:
			return errors.New("unsupported command option")
		}
	case "iptables":
		switch option {
		case Add, Delete:
			return IptablesForward(rules, target, option, protocol)
		default:
			return errors.New("unsupported command option")
		}
	default:
		return errors.New("port forwarding not supported on this system")
	}
}
func NftForward(rules string, target string, option string) error {
	if netManager != "nft" {
		return fmt.Errorf("nftables not found on the system")
	}

	rulePairs := strings.Split(rules, ",")
	for _, pair := range rulePairs {
		portRange := strings.Split(pair, "-")
		if len(portRange) != 2 {
			return fmt.Errorf("invalid port range format: %s", pair)
		}
		startPort := strings.TrimSpace(portRange[0])
		endPort := strings.TrimSpace(portRange[1])

		cmd := exec.Command("nft", option, "rule", "nat", "prerouting", "iif", "udp", "dport", fmt.Sprintf("{%s-%s}", startPort, endPort), "redirect", "to", ":"+target)
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("failed to %s nftables rule: %v", option, err)
		}
	}

	return nil
}

func IptablesForward(rules string, target string, option string, protocol string) error {
	if netManager != "iptables" {
		return fmt.Errorf("iptables not found on the system")
	}

	rulePairs := strings.Split(rules, ",")
	for _, pair := range rulePairs {
		portRange := strings.Split(pair, "-")
		if len(portRange) != 2 {
			return fmt.Errorf("invalid port range format: %s", pair)
		}
		startPort := strings.TrimSpace(portRange[0])
		endPort := strings.TrimSpace(portRange[1])

		cmd := exec.Command(protocol, "-t", "nat", option, "PREROUTING", "-p", "udp", "--dport", startPort+":"+endPort, "-j", "REDIRECT", "--to-port", target)
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("failed to %s iptables rule: %v", option, err)
		}
	}

	return nil
}
