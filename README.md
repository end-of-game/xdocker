## eXtended docker

Nicolas MULLER, 2015

Extended CLI for docker because typing long docker commands is boring...

## SETUP

```
cd ~
mkdir bin
cd bin
git clone git@github.com:Treeptik/xdocker.git
```
Now add this somewhere to your ``~/.bashrc``:

```
export PATH=$PATH:$HOME/bin/xdocker
```

**Warning**: Some of these commands may be destructive - use at your own risk
and test on a test system before using in a production environment.

## Command quick reference

* **killall** - Kill all named containers. Provide a prefix for the image name.
  The committed image will be in the form ``<namespace>/<container name>``.
* **dcalldated** - Like the above command but adds a date stamp as the tag.
  The committed image will be in the form ``<namespace>/<container name>:01-Septmber-2014``.


**Note:** Docker namespaces must be at least 3 characters long and less that 30
