## eXtended docker

Nicolas MULLER, 2015

Extended CLI for docker because typing long docker commands is boring...

Project built with :

* [Cobra](https://github.com/spf13/cobra), a commander for modern Go CLI interactions .
* [Viper](https://github.com/spf13/viper), Go configuration with fangs, companion to cobra
* [DockerClient](https://github.com/samalba/dockerclient), a Docker client library in Go

Thanks to all developpers about these libraries !

## INSTALL FROM SOURCE

You must know at least how to compile Go sources and have the golang runtime with you.
```
cd ~
mkdir bin
cd bin
git clone git@github.com:Treeptik/xdocker.git
cd xdocker
go build xdocker.go
```
Now add this somewhere to your ``~/.bashrc``:

```
export PATH=$PATH:$GOPATH/src/xdocker/xdocker
```

**Warning**: Some of these commands may be destructive - use at your own risk
and test on a test system before using in a production environment.

## INSTALL WITH DOCKER

Why install xDocker with docker ? Because it’s fun ! Therefore, you won’t mind if my method to build nsenter uses Docker itself.

If you want to install `xdocker` into `/usr/local/bin`, just do this:

    docker run --rm -v /usr/local/bin:/target treeptik/xdocker

The `treeptik/xdocker` container will detect that `/target` is a
mountpoint, and it will copy the `xdocker` binary into it.

If you don’t trust me, and prefer to extract the `xdocker` binary,
rather than allowing my container to potentially wreak havoc into
your system’s `$PATH`, you can also do this:

    docker run --rm treeptik/xdocker cat /xdocker > /tmp/xdocker && chmod +x /tmp/xdocker

Then do whatever you want with the binary in `/tmp/xdocker`.

## Command quick reference

* **kill [all|rex]** - Kill all containers or these named and provided by regular expression.
  The option ``all`` required an confirmation else you must force it.

**Note:** Nothing to say for the moment
