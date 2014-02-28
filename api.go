package zoopla

const baseURL = "http://api.zoopla.co.uk/api/v1/"

type Api struct {
	key string
}

func NewApi(key string) *Api {
	f := Api{key}
	return &f
}