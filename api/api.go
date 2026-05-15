package api

type PocketBase struct {
	authToken string
	baseURL   string
}

func NewPocketBase(baseURL string) (*PocketBase, error) {
	token, err := GetAuthToken(baseURL)
	if err != nil {
		return nil, err
	}

	return &PocketBase{
		authToken: token,
		baseURL:   baseURL,
	}, nil
}


