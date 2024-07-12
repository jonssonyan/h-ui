package util

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

var (
	netManager string
	Add        = "add"
	Delete     = "delete"
	Comment    = "h-ui"
)

func init() {
	// Check for nftables
	nft, err := Exec("command -v nft")
	if err == nil && len(nft) > 0 {
		netManager = "nft"
		return
	}

	// Check for iptables
	iptables, err := Exec("command -v iptables")
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
			return nftForward(rules, target, option)
		default:
			return errors.New("unsupported command option")
		}
	case "iptables":
		switch option {
		case Add:
			return iptablesForward(rules, target, "-A")
		case Delete:
			return iptablesForward(rules, target, "-D")
		default:
			return errors.New("unsupported command option")
		}
	default:
		return errors.New("port hopping not supported on this system")
	}
}

func RemoveByComment() error {
	switch netManager {
	case "nft":
		return ntfRemoveByComment(Comment)
	case "iptables":
		if err := iptablesRemoveByComment("iptables", Comment); err != nil {
			return err
		}
		if err := iptablesRemoveByComment("ip6tables", Comment); err != nil {
			return err
		}
		return nil
	default:
		return errors.New("port hopping not supported on this system")
	}
}

func nftForward(rules string, target string, option string) error {
	if netManager != "nft" {
		return fmt.Errorf("nftables not found on the system")
	}

	rulePairs := strings.Split(rules, ",")
	for _, pair := range rulePairs {
		ports := ""
		portRange := strings.Split(pair, "-")
		if len(portRange) == 1 {
			ports = strings.TrimSpace(portRange[0])
		} else if len(portRange) == 2 {
			startPort := strings.TrimSpace(portRange[0])
			endPort := strings.TrimSpace(portRange[1])
			ports = fmt.Sprintf("{%s-%s}", startPort, endPort)
		} else {
			return fmt.Errorf("invalid port range format: %s", pair)
		}

		cmd := exec.Command("nft", option, "rule", "ip", "nat", "prerouting", "iif", "eth0", "udp", "dport", ports, "redirect", "to", ":"+target, "comment", Comment)
		if err := cmd.Run(); err != nil {
			errMsg := fmt.Sprintf("failed to %s nftables rule: %v", option, err)
			logrus.Errorf(errMsg)
			return errors.New(errMsg)
		}
	}

	return nil
}

func ntfRemoveByComment(comment string) error {
	if netManager != "nft" {
		return fmt.Errorf("nftables not found on the system")
	}

	cmd := exec.Command("nft", "delete", "ip", "nat", "prerouting", "comment ", comment)
	if err := cmd.Run(); err != nil {
		errMsg := fmt.Sprintf("failed to delete nftables comment: %s err: %v", comment, err)
		logrus.Errorf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func iptablesForward(rules string, target string, option string) error {
	if netManager != "iptables" {
		return fmt.Errorf("iptables not found on the system")
	}

	rulePairs := strings.Split(rules, ",")
	for _, pair := range rulePairs {
		ports := ""
		portRange := strings.Split(pair, "-")
		if len(portRange) == 1 {
			ports = strings.TrimSpace(portRange[0])
		} else if len(portRange) == 2 {
			startPort := strings.TrimSpace(portRange[0])
			endPort := strings.TrimSpace(portRange[1])
			ports = startPort + ":" + endPort
		} else {
			return fmt.Errorf("invalid port range format: %s", pair)
		}

		if len(ports) != 0 {
			if err := iptablesAddRule(option, ports, target); err != nil {
				return err
			}
		}
	}

	return nil
}

func iptablesAddRule(option, ports, target string) error {
	Ipv4Cmd := exec.Command("iptables", "-t", "nat", option, "PREROUTING", "-i", "eth0", "-p", "udp", "--dport", ports, "-j", "REDIRECT", "--to-port", target, "-m", "comment", "--comment", Comment)
	if err := Ipv4Cmd.Run(); err != nil {
		errMsg := fmt.Sprintf("failed to %s iptables rule: %v", option, err)
		logrus.Errorf(errMsg)
		return errors.New(errMsg)
	}
	Ipv6Cmd := exec.Command("ip6tables", "-t", "nat", option, "PREROUTING", "-i", "eth0", "-p", "udp", "--dport", ports, "-j", "REDIRECT", "--to-port", target, "-m", "comment", "--comment", Comment)
	if err := Ipv6Cmd.Run(); err != nil {
		errMsg := fmt.Sprintf("failed to %s ip6tables rule: %v", option, err)
		logrus.Errorf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func iptablesRemoveByComment(protocol, comment string) error {
	rules, err := iptablesRules(protocol)
	if err != nil {
		return err
	}
	for _, rule := range rules {
		if strings.Contains(rule, comment) {
			cmd := exec.Command(protocol, "-t", "nat", "-D", "PREROUTING", rule)
			if err := cmd.Run(); err != nil {
				errMsg := fmt.Sprintf("failed to delete iptables rule: %v", err)
				logrus.Errorf(errMsg)
				return errors.New(errMsg)
			}
		}
	}

	return nil
}

func iptablesRules(protocol string) ([]string, error) {
	output, err := Exec(fmt.Sprintf("%s -t nat -S PREROUTING", protocol))
	if err != nil {
		errMsg := fmt.Sprintf("failed to list rules: %v", err)
		logrus.Errorf(errMsg)
		return nil, errors.New(errMsg)
	}

	rules := strings.Split(output, "\n")
	return rules, nil
}