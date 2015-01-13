## eXtended docker

Nicolas MULLER, 2015

Extended CLI for docker because typing long docker commands is boring...

Project built with Cobra, a commander for modern Go CLI interactions [https://github.com/spf13/cobra].

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

* **kill [all|rex]** - Kill all containers or these named and provided by regular expression.
  The option ``all`` required an confirmation else you must force it.

**Note:** Nothing to say...
