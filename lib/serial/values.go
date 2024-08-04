package serial

import "errors"

//###########################################################//

const BaudRateMin = 100

var ErrBaudRateMin = errors.New("baudRate too small")

////

var ErrListSerialUnsupportedPlatform = errors.New("unsupported platform")

type ParamLoadSerialObj struct {
	CMD  string
	Args []string
}

var ParamLoadSerialMap = map[string]ParamLoadSerialObj{
	"linux": {
		CMD:  "sh",
		Args: []string{"-c", "ls /dev/tty* | grep -E 'ttyUSB.*|ttyACM.*'"},
	},
	"windows": {
		CMD:  "powershell",
		Args: []string{"Get-ItemProperty HKLM:\\HARDWARE\\DEVICEMAP\\SERIALCOMM | Select-Object -Property * -ExcludeProperty PSPath, PSParentPath, PSChildName, PSDrive, PSProvider | ForEach-Object { $_.PSObject.Properties.Value }"},
	},
	"darwin": {
		CMD:  "sh",
		Args: []string{"-c", "ls /dev/tty.*"},
	},
}
