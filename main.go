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
	num       int
	host      string
	host_name string
}

func keyValueParser(line []string) (string, string) {
	b_key := false
	var key string
	var value string
	len := len(line)
	// 「#」で始まる行はスキップ（コメント）
	if regexp.MustCompile(`^#`).Match([]byte(line[0])) {
		key = SKIP_LINE
		return key, value
	}
	for i := 0; i < len; i++ {
		if line[i] == "" {
			continue
		} else if !b_key {
			key = line[i]
			b_key = true
		} else {
			value = line[i]
		}
	}

	// keyに何も入らなければスキップ
	if !b_key {
		key = SKIP_LINE
	}
	return key, value
}

func main() {
	homeDir, _ := os.UserHomeDir()
	data, _ := os.Open(homeDir + "/.ssh/config")

	var host_list []*Host

	defer data.Close()
	scanner := bufio.NewScanner(data)
	tmp_host := new(Host)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		key, value := keyValueParser(line)

		switch strings.ToUpper(key) {
		case "HOST":
			host_list = append(host_list, tmp_host)
			tmp_host = new(Host)
			tmp_host.host = value
			break
		case "HOSTNAME":
			tmp_host.host_name = value
			break
		case SKIP_LINE:
			break
		default:
			break
		}
	}
	// 最後のデータを追加
	host_list = append(host_list, tmp_host)
	fmt.Println("ssh hosts")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Host", "HostName"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

	// 1個目のデータは空データになってる(はず)
	for _, v := range host_list[1:] {
		str_list := []string{v.host, v.host_name}
		table.Append(str_list)
	}
	table.Render() // Send output
}
