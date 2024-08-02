package cmd

//###########################################################//

const (
	ParamLoadSerialLinux   = "ls /dev/tty* | grep -E 'ttyUSB.*|ttyACM.*'"
	ParamLoadSerialMac     = "ls /dev/tty.*"
	ParamLoadSerialWindows = "Get-ItemProperty HKLM:\\HARDWARE\\DEVICEMAP\\SERIALCOMM | Select-Object -Property * -ExcludeProperty PSPath, PSParentPath, PSChildName, PSDrive, PSProvider | ForEach-Object { $_.PSObject.Properties.Value }"
)
