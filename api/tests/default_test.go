package tests

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	// _, file, _, _ := runtime.Caller(0)
	// apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
}

// TestGet is a sample to run an endpoint test
func TestGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "/search/test", nil)
	w := httptest.NewRecorder()

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
