package util

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

var (
	netManager string
	Add        = "add"
	Delete     = "delete"
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

func PortForward(rules string, target string, option string) error {
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
		case Add:
			return IptablesForward(rules, target, "-A")
		case Delete:
			return IptablesForward(rules, target, "-D")
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

		cmd := exec.Command("nft", option, "rule", "ip", "nat", "prerouting", "iif", "eth0", "udp", "dport", fmt.Sprintf("{%s-%s}", startPort, endPort), "redirect", "to", ":"+target, "comment", "h-ui")
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("failed to %s nftables rule: %v", option, err)
		}
	}

	return nil
}

func NtfRemoveByComment(comment string) error {
	if netManager != "nft" {
		return fmt.Errorf("nftables not found on the system")
	}
	cmd := exec.Command("nft", "delete", "ip", "nat", "prerouting", "comment ", comment)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to delete nftables comment: %s err: %v", comment, err)
	}
	return nil
}

func IptablesForward(rules string, target string, option string) error {
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

		Ipv4Cmd := exec.Command("iptables", "-t", "nat", option, "PREROUTING", "-i", "eth0", "-p", "udp", "--dport", startPort+":"+endPort, "-j", "REDIRECT", "--to-port", target, "-m", "comment", "--comment", "h-ui")
		if err := Ipv4Cmd.Run(); err != nil {
			return fmt.Errorf("failed to %s iptables rule: %v", option, err)
		}
		Ipv6Cmd := exec.Command("ip6tables", "-t", "nat", option, "PREROUTING", "-i", "eth0", "-p", "udp", "--dport", startPort+":"+endPort, "-j", "REDIRECT", "--to-port", target, "-m", "comment", "--comment", "h-ui")
		if err := Ipv6Cmd.Run(); err != nil {
			return fmt.Errorf("failed to %s ip6tables rule: %v", option, err)
		}
	}

	return nil
}

func IptablesRemoveByComment(protocol, comment string) error {
	rules, err := iptablesRules(protocol)
	if err != nil {
		return err
	}
	for _, rule := range rules {
		if strings.Contains(rule, comment) {
			cmd := exec.Command(protocol, "-t", "nat", "-D", "PREROUTING", rule)
			if err := cmd.Run(); err != nil {
				return fmt.Errorf("failed to delete iptables rule: %v", err)
			}
		}
	}

	return nil
}

func iptablesRules(protocol string) ([]string, error) {
	output, err := exec.Command(protocol, "-t", "nat", "-S", "PREROUTING").Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list rules: %v", err)
	}

	rules := strings.Split(string(output), "\n")
	return rules, nil
}
