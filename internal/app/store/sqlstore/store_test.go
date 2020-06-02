package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M)  {
	databaseURL = os.Getenv(`DATABASE_URL`)
	if databaseURL == "" {
		databaseURL = "host=192.168.1.175 user=postgres dbname=restapi_test sslmode=disable password=postgres"
	}
	os.Exit(m.Run())
}
