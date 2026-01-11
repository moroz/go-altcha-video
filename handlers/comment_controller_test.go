package handlers_test

import (
	"bytes"
	"database/sql"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"testing"
	"time"

	_ "modernc.org/sqlite"

	"github.com/google/uuid"
	"github.com/moroz/go-altcha-video/config"
	"github.com/moroz/go-altcha-video/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ProjectRoot = config.MustGetenv("PROJECT_ROOT")
var DbConnString = config.GetenvWithDefault("TEST_DATABASE_URL", "./db/test.db")

func TestCreateComment(t *testing.T) {
	resolvedPath := filepath.Join(ProjectRoot, DbConnString)

	db, err := sql.Open("sqlite", resolvedPath)
	require.NoError(t, err)

	err = db.PingContext(t.Context())
	require.NoError(t, err)

	postId, err := uuid.NewV7()
	require.NoError(t, err)

	_, err = db.ExecContext(
		t.Context(),
		"insert into posts (id, title, body, slug, published_at) values (?, ?, ?, ?, ?)",
		postId, "Example Post", "example-post", "Content", time.Now().Unix(),
	)
	require.NoError(t, err)

	router := handlers.Router(db)
	w := httptest.NewRecorder()

	params := url.Values{
		"body":      {"I am a comment"},
		"signature": {"Clanker User"},
		"website":   {"https://farmacia-hombres.es"},
	}

	body := bytes.NewBufferString(params.Encode())
	req := httptest.NewRequest("POST", "/blog/example-post/comments", body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Result().StatusCode)
}
