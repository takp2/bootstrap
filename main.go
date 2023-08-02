package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	Version    = "0.0.0"
	subtle     = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	titleStyle = lipgloss.NewStyle().
			Align(lipgloss.Center).
			Foreground(lipgloss.Color("#43BF6D")). // dark green #43BF6D
			Background(subtle)
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(lipgloss.NewStyle().SetString("❌").
			Foreground(lipgloss.AdaptiveColor{Light: "#FF5555", Dark: "#FF5555"}).
			PaddingRight(1).
			String(), "Bootstrap failed:", err)
		os.Exit(1)
	}
}

func run() error {
	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))

	titleWidth := 77
	if physicalWidth < titleWidth {
		titleWidth = physicalWidth
	}
	fmt.Println(titleStyle.Width(titleWidth).Render("Bootstrap v" + Version))
	type fileEntry struct {
		name string
		path string
	}
	files := []fileEntry{
		{name: "manager", path: "manager"},
		{name: "config.yaml", path: "config.yaml"},
		{name: "zone", path: "zone"},
		{name: "world", path: "world"},
		{name: "ucs", path: "ucs"},
		{name: "loginserver", path: "loginserver"},
		{name: "queryserv", path: "queryserv"},
	}
	for _, file := range files {
		fi, err := os.Stat(file.path)
		if err != nil {
			return fmt.Errorf("%s not found", file.name)
		}
		if fi.IsDir() {
			return fmt.Errorf("%s is a directory", file.name)
		}
	}

	fmt.Println(lipgloss.NewStyle().SetString("✅").
		Foreground(lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}).
		PaddingRight(1).
		String(), "Success")

	cmd := exec.Command("./manager")
	cmd.Dir = "."
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("manager run: %w", err)
	}
	return nil
}
