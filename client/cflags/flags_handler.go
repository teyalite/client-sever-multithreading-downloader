package cflags

import (
	"client/network"
	"errors"
	"flag"
	"fmt"
	"os"
)

/// command line flags
type CLIFlags struct {
	URL        string
	Thread     int
	OutputPath string
	Verbose    bool
	Checksum   string
}

/// default thread count set to 10
const (
	DefaultThreadCount = 10
)

/// parse command line args
func (f *CLIFlags) Parse() error {
	urlString := flag.String("url", "", "Valid URL to download")
	out := flag.String("out", "", "Output path to store the downloaded file")
	t := flag.Int("t", DefaultThreadCount, "Thread count - Number of concurrent downloads")
	checksum := flag.String("checksum", "", "Checksum SHA256(currently supported) to verify file")
	flag.Parse()

	f.URL = *urlString
	f.OutputPath = *out
	f.Thread = *t
	f.Checksum = *checksum

	return nil
}

func (f *CLIFlags) HasValidDownloadURL() (bool, error) {
	if !network.IsValidURL(f.URL) {
		return false, errors.New("Invalid URL, a valid URL is mandatory, pass URL using -url flag")
	}
	return true, nil
}

func (f *CLIFlags) PerformEssentialChecks() {

	ok, err := f.HasValidDownloadURL()
	if !ok {
		fmt.Println(err)
		os.Exit(1)
	}
}
