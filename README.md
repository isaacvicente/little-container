# 📦 Little Container

This is a tiny container based in Docker. It uses a Ubuntu 22.04 Jammy Jellyfish
Base image.

The code was taken from Liz Rice talks: [1](https://youtu.be/oSlheqvaRso) and
[2](https://youtu.be/jeTKgAEyhsA).

> **Warning**
> It only works on Ubuntu 22.04 LTS Jammy Jellyfish.

## Dependencies

If you're not on Ubuntu 22.04 LTS, make sure to have [`vagrant`](https://github.com/hashicorp/vagrant)
installed on your machine as well as [Virtualbox](https://www.virtualbox.org/).

> I suggest you use the VM anyway.

Run:
```
$ vagrant init ubuntu/jammy64
$ vagrant up
$ vagrant ssh
```

You're now in a Ubuntu VM. There's still a few things you need to do.

Still in the VM, run:

```
$ sudo apt install qemu-user-static
$ sudo systemctl restart systemd-binfmt.service
```

## Setup

Clone the repo:

```
$ git clone https://github.com/isaacvicente/little-container.git
```

In order to have the Ubuntu Base filesystem, run the `setup` script:

```bash
$ cd little-conteiner && ./setup
```

## Run

Run the little container in rootless mode by:

```
$ go run main.go run bash
```

Now you're inside the container!

> **Note**
> Only the `run` command was implemented.
> Also, you cannot use the `apt` command to install or
> update packages inside the container, as the base
> system doesn't have the packages to sign GPG keys.

## What I've learned
Nowadays containers are everywhere. Everyone is talking about Docker wonders.
But how do containers works, underneath the hood?

A container is just a process. This is why containers are lighter than virtual
machines, as the containers uses the host's resources and kernel as opposed to
VMs, which have their own kernel and resources (of course, it's not *their*
resources, but it's a emulation of a real machine, so you have to give these
resources to the VM right away, or kinda of).

Containers are built from existing Linux kernel features, namely being:
namespaces, chroot and control groups (cgroups).

### Namespaces
Namespaces delimiters what a process can see. By default, a regular process can
see what's happening around it. But that's not what we want for containers,
right?

There's a few namespaces:
- **PID**: The PID namespace provides processes with an independent set of process IDs (PIDs) from other namespaces. The PID namespace makes the first process created within it assigned with PID 1.
- **MNT**: Mount namespaces control mount points, and provide you to mount and
  unmount folders without affecting other namespaces.
- **NET:** Network namespaces create their network stack for the process.
- **UTS:** UNIX Time-Sharing namespaces allow a process has a separate hostname
  and domain name.
- **USER:** User namespaces create their own set of UIDS and GIDS for the process.
- **IPC:** IPC namespaces isolate processes from inter-process communication,
  this prevents processes in different IPC namespaces from using.

### Chroot

A **`chroot`** on [Unix](https://en.wikipedia.org/wiki/Unix "Unix") and
[Unix-like](https://en.wikipedia.org/wiki/Unix-like "Unix-like") [operating
systems](https://en.wikipedia.org/wiki/Operating_system "Operating system") is
an operation that changes the apparent [root
directory](https://en.wikipedia.org/wiki/Root_directory "Root directory") for
the current running process and its
[children](https://en.wikipedia.org/wiki/Child_process "Child process"). A
program that is run in such a modified environment cannot name (and therefore
normally cannot access) files outside the designated directory tree.

### Cgroups (control groups)

Control groups, usually referred to as cgroups, are a Linux kernel feature
which allow processes to be organized into hierarchical groups whose usage of
various types of resources can then be limited and monitored. That way, a
process can have a limited I/O rate, or a limited number of children processes.

### What was implemented

- Namespaces
    - [x] PID
    - [x] MNT
    - [ ] NET
    - [x] UTS
    - [x] USER
    - [ ] IPC
- [x] Chroot
- [ ] Cgroups

> The code in Liz's talks didn't work for Cgroups, and I couldn't figure out
> how to make it work.

### Conclusion

This was a brief overview about containers. With all these building blocks you
can start to have a container. They are a little more complicated than shown.
Anyway, this is for educational sake, so I don't need to be really technical.

## License

Everything here is licensed under [GPLv3
license](https://www.gnu.org/licenses/gpl-3.0.en.html).
