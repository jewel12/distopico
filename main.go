package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jewel12/distopico/slack"
	"github.com/jewel12/distopico/telescreen"
)

func main() {
	var (
		device     = flag.Int("d", 1, "device ID")
		classifier = flag.String("c", "data/classifier.xml", "cascade classifier xml file")
		token      = flag.String("t", "", "slack token")
	)
	flag.Parse()

	c := telescreen.Conf{
		DeviceID:              *device,
		CascadeClassifierPath: *classifier,
		Slack:                 slack.NewSlackClient(*token),
	}

	err := telescreen.Surveil(&c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}
}
