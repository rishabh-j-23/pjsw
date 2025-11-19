package commands

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"pjsw/db"
	"strings"
)

func AddSSH(args []string, database *db.Projects) {
	if len(args) < 3 {
		fmt.Println("Usage: pjsw add-ssh <name> <user@host> [-p port] [-i identity_file]")
		return
	}

	name := args[0]
	connectionStr := args[1]
	
	parts := strings.Split(connectionStr, "@")
	if len(parts) != 2 {
		fmt.Println("Invalid connection string. Format: user@host")
		return
	}
	user := parts[0]
	host := parts[1]

	port := "22"
	identityFile := ""

	for i := 2; i < len(args); i++ {
		switch args[i] {
		case "-p":
			if i+1 < len(args) {
				port = args[i+1]
				i++
			}
		case "-i":
			if i+1 < len(args) {
				identityFile = args[i+1]
				i++
			}
		}
	}

	conn := db.SSHConnection{
		Name:         name,
		User:         user,
		Host:         host,
		Port:         port,
		IdentityFile: identityFile,
	}

	_, err := database.AddSSHConnection(conn)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("SSH connection '%s' added successfully.\n", name)
}

func ConnectSSH(name string, database *db.Projects) {
	conn, err := database.GetSSHConnection(name)
	if err != nil {
		log.Fatalln(err)
	}

	args := []string{"-p", conn.Port}
	if conn.IdentityFile != "" {
		args = append(args, "-i", conn.IdentityFile)
	}
	args = append(args, fmt.Sprintf("%s@%s", conn.User, conn.Host))

	cmd := exec.Command("ssh", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		// Don't log fatal here, as ssh exit code might be non-zero
		fmt.Println("SSH session ended:", err)
	}
}

func ListSSH(database *db.Projects) {
	conns, err := database.GetAllSSHConnections()
	if err != nil {
		log.Fatalln(err)
	}

	for _, conn := range conns {
		fmt.Printf("%s -> %s@%s:%s\n", conn.Name, conn.User, conn.Host, conn.Port)
	}
}

func InteractiveSelectSSH(database *db.Projects) {
	conns, err := database.GetAllSSHConnections()
	if err != nil {
		log.Fatalln(err)
	}

	if len(conns) == 0 {
		fmt.Println("No SSH connections found.")
		return
	}

	var keyBuffer bytes.Buffer
	for _, conn := range conns {
		// Format: name (user@host:port)
		line := fmt.Sprintf("%s (%s@%s:%s)\n", conn.Name, conn.User, conn.Host, conn.Port)
		keyBuffer.WriteString(line)
	}

	// Set up fzf command
	cmd := exec.Command("fzf", "--no-sort", "--reverse", "--height=50%")
	cmd.Stdin = bytes.NewReader(keyBuffer.Bytes())
	cmd.Stderr = os.Stderr // Forward fzf errors to terminal

	output, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 130 {
			return // User cancelled
		}
		fmt.Println("fzf error:", err)
		return
	}

	selectedLine := strings.TrimSpace(string(output))
	if selectedLine == "" {
		return
	}

	// Extract name (everything before the first space)
	parts := strings.Fields(selectedLine)
	if len(parts) > 0 {
		ConnectSSH(parts[0], database)
	}
}
