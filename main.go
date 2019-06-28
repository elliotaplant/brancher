package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

type Issue struct {
	id   int64
	name string
}

func (i Issue) String() string {
	return fmt.Sprintf("#%d  %s", i.id, i.name)
}

func (i Issue) mainBranch() string {
	titleCaseName := strings.Title(i.name)
	joinedName := strings.Join(strings.Fields(titleCaseName), "")
	return fmt.Sprintf("%s#%d", joinedName, i.id)
}

func (i Issue) personalBranch() string {
	initials := os.Args[1]
	return fmt.Sprintf("%s_%s", initials, i.mainBranch())
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Requires your initials (eg. 'eap') as the first and only argument")
		return
	}
	cmd := exec.Command("hub", "issue")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		fmt.Print(err)
		return
	}
	splitOutput := strings.Split(out.String(), "\n")
	issues := []Issue{}

	for _, line := range splitOutput {
		if line != "" {
			line := strings.TrimSpace(line)
			idStart := 1
			idEnd := idStart
			for line[idEnd] != ' ' {
				idEnd++
			}
			id, err := strconv.ParseInt(line[idStart:idEnd], 10, 64)
			if err != nil {
				panic(err)
			}
			issues = append(issues, Issue{
				id:   id,
				name: line[idEnd+2:],
			})
		}
	}

	prompt := promptui.Select{
		Label: "Select Issue",
		Items: issues,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	selectedBranch := issues[i]

	err = makeBranch(selectedBranch.mainBranch())
	if err != nil {
		fmt.Printf("Error creating branch %s: %s\nIt may already exist", selectedBranch.mainBranch(), err)
		return
	}

	err = makeBranch(selectedBranch.personalBranch())
	if err != nil {
		fmt.Printf("Error creating branch %s: %s\nIt may already exist", selectedBranch.personalBranch(), err)
		return
	}
}

func makeBranch(branchName string) error {
	cmd := exec.Command("git", "checkout", "-b", branchName)
	var out bytes.Buffer
	cmd.Stdout = &out
	return cmd.Run()
}
