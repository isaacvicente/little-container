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

> If it asks for `sudo`

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
