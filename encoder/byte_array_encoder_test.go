package encoder

import (
	"math/rand"
	"reflect"
	"testing"
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
			tracker := GetBufferTracker()
			encoded, err := encoder.Encode(val.Value, tracker)
			if err != nil != val.HasError {
				t.Errorf("Expected error value of %v but got %v", val.HasError, err != nil)
			}
			decoded, err := encoder.Decode(encoded)
			if err != nil != val.HasError {
				t.Errorf("Expected error value of %v but got %v", val.HasError, err != nil)
			}
			if !reflect.DeepEqual(val.Value, decoded) {
				t.Errorf("Expected %v but got %v", val.Value, decoded)
			}
			PutBufferTracker(tracker)
		})
	}
}
