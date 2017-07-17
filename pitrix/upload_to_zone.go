package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"os"
	// "bufio"
	// "log"
	// "syscall"
)

type SSHCommander struct {
	User string
	IP   string
	Port int
}

const ()

func exec_local(cmd []string) {

}

func (s *SSHCommander) Command() *exec.Cmd {
	_ = strconv.Itoa(s.Port)
	arg := []string{
		// "/usr/bin/rsync",
		// "-azP",
		// "-e",
		// "'ssh -p " + strconv.Itoa(s.Port) + "'",
		fmt.Sprintf("%s@%s:%s", s.User, s.IP, "/root/"),
	}
	fmt.Println(arg)
	return exec.Command("rsync", arg...)

	// cmd := []string{
	// 	// "apt-get",
	// 	// "update",

	// 	// "ls",
	// 	// "-l",
	// }

	// arg := append(
	// 	[]string{
	// 		fmt.Sprintf("%s@%s", s.User, s.IP),
	// 	},
	// 	cmd...,
	// )
	// return exec.Command("ssh", arg...)
}

func main() {
	commanders := []SSHCommander{
		{"root", "192.168.7.65", 22},
		{"root", "192.168.7.4", 22},
	}
	var (
		// cmdOut []byte
		// err    error
	)
	// rsync -azP -e "ssh -p port" user@ip:/src /dst
	for i := 0; i < len(commanders); i++ {

		// if cmd, e := exec.Run("/bin/ls", nil, nil, exec.DevNull, exec.Pipe, exec.MergeWithStdout); e == nil {
		// 	b, _ := ioutil.ReadAll(cmd.Stdout)
		// 	println("output: " + string(b))
		// }

		cmd := commanders[i].Command()

		if out, err := cmd.CombinedOutput(); err != nil {
			fmt.Fprintln(os.Stderr, "There was an error running SSH command: ", err)
		} else {
			fmt.Printf("The output is %s\n", out)
		}

		// cmd.Stdout = os.Stdout
		// cmd.Stderr = os.Stderr
		// cmdReader, err := cmd.StdoutPipe()
		// if err != nil {
		// 	fmt.Fprintln(os.Stderr, "There was an error running SSH command: ", err)
		// } else {
		// 	// fmt.Println(string(cmdOut))
		// }

		// scanner := bufio.NewScanner(cmdReader)
		// go func() {
		// 	for scanner.Scan() {
		// 		fmt.Printf("docker build out | %s\n", scanner.Text())
		// 	}
		// }()

		// if err := cmd.Start(); err != nil {
		// 	log.Fatalf("cmd.Start: %v")
		// 	// fmt.Fprintln(os.Stderr, "There was an error running SSH command: ", err)
		// 	// os.Exit(1)
		// }
		// if err := cmd.Wait(); err != nil {
		// 	if exiterr, ok := err.(*exec.ExitError); ok {
		// 		// The program has exited with an exit code != 0
		// 		// This works on both Unix and Windows. Although package
		// 		// syscall is generally platform dependent, WaitStatus is
		// 		// defined for both Unix and Windows and in both cases has
		// 		// an ExitStatus() method with the same signature.
		// 		if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
		// 			log.Printf("Exit Status: %d", status.ExitStatus())
		// 		}
		// 	} else {
		// 		log.Fatalf("cmd.Wait: %v", err)
		// 	}
		// 	// fmt.Println(string(cmdOut))
		// }
	}
}
