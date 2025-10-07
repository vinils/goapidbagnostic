package repository

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectionString(test *testing.T) {
	cnnConf := ConnectionConfig{
		Host:     "host",
		Port:     46554,
		User:     "user",
		DBName:   "dbname",
		Password: "password",
	}

	expected := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		cnnConf.Host, cnnConf.Port, cnnConf.User, cnnConf.DBName, cnnConf.Password)

	actual := cnnConf.CastToString()

	assert.Equal(test, expected, actual)
}
