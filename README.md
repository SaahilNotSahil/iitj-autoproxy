# IITJ Autoproxy (v0.1.2)

A CLI tool to automatically login to IIT Jodhpur's network firewall.

Supports Windows, Debian-based Linux distributions, Arch-based Linux distributions and MacOS.

## Note: OS keyring issues on MacOS. Please refrain from using this on MacOS until the issues are fixed.

### Installation 

#### Windows:

- Download the latest release from [here](https://github.com/SaahilNotSahil/iitj-autoproxy/releases).
- Unzip the file and open a powershell instance with administrator privileges in the unzipped directory:
- Run the following script to install autoproxy:

```
.\install.ps1
```

This will install the autoproxy CLI and daemon in "C:\Program Files\IITJ Autoproxy\" and also add it to the PATH environment variable.

#### Debian-based Linux Distributions:

- Run the following commands to install the package:

```
$ curl -s https://packagecloud.io/install/repositories/SaahilNotSahil/iitj-autoproxy/script.deb.sh | sudo bash
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

#### MacOS:

- Install homebrew if you haven't already. Instructions [here](https://brew.sh/).
- Run the following command to install the package:

```
$ brew install XanderWatson/iitj-autoproxy/iitj-autoproxy
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
