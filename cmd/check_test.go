package cmd

import (
	"testing"
)

func TestBuildConnString(t *testing.T) {
	tests := []struct {
		name string
		cfg  config
		want string
	}{
		{
			name: "all fields set",
			cfg: config{
				host:     "localhost",
				port:     5432,
				user:     "postgres",
				password: "secret",
				dbname:   "testdb",
				sslmode:  "disable",
			},
			want: "host=localhost port=5432 user=postgres password=secret dbname=testdb sslmode=disable connect_timeout=1",
		},
		{
			name: "custom port",
			cfg: config{
				host:     "db.example.com",
				port:     5433,
				user:     "admin",
				password: "pass",
				dbname:   "mydb",
				sslmode:  "require",
			},
			want: "host=db.example.com port=5433 user=admin password=pass dbname=mydb sslmode=require connect_timeout=1",
		},
		{
			name: "empty fields",
			cfg: config{
				host:     "",
				port:     0,
				user:     "",
				password: "",
				dbname:   "",
				sslmode:  "",
			},
			want: "host= port=0 user= password= dbname= sslmode= connect_timeout=1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildConnString(tt.cfg)
			if got != tt.want {
				t.Errorf("buildConnString() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestTryConnect_InvalidHost(t *testing.T) {
	connString := buildConnString(config{
		host:     "invalid-host-that-does-not-exist",
		port:     5432,
		user:     "postgres",
		password: "postgres",
		dbname:   "postgres",
		sslmode:  "disable",
	})

	err := tryConnect(connString)
	if err == nil {
		t.Error("tryConnect() expected error for invalid host, got nil")
	}
}

func TestCheckWithRetry_Failure(t *testing.T) {
	cfg := config{
		host:     "invalid-host-that-does-not-exist",
		port:     5432,
		user:     "postgres",
		password: "postgres",
		dbname:   "postgres",
		retry:    2,
		sleep:    0,
		sslmode:  "disable",
	}

	err := checkWithRetry(cfg)
	if err == nil {
		t.Error("checkWithRetry() expected error after exhausting retries, got nil")
	}
}

func TestCheckWithRetry_ZeroRetries(t *testing.T) {
	cfg := config{
		host:     "localhost",
		port:     5432,
		user:     "postgres",
		password: "postgres",
		dbname:   "postgres",
		retry:    0,
		sleep:    0,
		sslmode:  "disable",
	}

	err := checkWithRetry(cfg)
	if err == nil {
		t.Error("checkWithRetry() with 0 retries expected error, got nil")
	}
}
