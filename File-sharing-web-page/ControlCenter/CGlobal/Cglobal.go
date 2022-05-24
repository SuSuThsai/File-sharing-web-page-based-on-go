package CGlobal

var FilePath = "CConfig/Config.yaml"

var ServerIP map[string]int
var WebIP map[string]int

var GinPot *Merchine

type Merchine struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}
