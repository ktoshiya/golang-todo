package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ktoshiya/golang-todo/entity"
	"github.com/ktoshiya/golang-todo/store"
	"github.com/ktoshiya/golang-todo/testutil"
)

func TestAddTask(t *testing.T) {
	t.Parallel()
	type want struct {
		status  int
		rspFile string
	}
	test := map[string]struct {
		reqFile string
		want    want
	}{
		"ok": {
			reqFile: "testdata/add_task/ok_req.json.golden",
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/add_task/ok_rsp.json.golden",
			},
		},
		"badRequest": {
			reqFile: "testdata/add_task/bad_req.json.golden",
			want: want{
				status:  http.StatusBadRequest,
				rspFile: "testdata/add_task/bad_req_rsp.json.golden",
			},
		},
	}

	for n, tt := range test {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/tasks",
				bytes.NewReader(testutil.LoadFile(t, tt.reqFile)),
			)

			sut := AddTask{
				Store: &store.TaskStore{
					Tasks: map[entity.TaskID]*entity.Task{},
				},
			}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t, resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile))
		})
	}
}
