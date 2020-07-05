package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("Shoud: return default value", func(t *testing.T) {
		s := New()
		assert.Equal(t, s.GetPort(), defaultPort)
		assert.Equal(t, s.GetDbHost(), "")
		assert.Equal(t, s.GetDbPort(), "")
		assert.Equal(t, s.GetDbName(), "")
		assert.Equal(t, s.GetDbUser(), "")
		assert.Equal(t, s.GetDbPassword(), "")
		assert.Equal(t, s.GetIncidentServerAddress(), "")
		assert.Equal(t, s.WithIncident(), false)
	})
}

func TestCfg_GetClientAddress(t *testing.T) {

}

func TestCfg_GetPort(t *testing.T) {
	t.Run("Should: return from env", func(t *testing.T) {
		err := os.Setenv(ENV_PORT, "11124")
		if err != nil {
			assert.NotNil(t, nil)
		}
		s := New()
		assert.Equal(t, s.GetPort(), int32(11124))
	})
}

func TestCfg_GetDbHost(t *testing.T) {
	t.Run("Should: return from env", func(t *testing.T) {
		err := os.Setenv(ENV_DB_HOST, "dbhost")
		if err != nil {
			assert.NotNil(t, nil)
		}
		s := New()
		assert.Equal(t, s.GetDbHost(), "dbhost")
	})
}

func TestCfg_GetDbPort(t *testing.T) {
	t.Run("Should: return from env", func(t *testing.T) {
		err := os.Setenv(ENV_DB_PORT, "dbport")
		if err != nil {
			assert.NotNil(t, nil)
		}
		s := New()
		assert.Equal(t, s.GetDbPort(), "dbport")
	})
}

func TestCfg_GetDbName(t *testing.T) {
	t.Run("Should: return from env", func(t *testing.T) {
		err := os.Setenv(ENV_DB_NAME, "dbname")
		if err != nil {
			assert.NotNil(t, nil)
		}
		s := New()
		assert.Equal(t, s.GetDbName(), "dbname")
	})
}

func TestCfg_GetDbUser(t *testing.T) {
	t.Run("Should: return from env", func(t *testing.T) {
		err := os.Setenv(ENV_DB_USER, "dbuser")
		if err != nil {
			assert.NotNil(t, nil)
		}
		s := New()
		assert.Equal(t, s.GetDbUser(), "dbuser")
	})
}

func TestCfg_GetDbPassword(t *testing.T) {
	t.Run("Should: return from env", func(t *testing.T) {
		err := os.Setenv(ENV_DB_PASSWORD, "dbpassword")
		if err != nil {
			assert.NotNil(t, nil)
		}
		s := New()
		assert.Equal(t, s.GetDbPassword(), "dbpassword")
	})
}

func TestCfg_GetIncidentServerAddress(t *testing.T) {
	t.Run("Should: return from env", func(t *testing.T) {
		err := os.Setenv(ENV_INCIDENT_SERVER_HOST, "dbpassword")
		if err != nil {
			assert.NotNil(t, nil)
		}
		s := New()
		assert.Equal(t, s.GetIncidentServerAddress(), "dbpassword")
	})
}

func TestCfg_WithIncident(t *testing.T) {
	t.Run("Should: return from env", func(t *testing.T) {
		err := os.Setenv(ENV_ENABLE_INCIDENT, "true")
		if err != nil {
			assert.NotNil(t, nil)
		}
		s := New()
		assert.Equal(t, s.WithIncident(), true)
	})
}
