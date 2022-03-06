package rest

func getEndpoint(config *ClientConfig) string {
	if config == nil {
		return DefaultEndpoint
	} else {
		return config.GetConnectionString()
	}
}

func checkError(err error) error {
	if err != nil {
		return err
	} else {
		return nil
	}
}
