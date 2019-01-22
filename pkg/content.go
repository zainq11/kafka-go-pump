package root

import "time"

type Content struct {
	title string `json:"title"`
	platform string `json:"platform"`
	author string `json:"name"`
	data string `json:"data"`
	createdDate time.Time
}

type ContentService interface {
	SaveContent(content *Content) error
	GetContentByPlatform(platform string) Content

}