package account_api_lib

type AccountHandlerConfig struct {
	URI string
}

func NewAccountHandleConfig(uri string) *AccountHandlerConfig {
	config := new(AccountHandlerConfig)
	config.URI = uri
	return config
}
