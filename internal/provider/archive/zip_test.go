package archive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteZip(t *testing.T) {
	data := make(map[string]interface{})
	data["foo"] = "bar"

	path, err := writeZip(data)
	assert.NoError(t, err)
	assert.NotEmpty(t, path)
}
