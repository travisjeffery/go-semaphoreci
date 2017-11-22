package semaphoreci

type Client struct {
	token string
}

func New(token string) *Client struct{
	return &Client{token}
}


