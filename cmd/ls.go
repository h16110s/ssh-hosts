package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/h16110s/ssh-host-manager/model"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lsCmd)
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Print ssh hosts with table view",
	Long:  "Print ssh hosts with table view",
	Run:   ls,
}

const SKIP_LINE = "__SKIP_LINE__"

func ls(cmd *cobra.Command, args []string) {
	homeDir, _ := os.UserHomeDir()
	data, _ := os.Open(homeDir + CONFIG_PATH)

	var hosts []*model.Host

	defer data.Close()
	scanner := bufio.NewScanner(data)
	tmpHost := new(model.Host)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		key, value := keyValueParser(line)

		switch strings.ToUpper(key) {
		case "HOST":
			hosts = append(hosts, tmpHost)
			tmpHost = new(model.Host)
			tmpHost.Host = value
		case "HOSTNAME":
			tmpHost.HostName = value
		case SKIP_LINE:
		default:
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
		row := []string{v.Host, v.HostName}
		table.Append(row)
	}
	table.Render() // Send output
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
