# IITJ Autoproxy

A CLI tool to automatically login to IIT Jodhpur's network firewall.

#### (Currently only supports Linux)

### Installation 

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
