# Tutupen Botolmu [![Build Status](https://travis-ci.org/tiduronline/tutupBotol.svg?branch=master)](https://travis-ci.org/tiduronline/tutupBotol)

This is just a simple command line to help you to add blocked coin miner domains into your hosts file which probably the domain is used our RAM or other stuff in our device to do mining. 

This tool is just to help you to remove and append list of domains to your hosts'file in your Sistem. So where come from the list? basically, I don't collect the list my own. But I got from adBlock list coin domain. 

[NoCoin adblock list](https://raw.githubusercontent.com/hoshsadiq/adblock-nocoin-list/master/hosts.txt) from [Hoshsadiq](https://github.com/hoshsadiq) repository. 

The I choose that list because it's always update.


## How to use

```bash
$ git clone --single-branch --depth=1 --branch master git@github.com:tiduronline/tutupBotol.git 
$ cd tutupBotol
$ make build
$ sudo ./tutup-botol
```
If command successfully executed. The list will be appended to hosts file. To verify the list, you can open the hosts file and make sure that each coin domain is pointed to IP `0.0.0.0`

## Command Options

Maybe you want to see available options for the command, you can type on your favorite console something like this.

```bash
$ ./tutup-botol -h
Usage of ./tutup-botol:
  -block
        Block domain of coins
  -clean
        Unblock coin domain
  -update
        Update list of coin domain
```

##### Description of command:
- `block` and `update` option is used to save new list
- `clean` option is used to clean the list from hosts file

If the tool's executed, there will be two important comments in the hosts file 
`# BOF D COINS` and `# EOF D COINS`. The tool use that comments as mark. So please do not remove that comments. Tool can not work properly if those comments removed. 

But if those comments incidentally removed. Just add `# BOF D COINS` into hosts file and place it before coin domains list and place `# EOF D COINS` after coin domain list. 

## Thanks to 
- [Hoshsadiq](https://github.com/hoshsadiq)