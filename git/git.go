package git

import "os/exec"

func GetStagedDiff() string {
	cmd := exec.Command("git", "diff", "--staged", "--", ":!package-lock.json", ":!yarn.lock")

	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	return string(output)
}

func CreateCommit(message string) {
	cmd := exec.Command("git", "commit", "-m", message)

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
