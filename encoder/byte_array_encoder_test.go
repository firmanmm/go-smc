package encoder

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteArrayEncoder(t *testing.T) {

	large := make([]byte, 1000)
	for i := 0; i < len(large); i++ {
		large[i] = byte(i % 256)
	}
	largeRandom := make([]byte, 1000)
	rand.Read(largeRandom)

	testData := []struct {
		Name     string
		Value    []byte
		HasError bool
	}{
		{
			"Normal",
			[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
			false,
		},
		{
			"Empty",
			[]byte{},
			false,
		},
		{
			"Large",
			large,
			false,
		},
		{
			"Large Random",
			largeRandom,
			false,
		},
	}

	encoder := NewByteArrayEncoder()

	for _, val := range testData {
		t.Run(val.Name, func(t *testing.T) {
			writer := NewBufferWriter()
			err := encoder.Encode(val.Value, writer)
			assert.Nil(t, err)
			content, err := writer.GetContent()
			assert.Nil(t, err)
			reader := NewSliceReader(content)
			decoded, err := encoder.Decode(reader)
			assert.Nil(t, err)
			assert.EqualValues(t, val.Value, decoded)
		})
	}
}

func TestByteArrayCompabilityB64(t *testing.T) {
	data := make([]byte, 256)
	for i := 0; i < 256; i++ {
		data[i] = byte(i)
	}
	encoder := NewByteArrayEncoder()
	writer := NewBufferWriter()
	err := encoder.Encode(data, writer)
	assert.Nil(t, err)
	content, err := writer.GetContent()
	assert.Nil(t, err)
	encoded := base64.StdEncoding.EncodeToString(content)
	fmt.Println(encoded)
}
