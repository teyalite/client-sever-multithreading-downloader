package main

import (
	"client/cflags"
	"client/downloader"
	"client/util"
	"fmt"
)

/// execute the downloader
func execute() {
	flags := cflags.CLIFlags{}

	err := flags.Parse()
	if err != nil {
		/// error while parsing flags
		fmt.Println("Error parsing flags: ", err)
		return
	}

	/// flags checks
	flags.PerformEssentialChecks()

	/// a session id generated using the number of threads and the url
	sessionID := util.GenHash(flags.URL, flags.Thread)

	d := downloader.Downloader{
		Flags:     flags,
		SessionID: sessionID,
	}

	d.Run()
}

func main() {
	execute()
}
