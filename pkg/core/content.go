package core

type Content struct {
	Title string `json:"title"`
	Platform string `json:"platform"`
	Author string `json:"name"`
	Data string `json:"data"`
	CreatedDate string `json:"time"`
}

type ContentService interface {
	SaveContent(content Content, collection string) error
}