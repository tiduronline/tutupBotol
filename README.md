# Tutupen Botolmu [![Build Status](https://travis-ci.org/tiduronline/tutupBotol.svg?branch=master)](https://travis-ci.org/tiduronline/tutupBotol)

This is just a simple command line to help you to add blocked coin miner domains into your hosts file which probably the domain is used our RAM or other stuff in our device to do mining. 

This tool is just to help us easily remove and append domain list into our hosts file in our sistem. But where come from that list? 
Actually, I don't collect the list my own. But I got from Hoshsadiq repository for that list coin domain. 

[NoCoin adblock list](https://raw.githubusercontent.com/hoshsadiq/adblock-nocoin-list/master/hosts.txt) from [Hoshsadiq](https://github.com/hoshsadiq) repository. 

The I choose that list because it's always update.

Basically, the tool will modify hosts file which is pretty serious file for security reason. Try it at your own risk. **Backup your hosts file before execute the command**.

Tested on:
- OSX
- Windows
- Ubuntu


## Binary Version

- Version v0.1
    - [Windows](/binary/tutup-botol.win.v0.1.tar) ([Checksum](/binary/tutup-botol.win.sha256))
    - [Ubuntu](/binary/tutup-botol.ubuntu.v0.1.tar) ([Checksum](/binary/tutup-botol.ubuntu.sha256))
    - [MacOS](/binary/tutup-botol.osx.v0.1.tar) ([Checksum](/binary/tutup-botol.osx.sha256))


## How to use

```bash
$ git clone git@github.com:tiduronline/tutupBotol.git 
$ cd tutupBotol
$ make install
$ sudo ./bin/tutup-botol --block
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

If the tool executed, there's will two important comments in the hosts file 
`# BOF D COINS` and `# EOF D COINS`. Tool use that comments as marker. So please do not remove those comments. Tool can not work properly if those comments removed. 

But if those comments incidentally removed. Just add `# BOF D COINS` into hosts file and place it before coin domains list and place `# EOF D COINS` after coin domain list. 

#### Video
[![asciicast](https://asciinema.org/a/kiljuXTmOcAMamDzWsXjZgPd1.png)](https://asciinema.org/a/kiljuXTmOcAMamDzWsXjZgPd1)


## Thanks to 
- [Hoshsadiq](https://github.com/hoshsadiq)
- [Anggrainini](https://github.com/anggrainini), I borrow her device to test Win version. :P