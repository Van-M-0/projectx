package util

import (
	"runtime"
	"os"
	"path/filepath"
	"fmt"
	"log"
	"bytes"
	"io"
	"strings"
	"encoding/binary"
)

func CreateProcess(dir, cmd string, args []string, expect string) {
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatalf("Pipe: %v", err)
	}
	defer r.Close()
	attr := &os.ProcAttr{Dir: dir, Files: []*os.File{nil, w, os.Stderr}}
	p, err := os.StartProcess(cmd, args, attr)
	if err != nil {
		log.Fatalf("StartProcess: %v", err)
	}
	w.Close()

	var b bytes.Buffer
	io.Copy(&b, r)
	output := b.String()
	fmt.Println("output : ", output)
	fi1, _ := os.Stat(strings.TrimSpace(output))
	fi2, _ := os.Stat(expect)
	if !os.SameFile(fi1, fi2) {
		log.Fatalf("exec %q returned %q wanted %q",
			strings.Join(append([]string{cmd}, args...), " "), output, expect)
	}
	p.Wait()
}

func my_test_create_process() {
	var dir, cmd string
	var args []string
	switch runtime.GOOS {

	case "windows":
		cmd = os.Getenv("COMSPEC")
		dir = os.Getenv("SystemRoot")
		args = []string{"/c", "notepad"}
	default:
		cmd = "/bin/pwd"
		dir = "/"
		args = []string{}
	}
	cmddir, cmdbase := filepath.Split(cmd)
	args = append([]string{cmdbase}, args...)
	fmt.Println(dir, cmd, args, cmddir, cmdbase)
	// Test absolute executable path.
	CreateProcess(dir, cmd, args, dir)
	// Test relative executable path.
	CreateProcess(cmddir, cmdbase, args, cmddir)
}

func My_test_read_buf() {
	type A struct {
		One uint32
	}

	type B struct {
		A
		len uint32
	}

	var b B
	b.One = 1
	b.len = 10

	buf := new(bytes.Buffer)
	fmt.Println("b size is ", binary.Size(b))
	binary.Write(buf, binary.LittleEndian, b)
	fmt.Println("after write, buf is : ", buf.Bytes())
}