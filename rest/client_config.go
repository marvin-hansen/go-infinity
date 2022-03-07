package rest

const HTTP_PROT = "http://"

type ClientConfig struct {
	Protocol         string
	Host             string
	Port             string
	Timeout          int
	connectionString string
}

func NewClientConfig(host, port string, timeout int) *ClientConfig {
	return &ClientConfig{
		Host:             host,
		Port:             port,
		Timeout:          timeout,
		connectionString: HTTP_PROT + host + ":" + port,
	}
}

func (c ClientConfig) GetConnectionString() string {
	return c.connectionString
}
