package bcache_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/darmiel/yaxc/internal/server"
	"github.com/stretchr/testify/require"
)

func TestFiberMapMemoryCorruption(t *testing.T) {
	var s = server.NewServer(&server.YAxCConfig{})
	s.StartInternal()
	{
		req := httptest.NewRequest("POST", "/helloworld/8a6a8d0bd78b0da907b091a755e69f61", strings.NewReader("8a6a8d0bd78b0da907b091a755e69f61"))
		res, err := s.App.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.StatusCode)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		require.NoError(t, err)
		require.Equal(t, "8a6a8d0bd78b0da907b091a755e69f61", string(body))
	}
	{
		req := httptest.NewRequest("GET", "/hash/helloworld", nil)
		res, err := s.App.Test(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.StatusCode)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		require.NoError(t, err)
		require.Equal(t, "8a6a8d0bd78b0da907b091a755e69f61", string(body))
	}
}
