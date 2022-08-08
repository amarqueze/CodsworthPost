package firebasefilesprovider

import (
	"outergeekhub.com/codsworthpost/domain/model/publication"
	"outergeekhub.com/codsworthpost/domain/model/share"
)

type UploadServiceAdapter struct {
	Version string
}

func NewUploadService() publication.UploadService {
	return &UploadServiceAdapter{Version: "1.0.0"}
}

func (adapter *UploadServiceAdapter) UploadMedia() (string, share.BusinessError) {
	return "https://some-image.ong", nil
}
