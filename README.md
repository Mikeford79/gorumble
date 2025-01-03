# gorumbot
[rumble](https://rumble.com) livestream viewbot written in go!

## How does it work?
This viewbot doesn't depend on headless browsers or proxies. Instead, we obtain valid viewer IDs from Rumble and use them to increase livestream views.
It continues to run, sending those viewers every 60 seconds until you stop it.

## How to install golang

Here's are some resources to learn about golang and how to install it on windows
- [GeeksforGeeks - How to Install Go on Windows](https://www.geeksforgeeks.org/how-to-install-go-on-windows/)
- [YouTube - Install GO on Windows 11 in 2 minutes](https://www.youtube.com/watch?v=EPpZbwAr4k8)
Install Go (Golang): If you haven't already, you need to install Go on your system. You can download the installer from the official Go website and follow the installation instructions for your operating system.

## Set Up Your Project Directory
Create a new directory for your project and navigate to it in the Command Prompt.

```sh
mkdir myviewbot
cd myviewbot
```
## Create a Go Module
Initialize a new Go module in your project directory. This will create a go.mod file.

```sh
go mod init myviewbot
```
## Install Dependencies
Install the necessary dependencies for your project. In this case, you will need the github.com/itsryuku/gorumbot and github.com/eiannone/keyboard packages.

```sh
go get github.com/itsryuku/gorumbot
go get github.com/eiannone/keyboard
```
## Run the Code
Finally, you can run the code using the go run command followed by the necessary flags for your video URL and the number of bots.

```sh
go run main.go -u <your-video-url-here> -b <number-of-bots>
Replace <your-video-url-here> with the actual URL of the video you want to viewbot and <number-of-bots> with the number of bot viewers you want to start with.
```


