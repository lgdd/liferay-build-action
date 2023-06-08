package main

import (
	"os"
	"os/exec"
)

func main() {
	buildTool := getBuildTool()
	buildCmd := getBuildCmd(buildTool)
	runCmd(buildCmd)
}

func runCmd(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}

func getBuildTool() string {
	if _, err := os.Stat("pom.xml"); !os.IsNotExist(err) {
		return "mvn"
	}
	return "gradle"
}

func getBuildCmd(buildTool string) string {
	if buildTool == "mvn" {
		return "package"
	}
	return "build"
}
