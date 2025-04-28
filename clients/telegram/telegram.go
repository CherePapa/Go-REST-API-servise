package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"telegram-bot/lib/e"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

func New(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	querry := url.Values{}
	querry.Add("offset", strconv.Itoa(offset))
	querry.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, querry)
	if err != nil {
		return nil, err
	}
	var res UpdateResponce

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Results, err
}

func (c *Client) SendMessage(chatID int, text string) error {
	querry := url.Values{}

	querry.Add("chat_id", strconv.Itoa(chatID))
	querry.Add("text", (text))

	_, err := c.doRequest(sendMessageMethod, querry)
	if err != nil {
		return e.Wrap("cant send message", err)
	}
	return nil
}

func (c *Client) doRequest(method string, querry url.Values) (data []byte, err error) {
	defer func() { err = e.WrapIfErr("cant do requests:", err) }()
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = querry.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
