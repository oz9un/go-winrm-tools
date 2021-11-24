package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/masterzen/winrm"
	"golang.org/x/crypto/ssh/terminal"
)

func command_runner(slaveIP string, slavePort int, command string, username string, password string) bool {
	endpoint := winrm.NewEndpoint(slaveIP, slavePort, true, true, nil, nil, nil, 0)
	params := winrm.DefaultParameters

	params.TransportDecorator = func() winrm.Transporter { return &winrm.ClientNTLM{} }

	client, err := winrm.NewClientWithParameters(endpoint, username, password, params)
	if err != nil {
		panic(err)
	}

	stdout, _, _, err := client.RunPSWithString(command, "")
	if strings.TrimSpace(stdout) == "" {
		fmt.Println(err)
		fmt.Println(stdout)
	}

	color.Red("\nAnswer:")
	fmt.Println(stdout)
	return true
}

func input_getter(input_message string) string {
	reader := bufio.NewReader(os.Stdin)
	input_color := color.New(color.FgGreen).Add(color.Bold)

	input_color.Print(input_message)
	user_input, user_input_err := reader.ReadString('\n')
	if user_input_err != nil {
		fmt.Println("Error: ", user_input_err)
		return ""
	}

	user_input = strings.TrimRight(user_input, "\n")
	return user_input
}

func pass_getter(input_message string) string {
	input_color := color.New(color.FgGreen).Add(color.Bold)
	input_color.Print(input_message)

	user_pass_byte, user_pass_err := terminal.ReadPassword(0)
	if user_pass_err != nil {
		fmt.Println("Error: ", user_pass_err)
		return ""
	}

	user_pass := string(user_pass_byte)
	return user_pass
}

func main() {
	ip_address := input_getter("Enter the target IP address: ")
	port_no := input_getter("Enter the target port (5986 or 5985): ")
	username := input_getter("Enter your username: ")
	password := pass_getter("Enter your password: ")
	int_port_no, _ := strconv.Atoi(port_no)
	command := input_getter("\nCommand to be executed: ")

	command_runner(ip_address, int_port_no, command, username, password)
}
