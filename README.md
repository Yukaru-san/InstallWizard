# InstallWizard
With this library you are able to pack and deliver any amount of files to Windows, Linux and Mac OS.
The main goal of this project was delivering projects, consisting of more than just one binary (e.g. projects with graphics
like UI's), without having to include them into the binary itself and making it unreasonable huge.

Note that this program is written in go but can be used for projects in any language - or just to pack some data. It's all
up to you.

# Usage
When compiled, this project will create a file that you can use to create any number of installers that you want.
Simply put the binary into your desired folder and any data except .git files will be included into your new installer.

# Installation
As said before, you will only need the compiled version of this project. Anyway, if you want to compile it yourself you will need
to install packr beforehand. Head over to [packr's repository](https://github.com/gobuffalo/packr) and install it's first version.
When you are done, you can build the InstallWizard using:
```packr build```<br>
And that's pretty much it. You will receive the binary which you can then use to create new installers.

# Example
Using the InstallWizard is pretty straight forward. But here is a complete example on how it can be used:
<br>(Note that it will be a bit different on linux since it's done from command-line only!)

1. Find a directory that you want to pack:
![example1](https://very.highly.illegal-dark-web-server.xyz/preview/raw/KbwmRXtLV1FbRCtWHOigFivwV)

2. Include the InstallWizard binary into your desired directory:
![example2](https://very.highly.illegal-dark-web-server.xyz/preview/raw/NTgIiDA0ugtl6VcbHpWxqk58e)

3. Run the InstallWizard and tell the program how you want to name your installer
![example3](https://very.highly.illegal-dark-web-server.xyz/preview/raw/gSR9ERD5IZBYudWpChikwsEwm)

4. After the wizard is done, you will see an output directory:
![example4](https://very.highly.illegal-dark-web-server.xyz/preview/raw/5vLWcZV2VUnfpoiHhzeTtsojX)

5. Inside, you will find your compiled installer for Windows, Linux and Mac OS
![example5](https://very.highly.illegal-dark-web-server.xyz/preview/raw/LyVyshiSgJPsSg5QX2DMxJG2z)

Now you can ship the installers to your clients. Upon installation their POV will be:

1. Running the installer:<br>
![example6](https://very.highly.illegal-dark-web-server.xyz/preview/raw/5W9FwYXPn15HqOdqpTxKPLitF)

2. Selecting the directory for the program to be installed in:<br>
![example7](https://very.highly.illegal-dark-web-server.xyz/preview/raw/l3YjZEhWczdgQiTdEh6Z8r7zm)

3. The data will be extracted and you can find everything inside the desired directory:
![example8](https://very.highly.illegal-dark-web-server.xyz/preview/raw/irHjVghKajLygWvntLqFBJRL6)

And that's pretty much it. If you have any suggestions or ideas on how to improve this project, feel free to create an issue!
