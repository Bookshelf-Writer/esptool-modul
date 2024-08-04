# esptool


## Compilation

```bash
go get esptool
cd ${GOPATH-$HOME/go}/src/esptool/cmd
go build
./esptool <args>
```

## Info

An attempt was made to improve the fork and it failed.
The repository is being moved to the archive, rewriting on documentation will continue here: [esp32tool](https://github.com/Bookshelf-Writer/esp32tool)

What was done in this fork:
- updated go from 1.15 to 1.22.5
- all output is transferred to zerolog for convenience on all platforms
- work with serial-port rewritten from term to built-in in go (now supports all platforms)
- CLI has been completely rewritten
- Added command to output the list of available serial devices (supported platforms: Linux, Windows, Mac).
- General code optimizations (unfinished)
- Removed unused code


The release [v1.1.1](https://github.com/Bookshelf-Writer/esptool-modul/releases/tag/v1.1.1) is the most up to date, ready compiled applications can be taken from there.
For self-building it is better to take the latest actual version of the master (everything unnecessary was removed before archiving)


#### Optimization is completely finished for:
- serial-port dialog
- CLI

#### Partially realized:
- command generation
- sending messages and parsing received messages
- output (and logs)



