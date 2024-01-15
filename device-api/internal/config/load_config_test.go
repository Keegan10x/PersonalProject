package config_test

import (
	"fmt"
	"main/device-api/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	testPath := "./testpath/test.yaml"
	values, err := config.LoadTemplateConfigMap(testPath)
	assert.NoError(t, err, "expected no error got: %v", err)

	fmt.Println("logging values", values)

	// assert port
	assert.Equal(t, 8080, values.Port)
}
