package databases

import (
	"database/sql"
	"testing"

	"github.com/blackhorseya/lobster/internal/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestNewMongoDB(t *testing.T) {
	type args struct {
		cfg *config.Config
	}
	tests := []struct {
		name    string
		args    args
		want    *mongo.Client
		wantErr bool
	}{
		{
			name: "local test",
			args: args{&config.Config{
				DB: &config.DB{
					URL:   "mongodb://okr-app:changeme@localhost:27018/okr-db/?connect=direct&authSource=okr-db",
					Debug: true,
				},
			}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewMongoDB(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMongoDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNewMariaDB(t *testing.T) {
	type args struct {
		cfg *config.Config
	}
	tests := []struct {
		name    string
		args    args
		want    *sql.DB
		wantErr bool
	}{
		{
			name: "local test",
			args: args{&config.Config{
				DB: &config.DB{
					URL:   "lobster:changeme@/lobster",
					Debug: true,
				},
			}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewMariaDB(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMariaDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
