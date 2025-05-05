# PVM for Windows

[Support this project](https://github.com/sponsors/hjbdev)

Removing the hassle of changing PHP versions in the CLI on Windows.

This package has a much more niche use case than nvm does. When developing on Windows and using the integrated terminal, it's quite difficult to get those terminals to _actually_ listen to PATH changes.

This utility changes that.

## Build this project
Download and install GO from [here](https://go.dev/doc/install)

To compile this project use:
```shell
set GOOS=windows
set GOARCH=amd64
go build -o pvm.exe
```

## Installation

Create the destination folder (e.g. `C:\pvm`) and drop the compiled `pvm.exe` in there. Add the folder to your PATH.

## Commands
```
pvm list
```
Will list out all the available PHP versions you have installed

```
pvm path
```
Will tell you what to put in your Path variable.

```
pvm use 8.2.9
```
> [!NOTE]  
> Versions must have major.minor specified in the *use* command. If a .patch version is omitted, newest available patch version is chosen.

Will switch your currently active PHP version to PHP 8.2.9

```
pvm install 8.2
```
> [!NOTE]  
> The install command will automatically determine the newest minor/patch versions if they are not specified

Will install PHP 8.2 at the latest patch.
