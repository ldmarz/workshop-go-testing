package go_testing_test

import (
	"fmt"
	"github.com/ldmarz/workshop-go-testing/go_testing"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type want struct {
	err    error
	result int
}

func Test_SomeFancyHandler(t *testing.T) {

	tests := []struct {
		name string
		url  string
		want string
	}{
		{name: "should fail when send_error argument is missing",
			url:  "/test",
			want: "{\"error\": \"400\",\"message\": \"send_error argument is required\"}",
		},
		{name: "should response and error if send_error is true",
			url:  "/test?send_error=true",
			want: "{\"error\": \"400\",\"message\": \"You ask me for and error\"}",
		},
		{name: "should response the best all seasons pokemon",
			url:  "/test?send_error=false",
			want: "{\"name\": \"charmeleon\",\"type\": \"fire\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.NewServeMux()
			handler.HandleFunc("/test", go_testing.SomeFancyHandler())
			srv := httptest.NewServer(handler)
			defer srv.Close()

			res, _ := http.Get(fmt.Sprintf("%v%v", srv.URL, tt.url))
			bodyBytes, _ := io.ReadAll(res.Body)

			assert.Equal(t, tt.want, string(bodyBytes))
		})
	}
}
