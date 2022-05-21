package ALiconfig

type DriverConfig struct {
	AutoRefresh     bool   `mapstructure:"autoRefresh" json:"auto_refresh"`
	UploadRate      int    `mapstructure:"uploadRate" json:"upload_rate"`
	RefreshDuration string `mapstructure:"refreshDuration" json:"refresh_duration"`
}

type RefreshTokenConfig struct {
	RefreshToken string `mapstructure:"refreshToken" json:"refresh_token"`
}
