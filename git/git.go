package git

import (
	"fmt"
	"os"
	"os/exec"
)

func GetStagedDiff() string {
	excludedFiles := []string{
		":!package-lock.json",
		":!yarn.lock",
		":!go.sum",
		":!Pipfile.lock",
		":!Gemfile.lock",
		":!pom.xml",
		":!build.gradle",
		":!composer.lock",
		":!Cargo.lock",
		":!Package.resolved",
		":!stack.yaml.lock",
	}

	args := append([]string{"git", "diff", "--staged", "--"}, excludedFiles...)
	cmd := exec.Command(args[0], args[1:]...)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running git diff:", err)
		panic(err)
	}

	return string(output)
}

func GetShortStatus() string {
	cmd := exec.Command("git", "status", "--short")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running git status:", err)
		panic(err)
	}

	return string(output)
}

func CreateCommit(commitMessage string) {
	cmd := exec.Command("git", "commit", "-m", commitMessage)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		os.Exit(1)
	}
	fmt.Println("Commit successful!")
}
