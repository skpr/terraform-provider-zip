package archive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMD5(t *testing.T) {
	md5, err := getMD5("./testdata/archive.zip")
	assert.NoError(t, err)
	assert.Equal(t, "d851164ee8c9861e921fab72b1433682", md5)
}
