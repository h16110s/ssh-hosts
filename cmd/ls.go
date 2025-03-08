package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/h16110s/ssh-hosts/model"
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
	// ホームディレクトリを取得
	homeDir, _ := os.UserHomeDir()
	data, err := os.Open(homeDir + CONFIG_PATH)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer data.Close()

	// ホスト情報をパース
	hosts := parseHosts(data)

	fmt.Println("ssh hosts")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Host", "HostName"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

	for _, v := range hosts {
		row := []string{v.Host, v.HostName}
		table.Append(row)
	}
	table.Render() // 出力
}

func parseHosts(data *os.File) []*model.Host {
	scanner := bufio.NewScanner(data)
	hosts := []*model.Host{}
	tmpHost := new(model.Host)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		key, value := keyValueParser(line)

		switch strings.ToUpper(key) {
		case "HOST":
			hosts = append(hosts, tmpHost)
			tmpHost = new(model.Host)
			tmpHost.Host = value
		case SKIP_LINE:
		default:
			tmpHost.Setter(key, value)
		}
	}

	if tmpHost.Host != "" {
		hosts = append(hosts, tmpHost)
	}

	return hosts
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
