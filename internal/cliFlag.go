package internal

import (
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

var CheckFlag = []cli.Flag{
	&cli.StringFlag{Name: "url", Aliases: []string{"u"}, Required: true, Usage: "input target, eg: http://127.0.0.1:8080"},
	&cli.StringFlag{Name: "plugin", Aliases: []string{"p"}, Usage: "input plugin, eg: graph, default: graph"},
	&cli.StringFlag{Name: "conf", Aliases: []string{"c"}, Usage: "config file path, default: /etc/grafana/grafana.ini"},
	&cli.StringFlag{Name: "db", Aliases: []string{"d"}, Usage: "db file path, default: /var/lib/grafana/grafana.db"},
	&cli.StringFlag{Name: "key", Aliases: []string{"k"}, Usage: "input key, eg: SW2YcwTIb9zpOOhoPsMm, default get from server"},
	&cli.StringFlag{Name: "outfile", Aliases: []string{"o"}, Usage: "output dbfile name."},
}

var Exp = cli.Command{
	Name:        "exp",
	Usage:       "-u [url] -p [plugin] -c [config] -d [db] -k [key]",
	Description: "Get datasource message from server.",
	Action:      CheckTargetFunc,
	Flags:       CheckFlag,
}

var DecodeFlag = []cli.Flag{
	&cli.StringFlag{Name: "file", Aliases: []string{"f"}, Required: true, Usage: "input db file name, eg: grafana.db"},
	&cli.StringFlag{Name: "key", Aliases: []string{"k"}, Required: true, Usage: "input key, eg: SW2YcwTIb9zpOOhoPsMm"},
}

var Decode = cli.Command{
	Name:        "decode",
	Usage:       "-f [dbfile] -k [key]",
	Description: "Decode data_source message from local file.",
	Action:      DecodetFunc,
	Flags:       DecodeFlag,
}

func CliParse(ctx *cli.Context) {
	if ctx.IsSet("url") {
		Target = ctx.String("url")
		if !strings.Contains(Target, "http") {
			Logger.Errorf("Target input error! ==> [%s], need startswith 'http'", Target)
			os.Exit(1)
		}
		if strings.HasSuffix(Target, "/") {
			Target = Target[:len(Target)-1]
		}
	}
	if ctx.IsSet("plugin") {
		Plugin = ctx.String("plugin")
	}
	if ctx.IsSet("conf") {
		ConfFile = ctx.String("conf")
	}
	if ctx.IsSet("db") {
		DBFile = ctx.String("db")
	}
	if ctx.IsSet("outfile") {
		OutFile = ctx.String("outfile")
	}
	if ctx.IsSet("file") {
		LDBFile = ctx.String("file")
	}
	if ctx.IsSet("key") {
		Key = ctx.String("key")
	}
}

func DecodetFunc(ctx *cli.Context) (err error) {
	CliParse(ctx)
	// Open DB.
	printDBInfo(LDBFile)
	return nil
}

func CheckTargetFunc(ctx *cli.Context) (err error) {
	CliParse(ctx)
	Plugin = checkVuln()
	if Plugin == "" {
		return
	}
	if Key == "" {
		Key = getTargetKey()
	}
	if Key == "" {
		Logger.Error("Get Key Failed!")
		return nil
	}
	getAllDatasource()
	Logger.Criticalf("All Done, have nice day!")
	return nil
}
