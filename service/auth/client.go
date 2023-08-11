package auth

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type IClient interface {
	GetUserInfo(token string) (UserInfo, error)
}

type Client struct {
	httpClient  *http.Client
	iamEndpoint string
}

var _ IClient = &Client{}

func NewClient(iamEndpoint string) *Client {
	fmt.Println("IAM ENDPOINT:", iamEndpoint)

	return &Client{
		httpClient: &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   30 * time.Second,
		},
		iamEndpoint: iamEndpoint,
	}
}

func (c *Client) GetUserInfo(token string) (UserInfo, error) {
	req, err := http.NewRequest(http.MethodGet, c.iamEndpoint, nil)
	if err != nil {
		log.Printf("error init http request %s", err.Error())
		return UserInfo{}, err
	}
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return UserInfo{}, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode == http.StatusUnauthorized {
		log.Printf("error get user info failed with 401")
		return UserInfo{}, ErrInvalidIAMToken
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("error get user info: %s", string(body))
		return UserInfo{}, ErrCallIAM
	}

	return unmarshalUserInfo(resp.Body)
}
