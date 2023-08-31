# IITJ Autoproxy (v0.1.1)

A CLI tool to automatically login to IIT Jodhpur's network firewall.

#### (Currently only supports Windows and Linux)

### Installation 

#### Windows:

- Download the latest release from [here](https://github.com/XanderWatson/iitj-autoproxy/releases).
- Unzip the file and open a powershell instance with administrator privileges in the unzipped directory:
- Run the following script to install autoproxy:

```
.\install.ps1
```

This will install the autoproxy CLI and daemon in "C:\Program Files\IITJ Autoproxy\" and also add it to the PATH environmment variable.

#### Debian-based Linux Distributions:
- Run the following commands to install the package:

```
$ curl -s https://packagecloud.io/install/repositories/XanderWatson/iitj-autoproxy/script.deb.sh?any=true | sudo bash
$ sudo apt update
$ sudo apt install iitj-autoproxy
```

- The daemon service will be started automatically. Incase it doesn't, run the following command:

```
$ sudo autoproxyd-start
```

- To stop the daemon service, run the following command:

```
$ sudo autoproxyd-stop
```

#### Arch-based Linux Distributions:
- Install the [AUR package](https://aur.archlinux.org/packages/iitj-autoproxy-bin/) using your favourite AUR helper.

For example, using `yay`:
```
$ yay -S iitj-autoproxy-bin
```

- Start the daemon service:

```
$ sudo autoproxyd-start
```

- To stop the daemon service, run the following command:

```
$ sudo autoproxyd-stop
```

### Usage

- First, you need to set your username and password. To do so, run the following command:

```
$ autoproxy config
```

- Enter your username and password when prompted.


- To login to the firewall once, run the following command:

```
$ autoproxy login
```

- To automatically login to the firewall everytime it is required, run the following command:

```
$ autoproxy schedule
```

Note: You need to be logged out before running the above command.

- To logout from the firewall, run the following command:

```
$ autoproxy logout
```

#### Support for other operating systems coming soon :)
