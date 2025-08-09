package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stephenZ22/mini_dash/pkg/config"
)

var configFile string

func printBanner() {
	banner := "" +
		"            _       _     _           _                                        \n" +
		"           (_)     (_)   | |         | |                                       \n" +
		"  _ __ ___  _ _ __  _  __| | __ _ ___| |__ ______ ___  ___ _ ____   _____ _ __ \n" +
		" | '_ ` + \"`\" + ` _ \\| | '_ \\| |/ _` |/ _` / __| '_ \\______/ __|/ _ \\ '__\\ \\ / / _ \\ '__|\n" +
		" | | | | | | | | | | | (_| | (_| \\__ \\ | | |     \\__ \\  __/ |   \\ V /  __/ |   \n" +
		" |_| |_| |_|_|_| |_|_|\\__,_|\\__,_|___/_| |_|     |___/\\___|_|    \\_/ \\___|_|   \n" + "\n\n\n"

	fmt.Print(banner)
}

func main() {
	// This is the entry point for the server application.
	// The actual implementation will be added later.
	printBanner()
	fmt.Println("welcome to the mini-dash server application")

	rootCmd := &cobra.Command{
		Use:   "minidash-server",
		Short: "Mini-Dash-Server - Lightweight project management backend",
		Long: `Mini-Dash-Server is a lightweight project management backend 
designed for small teams and studios, offering a clean and efficient way 
to manage projects and tasks.

Features:
  - Multi-project support: Create, edit, and archive multiple projects
  - Task collaboration: Assign tasks, update statuses, and add comments
  - Lightweight deployment: Single binary executable, supports Docker deployment
  - RESTful API: Easily integrate with Web, CLI, or third-party applications
  - Optional features: Structured logging, database migrations, user access control

Example usage:
  # Start the server
  minidash-server start --url http://0.0.0.0 --port 8080 --loglevel info

  # Run database migrations
  minidash-server migrate up

  # Check server status
  minidash-server status
`,
		// PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// 	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		// 		return fmt.Errorf("config file does not exist: %s", configFile)
		// 	}

		// 	config.LoadConfig(configFile)
		// 	fmt.Printf("Loaded configuration from: %s\n", configFile)
		// 	return nil
		// },
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("No command specified. Use 'start' to run the server.")
		// },
	}

	rootCmd.AddCommand(startCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)

		os.Exit(1)
	}
}

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start the Mini-Dash server",
		Long:  "Start the Mini-Dash server with the specified config file and http listen port.",
		RunE:  start_func,
	}

	cmd.Flags().StringVarP(
		&configFile,
		"config",
		"c",
		"configs/app.yaml",
		"Path to config file",
	)

	return cmd
}

func start_func(cmd *cobra.Command, args []string) error {
	if configFile == "" {
		return fmt.Errorf("config file must be specified with --config")
	}

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return fmt.Errorf("config file does not exist: %s", configFile)
	}

	config.LoadConfig(configFile)
	fmt.Printf("Loaded configuration from: %s\n", configFile)

	// Here you would start the server using the loaded configuration.
	fmt.Println("Starting Mini-Dash server...")

	fmt.Printf("Server will run on port: %d, database port: %d\n", config.Cfg.Server.Port, config.Cfg.Database.Port)

	// TODO: Initialize database connection with gorm
	// TODO：add Run MiniDashApp http server
	return nil
}
