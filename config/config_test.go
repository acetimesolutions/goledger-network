package config_test

import (
	"testing"

	"github.com/acetimesolutions/blockchain-golang/config"
)

func TestInit(t *testing.T) {
	var config config.Config
	config.LoadConfigs()

}