# little container - written in Go

This is a tiny container based in Docker. It has just the "run" command
implemented, without the "image' implementation. As of now, only Ubuntu works
inside the container.

The code was taken from [Liz Rice talk](https://youtu.be/oSlheqvaRso).

> Warning: It only works on Ubuntu 22.04 LTS Jammy Jellyfish.

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

You're now in a Ubuntu VM. There's still a few things you neeed to do.

Still in the VM, run:

```
$ sudo apt install qemu-user-static
$ sudo systemctl restart systemd-binfmt.service
```

## Run

Switch to *root* user with `sudo su -` and then:

```
# go run main.go run /bin/bash
```

Now you're inside the container!
