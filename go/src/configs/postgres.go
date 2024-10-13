package configs

type PsqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	TimeZone string
}

func (p *PsqlConfig) FormatDsn() string {
	return "host=" + p.Host + " port=" + p.Port + " user=" + p.User + " password=" + p.Password + " dbname=" + p.Dbname + " sslmode=disable TimeZone=" + p.TimeZone
}

func (p *PsqlConfig) String() string {
	return "Host: " + p.Host + "\nPort: " + p.Port + "\nUser: " + p.User + "\nPassword: " + p.Password + "\nDbname: " + p.Dbname + "\nTimeZone: " + p.TimeZone
}
