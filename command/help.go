package command

import "fmt"

func Help() string {
	const HELP = "Go SolusVM Client API help command\n" +
		"usage: ./gosolusvmclientapi [flag]\n" +
		"\n" +
		"Flags:\n" +
		"\tboot : boot your VPS\n" +
		"\treboot : reboot your VPS\n" +
		"\tshutdown : turn off your VPS\n" +
		"\tstatus : get your VPS status\n" +
		"\tinfo : get your VPS infos\n" +
		"\thelp : this command\n"
	fmt.Printf(HELP)
	return HELP
}
