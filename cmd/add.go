package cmd

import (
	"fmt"
	"os"

	"github.com/h16110s/ssh-hosts/model"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new SSH host",
	Long:  `Add a new SSH host to the SSH config file`,
	Run:   addHost,
}

func addHost(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: ssh-hosts add <Host> <HostName> [<Property>=<Value>...]")
		return
	}

	host := args[0]
	hostName := args[1]

	h := model.Host{
		Host:     host,
		HostName: hostName,
	}

	// for _, arg := range args[2:] {
	// 	parts := strings.SplitN(arg, "=", 2)
	// 	if len(parts) != 2 {
	// 		fmt.Printf("Invalid property format: %s\n", arg)
	// 		continue
	// 	}
	// 	h.Setter(parts[0], parts[1])
	// }

	homeDir, _ := os.UserHomeDir()
	configPath := homeDir + CONFIG_PATH

	file, err := os.OpenFile(configPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Error opening SSH config file:", err)
		return
	}
	defer file.Close()

	newHost := fmt.Sprintf("\nHost %s\n\tHostName %s\n", h.Host, h.HostName)

	// for _, prop := range h.GetFieldNames() {
	// 	value := reflect.ValueOf(&h).Elem().FieldByName(prop).String()
	// 	if value != "" {
	// 		newHost += fmt.Sprintf("\t%s %s\n", prop, value)
	// 	}
	// }

	if _, err := file.WriteString(newHost); err != nil {
		fmt.Println("Error writing to SSH config file:", err)
	}

	fmt.Printf("Host %s added to SSH config file\n", h.Host)
}
