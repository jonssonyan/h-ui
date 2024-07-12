package service

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"h-ui/dao"
	"h-ui/model/constant"
	"h-ui/util"
	"os/exec"
	"strings"
)

var (
	netManager       string
	ingressInterface string
	Add              = "add"
	Delete           = "delete"
	Comment          = "h-ui"
)

func init() {
	if nft, err := util.Exec("command -v nft"); err == nil && nft != "" {
		netManager = "nft"
	} else if iptables, err := util.Exec("command -v iptables"); err == nil && iptables != "" {
		netManager = "iptables"
	}

	if ii, err := util.Exec("ls /sys/class/net | grep -E '^en'"); err == nil && ii != "" {
		ingressInterface = ii
	}
}

func InitTableAndChain() error {
	if netManager == "nft" {
		_, err := util.Exec("nft add table inet hysteria_porthopping")
		if err != nil {
			return err
		}
		_, err = util.Exec("nft add chain inet hysteria_porthopping prerouting { type nat hook prerouting priority dstnat\\; policy accept\\; }")
		if err != nil {
			return err
		}
	} else if netManager == "iptables" {

	}
	return nil
}

func InitPortHopping() error {
	if err := RemoveByComment(); err != nil {
		return err
	}

	hysteria2Config, err := GetHysteria2Config()
	if err != nil {
		return err
	}

	// set port forward
	hysteria2ConfigPortHopping, err := dao.GetConfig("key = ?", constant.Hysteria2ConfigPortHopping)
	if err != nil {
		return err
	}
	if *hysteria2ConfigPortHopping.Value != "" {
		listen := strings.Split(*hysteria2Config.Listen, ":")
		if len(listen) == 2 {
			if err := PortForward(*hysteria2ConfigPortHopping.Value, listen[1], Add); err != nil {
				return err
			}
		}
	}
	return nil
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
	// nft list ruleset
	// 创建表：nft add table inet hysteria_porthopping
	// 创建链：nft add chain inet hysteria_porthopping prerouting { type nat hook prerouting priority dstnat\; policy accept\; }
	// 添加规则：nft add rule inet hysteria_porthopping prerouting iifname eth0 udp dport {30000-40000} counter redirect to :444 comment h-ui
	_, err := util.Exec(fmt.Sprintf("nft %s rule inet hysteria_porthopping prerouting iifname %s udp dport {%s} counter redirect to :%s comment %s", option, ingressInterface, rules, target, Comment))
	if err != nil {
		return err
	}

	return nil
}

func ntfRemoveByComment(comment string) error {
	if netManager != "nft" {
		return errors.New("nftables not found on the system")
	}

	rules, err := nftRules()
	if err != nil {
		return err
	}
	for _, rule := range rules {
		if strings.Contains(rule, comment) {
			_, err := util.Exec(fmt.Sprintf("nft delete rule inet hysteria_porthopping prerouting %s", strings.TrimSpace(rule)))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func nftRules() ([]string, error) {
	output, err := util.Exec("nft list chain inet hysteria_porthopping prerouting")
	if err != nil {
		return nil, err
	}

	rules := strings.Split(output, "\n")
	return rules, nil
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
	output, err := util.Exec(fmt.Sprintf("%s -t nat -S PREROUTING", protocol))
	if err != nil {
		return nil, err
	}

	rules := strings.Split(output, "\n")
	return rules, nil
}
