```go
package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	config := GetConfig()

	assert.Equal(t, config.DB.Driver, "postgres")
	assert.Equal(t, config.DB.Host, "localhost")
	assert.Equal(t, config.DB.Port, "5432")
	assert.Equal(t, config.DB.User, "postgres")
	assert.Equal(t, config.DB.Password, "password")
	assert.Equal(t, config.DB.Name, "testdb")
}
```