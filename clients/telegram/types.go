package telegram

type UpdateResponce struct {
	Ok      bool     `json:"ok"`
	Results []Update `json:"results"`
}

type Update struct {
	ID      int    `json:"update_id"`
	Message string `json:"message"`
}
