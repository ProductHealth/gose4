package time

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestSecondsDuration(t *testing.T) {
	parsed, _ := time.ParseDuration("2s")
	assert.Equal(t, parsed, SecondsDuration(2))
}
