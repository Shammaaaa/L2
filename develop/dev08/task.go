package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	for {
		dir, err := os.Getwd()
		if err != nil {
			return
		}
		fmt.Print(dir, ":: ")
		if stdin.Scan() {
			cmd := stdin.Text()
			cmdSlice := strings.Split(cmd, "|")
			GoToExec(cmdSlice)
		}
	}
}

func cmdCd(cmd []string) {
	if len(cmd) != 2 {
		fmt.Fprint(os.Stderr, "please insert a path\n")
	}
	err := os.Chdir(cmd[1])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}

func cmdPwd(cmd []string) {
	if len(cmd) > 1 {
		fmt.Fprint(os.Stderr, "too many arguments\n")
	} else {
		path, err := os.Getwd()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		fmt.Println(path)
	}
}

func cmdEcho(cmd []string) {
	res := make([]string, 0)
	if len(cmd) == 1 {
		fmt.Println()
	} else {
		for i := 1; i < len(cmd); i++ {
			res = append(res, cmd[i])
		}
		fmt.Println(strings.Join(res, " "))
	}
}

func cmdKill(cmd []string) {
	pid, err := strconv.Atoi(cmd[1])
	if err != nil {
		fmt.Fprint(os.Stderr, err, "\n")
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprint(os.Stderr, err, "\n")
	}
	err = proc.Kill()
	if err != nil {
		fmt.Fprint(os.Stderr, err, "\n")
	}
}

func cmdPs() {
	sliceProc, _ := ps.Processes()

	for _, proc := range sliceProc {
		fmt.Printf("Process name: %v process id: %v\n", proc.Executable(), proc.Pid())
	}

}

func anotherCmd(cmd []string) {
	var comm *exec.Cmd
	if len(cmd) > 1 {
		comm = exec.Command(cmd[0], cmd[1:]...)
	} else {
		comm = exec.Command(cmd[0])
	}
	output, err := comm.Output()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	fmt.Println(string(output))
}

func GoToExec(cmd []string) {

	for _, c := range cmd {
		partCmd := strings.Split(c, " ")
		switch partCmd[0] {
		case "cd":
			cmdCd(partCmd)
		case "pwd":
			cmdPwd(partCmd)
		case "echo":
			cmdEcho(partCmd)
		case "kill":
			cmdKill(partCmd)
		case "ps":
			cmdPs()
		case "q":
			fmt.Fprint(os.Stdout, "quit\n")
			os.Exit(0)
		default:
			anotherCmd(partCmd)
		}
	}
}
