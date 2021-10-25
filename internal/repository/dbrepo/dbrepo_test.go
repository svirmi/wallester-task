package dbrepo

import (
	"github.com/ekateryna-tln/wallester-task/internal/config"
	"reflect"
	"testing"
)

func TestNewPostgresRepo(t *testing.T) {
	if got := NewPostgresRepo(db.SQL, &config.App{}); !reflect.DeepEqual(got, dbRepo) {
		t.Errorf("NewPostgresRepo() = %v, want %v", got, dbRepo)
	}
}
