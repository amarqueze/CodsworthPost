package publication

import "outergeekhub.com/codsworthpost/domain/helper"

type ImagePost struct {
	Id   string
	File []byte
}

func NewImagePost(file []byte) *ImagePost {
	id := helper.Util{}.GenerateId()
	return &ImagePost{Id: id, File: file}
}
