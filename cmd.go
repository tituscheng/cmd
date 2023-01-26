package cmd

import (
	"fmt"
	"os"
)

type Cmd struct {
	Name        string
	Action      func(...string)
	Description string
	Args        []string
}

var cmdMap map[string]Cmd
var cmdTitle string

func Title(title string) {
	cmdTitle = title
}

func Add(name, description string, action func(...string)) {
	if cmdMap == nil {
		cmdMap = make(map[string]Cmd)
	}
	cmdMap[name] = Cmd{
		Name:        name,
		Description: description,
		Action:      action,
	}
}

func GetMap() map[string]Cmd {
	return cmdMap
}

func Run() {
	if len(os.Args) == 1 {
		print()
		return
	}
	cmd := findCommand(os.Args)
	if cmd != nil {
		cmd.Action(cmd.Args...)
	} else {
		print()
		return
	}
}

func print() {
	const (
		LEVEL1 = 1
		LEVEL2 = 2
	)
	var tab = func(level int, name, description string) {
		for i := 0; i < level; i++ {
			fmt.Printf("\t")
		}
		fmt.Printf("%s: %s", name, description)
		fmt.Printf("\n")
	}
	fmt.Println()
	fmt.Print(cmdTitle)
	fmt.Println()
	for key, val := range cmdMap {
		tab(LEVEL1, key, val.Description)
	}
	fmt.Println()
}

func findCommand(list []string) *Cmd {
	if cmdMap == nil {
		return nil
	}
	for idx, arg := range list {
		if cmd, ok := cmdMap[arg]; ok {
			cmd.Args = list[idx+1:]
			return &cmd
		}
	}
	return nil
}
