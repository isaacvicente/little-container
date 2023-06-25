# 📦 Little Container

This is a tiny container based in Docker. It has just the "run" command
implemented, without the "image' implementation. As of now, only Ubuntu works
inside the container.

The code was taken from Liz Rice talks: [1](https://youtu.be/oSlheqvaRso) and [2](https://youtu.be/jeTKgAEyhsA).

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

Run the `setup` script:

```bash
$ cd little-conteiner && ./setup
``

## Run

Run the little container in rootless mode by:

```
$ go run main.go run bash
```

Now you're inside the container!

> **Note**
> Only the `run` command was implemented.
