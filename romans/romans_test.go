package romans

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomans(t *testing.T) {
	assert.Equal(t, "IV", romans(4))
	assert.Equal(t, "V", romans(5))
}
func TestThree(t *testing.T) {
	assert.Equal(t, "I", lastThree(1))
	assert.Equal(t, "II", lastThree(2))
	assert.Equal(t, "III", lastThree(3))
}
