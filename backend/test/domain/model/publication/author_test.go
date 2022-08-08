package publication_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"outergeekhub.com/codsworthpost/domain/model/publication"
	"outergeekhub.com/codsworthpost/domain/model/share"
)

func TestCreatedAuthorSuccess(t *testing.T) {
	author := publication.NewAuthor("Alan Marquez", "This is a great Writer")
	assert.EqualValues(t, author.Fullname, "Alan Marquez")
}

func TestCreatedAuthorFailed_AuthorNameIsRequired(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			if serr, ok := err.(share.BusinessError); ok {
				assert.EqualValues(t, serr.Error(), "Author's name is required")
			}
		}
	}()

	publication.NewAuthor("", "")
	t.Errorf("it should throw a BusinessError")
}
