package configuration

type CoreConfiguration struct {
	Modeling           bool    `json:"modeling"`
	LogFile            string  `json:"log_file"`
	LocalizationRadius float64 `json:"localization_radius"`
	ThreadAmount       int8    `json:"thread_amount"`
}

type ServerConfiguration struct {
	ServerAddress      string `json:"server_address"`
	ServerPort         int    `json:"server_port"`
	LocalizationTrials int8   `json:"localization_trials"`
}

type ClientConfiguration struct {
	AudioPath     string `json:"audio_path"`
	ProxiesAmount int    `json:"proxies_amount"`
}
