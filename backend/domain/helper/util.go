package helper

import "github.com/google/uuid"

type Util struct {
}

func (util Util) GenerateId() string {
	return uuid.New().String()
}
