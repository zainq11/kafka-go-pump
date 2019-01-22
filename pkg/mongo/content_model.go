package mongo

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//todo: Make platform enum
type contentModel struct {
	title string
	platform string
	author string
	data string
	createdDate primitive.DateTime
}

