# InstallWizard
With this library you are able to pack and deliver any amount of files to Windows, Linux and Mac OS.
The main goal of this project was delivering projects, consisting of more than just one binary (e.g. projects with graphics
like an UI), without having to include them into the binary itself and making it unreasonable huge and hard to update.

Note that this program is written in go but can be used for projects in any language - or just to pack some data. It's all
up to you.

<b>Please note:</b><br>
The generated installer uses a GUI on Windows and Darwin but CLI on Linux!

# Usage
When compiled, this project will create a file that you can use to create any number of installers you want.
Simply call the binary from your desired directory and all files will be included into your new installer.

All files and directories containing ".git" will be ignored.

# Installation
Either download the precompiled version here:<br>
[Release 2.0](https://github.com/Yukaru-san/InstallWizard/releases/tag/2.0)<br>

Or build your own following these steps:<br>
```git pull https://github.com/Yukaru-san/InstallWizard.git```<br>
```cd InstallWizard```<br>
```go mod tidy```<br>
```go build```<br>
