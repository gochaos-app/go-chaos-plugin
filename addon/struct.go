package modules

type ChaosConfig struct {
	Region    string
	Service   string
	Project   string
	Namespace string
	Tag       string
	Chaos     string
	Count     int
}
