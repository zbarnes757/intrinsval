package iex

type serviceEndpoint struct {
	url string
}

// API struct for working with IEX
type API struct {
	serviceEndpoint
}

const iexURI = "https://api.iextrading.com/1.0"

// New will return a new instance of the iex api
func New() *API {
	return &API{
		serviceEndpoint{
			iexURI,
		},
	}
}
