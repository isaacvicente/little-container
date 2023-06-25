package main

import (
    "fmt"
    "os"
    "os/exec"
    "syscall"
    "path/filepath"
    "strconv"
    "io/ioutil"
)

// docker         run <image> <cmd> <params>
// go run main.go run         <cmd> <params>

func main() {
    switch os.Args[1] {
    case "run":
	run()
    case "child":
	child()
    default:
	panic("Help!")
    }
}

func run() {
    fmt.Printf("Running %v as PID %d\n", os.Args[2:], os.Getpid())

    cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    cmd.SysProcAttr = &syscall.SysProcAttr {
	Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWUSER | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	Unshareflags: syscall.CLONE_NEWNS,
	UidMappings: []syscall.SysProcIDMap {{
	    ContainerID: 0,
	    HostID: 1000,
	    Size: 1,
	}},

	GidMappings: []syscall.SysProcIDMap {{
	    ContainerID: 0,
	    HostID: 1000,
	    Size: 1,
	}},
    }

    must(cmd.Run())
}

func child() {
    fmt.Printf("Running %v as PID %d\n", os.Args[2:], os.Getpid())

    //cg()

    syscall.Sethostname([]byte("little-ubuntu"))
    must(syscall.Chroot("/home/vagrant/little-container/images/ubuntu"))
    must(syscall.Chdir("/"))
    must(syscall.Mount("proc", "proc", "proc", 0, ""))

    cmd := exec.Command(os.Args[2], os.Args[3:]...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    must(cmd.Run())

    syscall.Unmount("proc", 0)
}

func cg() {
    pids := "/sys/fs/cgroup/pids"
    os.Mkdir(filepath.Join(pids, "zak"), 0755)
    must(ioutil.WriteFile(filepath.Join(pids, "zak/pids.max"), []byte("20"), 0700))
    must(ioutil.WriteFile(filepath.Join(pids, "zak/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func must(err error) {
    if err != nil {
	panic(err)
    }
}

