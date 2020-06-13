# InstallWizard
With this library you are able to pack and deliver any amount of files to Windows, Linux and Mac OS.
The main goal of this project was delivering projects, consisting of more than just one binary (e.g. projects with graphics
like an UI), without having to include them into the binary itself and making it unreasonable huge and hard to update.

Note that this program is written in go but can be used for projects in any language - or just to pack some data. It's all
up to you.

# Usage
When compiled, this project will create a file that you can use to create any number of installers you want.
Simply put the binary into your desired folder and any data except .git files will be included into your new installer.

# Installation
As said before, you will only need the compiled version of this project. Anyway, if you want to compile it yourself you will need
to install packr beforehand. Head over to [packr's repository](https://github.com/gobuffalo/packr) and install it's first version.
When you are done, you can build the InstallWizard using:
```packr build```<br>
And that's pretty much it. You will receive the binary which you can then use to create new installers.

# Download binaries
If you want to spare yourself the trouble of creating your own binary, you can head over to
the [release tab](https://github.com/Yukaru-san/InstallWizard/releases/tag/1.0) and download
a precompiled version for your OS.


# Example
## Creating an Installer
Using the InstallWizard is pretty straight forward. But here is a complete example on how it can be used:
<br>(Note that it will be a bit different on linux since it's done from command-line only!)

1. Find a directory that you want to pack:<br>
![example1](https://very.highly.illegal-dark-web-server.xyz/preview/raw/qSCYeEawhci7d4qROuqPYOnVy)

2. Include the InstallWizard binary into your desired directory:<br>
![example2](https://very.highly.illegal-dark-web-server.xyz/preview/raw/5FwBaK2tf93xmWoe6c1Gp4aNA)

3. Run the InstallWizard and tell the program how you want to name your installer:<br>
![example3](https://very.highly.illegal-dark-web-server.xyz/preview/raw/djDx1VbYyAs2fmVFXV54tXhhD)

4. After the wizard is done, you will see an output directory:<br>
![example4](https://very.highly.illegal-dark-web-server.xyz/preview/raw/opi5loj0Cv7poWw1Mts2cjI2f)

5. Inside, you will find your compiled installer for Windows, Linux and Mac OS:<br>
![example5](https://very.highly.illegal-dark-web-server.xyz/preview/raw/6XzXxlptCdwgp1OoJsfTFsTLH)

Now you can ship the installers to your clients.

## Client perspective: How to use the installer

1. Running the installer:<br>
![example6](https://very.highly.illegal-dark-web-server.xyz/preview/raw/8YnbLNlgclgEbVXKpg4pK9Jba)

2. Selecting the directory for the program to be installed in:<br>
![example7](https://very.highly.illegal-dark-web-server.xyz/preview/raw/Tc1nAoqKnjvqhkjQZmQ1di8K9)

3. The data will be extracted and you can find everything inside the desired directory:<br>
![example8](https://very.highly.illegal-dark-web-server.xyz/preview/raw/ItAriACKIVja5Bs7JhGCpAqyB)

And that's pretty much it. If you have any suggestions or ideas on how to improve this project, feel free to create an issue!
