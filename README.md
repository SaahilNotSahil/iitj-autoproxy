# IITJ Autoproxy (v0.1.2)

A CLI tool to automatically login to IIT Jodhpur's network firewall.

## Installation 

### Windows:

#### Scoop:

- First, you need to install [Scoop](https://scoop.sh/).
- Next, run the following commands in PowerShell

```shell
scoop bucket add org https://github.com/SaahilNotSahil/scoop-iitj-autoproxy.git
scoop install iitj-autoproxy
```

#### From archive:

- Download the latest release (`iitj-autoproxy_<version>_windows_amd64.zip`) from [here](https://github.com/SaahilNotSahil/iitj-autoproxy/releases).
- Extract the zip file.
- Open Powershell as administrator and run the following commands

```shell
cd <path-to-extracted-folder>
.\install.ps1
```

- IITJ Autoproxy is now installed on your system in the `C:\Program Files\IITJ Autoproxy` directory, and is added to the PATH.

### Linux:

#### Debian-based Distros (Debian, Ubuntu, Linux Mint, etc.):

- You can install by running the following commands

```bash
curl -s https://packagecloud.io/install/repositories/SaahilNotSahil/iitj-autoproxy/script.deb.sh?any=true | sudo bash
sudo apt update
sudo apt install iitj-autoproxy
```

#### Arch-based Distros (Arch Linux, EndeavourOS, Manjaro, etc.):

- Download the AUR package for IITJ Autoproxy:
  (You can install it with your favourite AUR helper)

```bash
yay -S iitj-autoproxy-bin
```

#### RHEL-based Distros (RHEL, Fedora, CentOS, etc.):

- You can install by running the following commands

```shell
curl -s https://packagecloud.io/install/repositories/SaahilNotSahil/iitj-autoproxy/script.rpm.sh?any=true | sudo bash
sudo yum install iitj-autoproxy
```

#### From archive:

- Download the latest release (`iitj-autoproxy_<version>_linux_<amd64/arm64>.tar.gz`) from [here](https://github.com/SaahilNotSahil/iitj-autoproxy/releases).
- Extract the `tar.gz` file and run the installer script
```shell
tar zxvf iitj-autoproxy_<version>_linux_<amd64/arm64>.tar.gz
chmod +x install.sh
./install.sh
```
- IITJ Autoproxy is now installed on your system in the `/usr/bin` directory, which is already in the PATH.

### Mac:

#### Homebrew:

- You can install by running the following commands

```shell
brew tap SaahilNotSahil/iitj-autoproxy
brew install iitj-autoproxy
```

- To upgrade the package:

```shell
brew upgrade iitj-autoproxy
```

## Usage

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
