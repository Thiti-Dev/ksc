package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

// Execute the shell command via sh, also support both showing stderr and stdout
func ExecuteShellCommand(cmdStr string) {
	// Use /bin/sh to execute the command on Unix-like systems
	cmd := exec.Command("/bin/sh", "-c", cmdStr)

	// Get the command's stdout pipe
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Failed to get stdout pipe: %v", err)
	}

	// Get the command's stderr pipe
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("Failed to get stderr pipe: %v", err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		log.Fatalf("Command start failed: %v", err)
	}

	// Create a scanner for stdout
	stdoutScanner := bufio.NewScanner(stdout)
	go func() {
		for stdoutScanner.Scan() {
			fmt.Println(stdoutScanner.Text())
		}
	}()

	// Create a scanner for stdout
	stderrScanner := bufio.NewScanner(stderr)
	go func() {
		for stderrScanner.Scan() {
			fmt.Println(stderrScanner.Text())
		}
	}()

	// Check for scanner errors
	if err := stdoutScanner.Err(); err != nil {
		log.Fatalf("Error reading command output (stdout): %v", err)
	}
	if err := stderrScanner.Err(); err != nil {
		log.Fatalf("Error reading command output (stderr): %v", err)
	}

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		log.Fatalf("Command execution failed: %v", err)
	}
}
