package app

import (
	"bytes"
	"fmt"
	"os/exec"
)

type BashCommand struct {
	Path    string `json:"path" xml:"path" form:"path" query:"path"`
	Command string `json:"command" xml:"command" form:"command" query:"command"`
}

func NewBashCommand(command string) BashCommand {
	return BashCommand{Path: "/home/vladislav", Command: command}
}

func (p *BashCommand) ToString() string {
	/**
	* Формирует строку для выполнения
	 */
	strings := []string{"cd ", p.Path, " && ", p.Command}
	buffer := bytes.Buffer{}
	for _, val := range strings {
		buffer.WriteString(val)
	}
	return buffer.String()
}

func (p *BashCommand) RunBashCommand() string {
	/**
	* Запускает баш команду. Возвращает output команды
	 */
	cmd := exec.Command(p.Command)
	cmd.Dir = p.Path
	// out, err := exec.Command(bashCommand).Output()
	out, err := cmd.Output()
	if err != nil {
		db.Create(&BashHistory{Command: p.Command, Path: p.Path, Output: fmt.Sprintf("%v", err), Status: 1})
		return fmt.Sprintf("%v", err)
	}
	db.Create(&BashHistory{Command: p.Command, Path: p.Path, Output: string(out), Status: 0})
	return fmt.Sprintf("%v", string(out))
}

func GetBashHistory() []map[string]string {
	var m []map[string]string
	var history []BashHistory
	db.Find(&history)
	for i, v := range history {
		m = append(m, make(map[string]string))
		m[i]["command"] = fmt.Sprintf("%v", v.Command)
		m[i]["path"] = fmt.Sprintf("%v", v.Path)
		m[i]["id"] = fmt.Sprintf("%v", v.ID)
		m[i]["output"] = fmt.Sprintf("%v", v.Output)
		m[i]["status"] = fmt.Sprintf("%v", v.Status)
		m[i]["created"] = fmt.Sprintf("%v", v.CreatedAt)
	}
	return m
}
