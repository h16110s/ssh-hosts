package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/olekukonko/tablewriter"
)

const SKIP_LINE = "__SKIP_LINE__"

type Host struct {
	num      int
	host     string
	hostName string
}

func keyValueParser(line []string) (string, string) {
	hasKey := false
	var key string
	var value string

	// 「#」で始まる行はスキップ（コメント）
	if regexp.MustCompile(`^#`).Match([]byte(line[0])) {
		key = SKIP_LINE
		return key, value
	}
	for _, word := range line {
		if word == "" {
			continue
		} else if !hasKey {
			key = word
			hasKey = true
		} else {
			value = word
		}
	}

	// keyに何も入らなければスキップ
	if !hasKey {
		key = SKIP_LINE
	}
	return key, value
}

func main() {
	homeDir, _ := os.UserHomeDir()
	data, _ := os.Open(homeDir + "/.ssh/config")

	var hosts []*Host

	defer data.Close()
	scanner := bufio.NewScanner(data)
	tmpHost := new(Host)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		key, value := keyValueParser(line)

		switch strings.ToUpper(key) {
		case "HOST":
			hosts = append(hosts, tmpHost)
			tmpHost = new(Host)
			tmpHost.host = value
			break
		case "HOSTNAME":
			tmpHost.hostName = value
			break
		case SKIP_LINE:
			break
		default:
			break
		}
	}
	// 最後のデータを追加
	hosts = append(hosts, tmpHost)
	fmt.Println("ssh hosts")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Host", "HostName"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

	// 1個目のデータは空データになってる(はず)
	for _, v := range hosts[1:] {
		row := []string{v.host, v.hostName}
		table.Append(row)
	}
	table.Render() // Send output
}
