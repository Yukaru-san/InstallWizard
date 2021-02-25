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


will need
to install packr beforehand. Head over to [packr's repository](https://github.com/gobuffalo/packr) and install it's first version.
When you are done, you can build the InstallWizard using:
```packr build```<br>
And that's pretty much it. You will receive the binary which you can then use to create new installers.

# Example
## Creating an Installer
Using the InstallWizard is pretty straight forward. But here is a complete example on how it can be used:
<br>**Note that the user experience will be a bit different on linux since it's done from command-line only!**<br>

1. Find a directory that you want to pack:<br>
![example1](https://files.jojii.de/preview/raw/qSCYeEawhci7d4qROuqPYOnVy)

2. Include the InstallWizard binary into your desired directory:<br>
![example2](https://files.jojii.de/preview/raw/5FwBaK2tf93xmWoe6c1Gp4aNA)

3. Run the InstallWizard and tell the program how you want to name your installer:<br>
![example3](https://files.jojii.de/preview/raw/djDx1VbYyAs2fmVFXV54tXhhD)

4. After the wizard is done, you will see an output directory:<br>
![example4](https://files.jojii.de/preview/raw/opi5loj0Cv7poWw1Mts2cjI2f)

5. Inside, you will find your compiled installer for Windows, Linux and Mac OS:<br>
![example5](https://files.jojii.de/preview/raw/6XzXxlptCdwgp1OoJsfTFsTLH)

Now you can ship the installers to your clients.

## Client perspective: How to use the installer

1. Running the installer:<br>
![example6](https://files.jojii.de/preview/raw/8YnbLNlgclgEbVXKpg4pK9Jba)

2. Selecting the directory for the program to be installed in:<br>
![example7](https://files.jojii.de/preview/raw/Tc1nAoqKnjvqhkjQZmQ1di8K9)

3. The data will be extracted and you can find everything inside the desired directory:<br>
![example8](https://files.jojii.de/preview/raw/ItAriACKIVja5Bs7JhGCpAqyB)

And that's pretty much it. If you have any suggestions or ideas on how to improve this project, feel free to create an issue!
