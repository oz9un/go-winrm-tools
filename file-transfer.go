package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/masterzen/winrm"
	"golang.org/x/crypto/ssh/terminal"
)

func command_runner(target_ip string, target_port int, direction string, file_1 string, file_2 string, username_target string, password_target string, target_dn_computer string) bool {
	endpoint := winrm.NewEndpoint(target_ip, target_port, true, true, nil, nil, nil, 0)
	params := winrm.DefaultParameters

	params.TransportDecorator = func() winrm.Transporter { return &winrm.ClientNTLM{} }

	_, err := winrm.NewClientWithParameters(endpoint, username_target, password_target, params)
	if err != nil {
		panic(err)
	}

	// Store your password securely: (Be care to injections!)
	command_password_store := "$temp_passwd = ConvertTo-SecureString -AsPlainText -Force -String \"" + password_target + "\";"

	// Create a variable for your credentials:
	command_create_credentials := "$temp_cred = New-Object -TypeName System.Management.Automation.PSCredential -ArgumentList \"" + target_dn_computer + "\", $temp_passwd;"

	// Create session:
	command_create_session := "$my_new_session1 = New-PSSession -ComputerName " + target_ip + " -Credential $temp_cred -Authentication Negotiate;"

	// Copy file:
	var command_copy_file string
	if direction == "1" {
		command_copy_file = "Copy-Item " + file_1 + " " + file_2 + " -ToSession $my_new_session1"
	} else {
		command_copy_file = "Copy-Item " + file_1 + " " + file_2 + " -FromSession $my_new_session1"
	}

	complete_command := command_password_store + command_create_credentials + command_create_session + command_copy_file
	cmd := exec.Command("pwsh", "-c", complete_command)

	color.Red("\nAnswer:")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	// Print the output
	fmt.Println(string(stdout))
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
	ip_address := input_getter("Enter the target machine's IP address: ")
	port_no := input_getter("Enter the target port (5986 or 5985): ")
	username_target := input_getter("Enter target machine's username: ")
	password_target := pass_getter("Enter target machine's password: ")
	dn_user_target := input_getter("\nEnter target machine's name (in format: ozg.winservdc.com\\Administrator): ")

	int_port_no, _ := strconv.Atoi(port_no)

	direction := os.Args[1]
	file_1 := os.Args[2]
	file_2 := os.Args[3]

	command_runner(ip_address, int_port_no, direction, file_1, file_2, username_target, password_target, dn_user_target)
}
