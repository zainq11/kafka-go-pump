package core

type Content struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Data        string `json:"content"`
	CreatedTime string `json:"time"`
}

type Platform string

const (
	TWITTER Platform = "twt"
	REDDIT Platform = "rdt"
)

type ContentService interface {
	SaveContent(content Content, collection string) error
}