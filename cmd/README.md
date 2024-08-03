# CLI esptool

This is a console application for working with esptool-mod

## Build

```bash
go build -o esptool.bin
```

## work

### Help
```bash
./esptool.bin h
```

### Version
```bash
./esptool.bin ver
```

### Serial Device List
```bash
./esptool.bin -list
```

---

### Info
```bash
./esptool.bin -info -port /dev/ttyACM0 logTrase
```