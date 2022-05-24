package Temp

type DriverConfig struct {
	AutoRefresh     bool   `mapstructure:"autoRefresh" json:"auto_refresh"`
	UploadRate      int    `mapstructure:"uploadRate" json:"upload_rate"`
	RefreshDuration string `mapstructure:"refreshDuration" json:"refresh_duration"`
}

type RefreshTokenConfig struct {
	RefreshToken string `mapstructure:"refreshToken" json:"refresh_token"`
}

type MysqlInfo struct {
	Db         string `mapstructure:"db" json:"db"`
	Host       string `mapstructure:"host" json:"host"`
	Port       int    `mapstructure:"port" json:"port"`
	DbUser     string `mapstructure:"dbUser" json:"db_user" gorm:"DEFAULT:root"`
	DbPassword string `mapstructure:"dbPassword" json:"db_password"`
	DbName     string `mapstructure:"dbName" json:"db_name"`
}

type RedisInfo struct {
	DbR         string `mapstructure:"dbr" json:"db_r"`
	HostR       string `mapstructure:"hostr" json:"host_r"`
	PortR       int    `mapstructure:"portr" json:"port_r"`
	DbUserR     string `mapstructure:"dbuserr" json:"dbuser_r" gorm:"DEFAULT:defalut"`
	DbPassWordR string `mapstructure:"dbpasswordr" json:"db_pass_word_r"`
	DBModel     int    `mapstructure:"dbModel" json:"db_model"`
}

type JWT struct {
	JwtKey string `mapstructure:"jwtKey" json:"jwt_key"`
}
