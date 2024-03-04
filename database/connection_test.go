package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectionDSNString(t *testing.T) {
	testCases := []struct {
		conn        Connection
		expectedDSN string
	}{
		{Connection{}, "@/"},
		{Connection{Username: "foo"}, "foo@/"},
		{Connection{Username: "foo", Password: "bar"}, "foo:bar@/"},
		{Connection{Username: "foo", Password: "bar", DBName: "foobar"}, "foo:bar@/foobar"},
		{Connection{
			Username: "foo",
			Password: "bar",
			DBName:   "foobar",
			Address:  "localhost:3600",
		},
			"foo:bar@tcp(localhost:3600)/foobar",
		},
		{Connection{
			Username: "foo",
			Password: "bar",
			DBName:   "foobar",
			Address:  "/path/to/socket",
			Protocol: "unix",
		},
			"foo:bar@unix(/path/to/socket)/foobar",
		},
		{Connection{
			Username: "foo",
			Password: "bar",
			DBName:   "foobar",
			Address:  "/path/to/socket",
			Protocol: "unix",
			Params:   map[string]string{"foo": "bar"},
		},
			"foo:bar@unix(/path/to/socket)/foobar?foo=bar",
		},
		{Connection{
			Username: "foo",
			DBName:   "foobar",
			Params:   map[string]string{"foo": "bar"},
		},
			"foo@/foobar?foo=bar",
		},
		{Connection{
			Username: "foo",
			Password: "bar",
			DBName:   "foobar",
			Address:  "/path/to/socket",
			Protocol: "unix",
			Params:   map[string]string{"foo": "bar", "baz": "fee"},
		},
			"foo:bar@unix(/path/to/socket)/foobar?foo=bar&baz=fee",
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expectedDSN, tc.conn.DSN())
	}
}
