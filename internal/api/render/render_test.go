// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package render

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/harness/gitness/internal/api/usererror"
)

func TestWriteErrorf(t *testing.T) {
	w := httptest.NewRecorder()

	e := usererror.New(500, "abc")
	UserError(w, e)

	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &usererror.Error{}
	if err := json.NewDecoder(w.Body).Decode(errjson); err != nil {
		t.Error(err)
	}
	if got, want := errjson.Message, e.Message; got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteErrorCode(t *testing.T) {
	w := httptest.NewRecorder()

	ErrorMessagef(w, 418, "pc load letter %d", 1)

	if got, want := w.Code, 418; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &usererror.Error{}
	if err := json.NewDecoder(w.Body).Decode(errjson); err != nil {
		t.Error(err)
	}
	if got, want := errjson.Message, "pc load letter 1"; got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteNotFound(t *testing.T) {
	w := httptest.NewRecorder()

	NotFound(w)

	if got, want := w.Code, 404; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &usererror.Error{}
	if err := json.NewDecoder(w.Body).Decode(errjson); err != nil {
		t.Error(err)
	}
	if got, want := errjson.Message, usererror.ErrNotFound.Message; got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteUnauthorized(t *testing.T) {
	w := httptest.NewRecorder()

	Unauthorized(w)

	if got, want := w.Code, 401; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &usererror.Error{}
	if err := json.NewDecoder(w.Body).Decode(errjson); err != nil {
		t.Error(err)
	}
	if got, want := errjson.Message, usererror.ErrUnauthorized.Message; got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteForbidden(t *testing.T) {
	w := httptest.NewRecorder()

	Forbidden(w)

	if got, want := w.Code, 403; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &usererror.Error{}
	if err := json.NewDecoder(w.Body).Decode(errjson); err != nil {
		t.Error(err)
	}
	if got, want := errjson.Message, usererror.ErrForbidden.Message; got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteBadRequest(t *testing.T) {
	w := httptest.NewRecorder()

	BadRequest(w)

	if got, want := w.Code, 400; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	errjson := &usererror.Error{}
	if err := json.NewDecoder(w.Body).Decode(errjson); err != nil {
		t.Error(err)
	}
	if got, want := errjson.Message, usererror.ErrBadRequest.Message; got != want {
		t.Errorf("Want error message %s, got %s", want, got)
	}
}

func TestWriteJSON(t *testing.T) {
	// without indent
	{
		w := httptest.NewRecorder()
		JSON(w, http.StatusTeapot, map[string]string{"hello": "world"})
		if got, want := w.Body.String(), "{\"hello\":\"world\"}\n"; got != want {
			t.Errorf("Want JSON body %q, got %q", want, got)
		}
		if got, want := w.Header().Get("Content-Type"), "application/json; charset=utf-8"; got != want {
			t.Errorf("Want Content-Type %q, got %q", want, got)
		}
		if got, want := w.Code, http.StatusTeapot; got != want {
			t.Errorf("Want status code %d, got %d", want, got)
		}
	}
	// with indent
	{
		indent = true
		defer func() {
			indent = false
		}()
		w := httptest.NewRecorder()
		JSON(w, http.StatusTeapot, map[string]string{"hello": "world"})
		if got, want := w.Body.String(), "{\n  \"hello\": \"world\"\n}\n"; got != want {
			t.Errorf("Want JSON body %q, got %q", want, got)
		}
	}
}
