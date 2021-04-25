package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	var (
		base *Error
	)
	err := New(400, "domain", "reason", "message")
	err2 := New(400, "domain", "reason", "message")
	err3 := err.WithMetadata(map[string]string{
		"foo": "bar",
	})
	werr := fmt.Errorf("wrap %w", err)

	if errors.Is(err, new(Error)) {
		t.Errorf("should not be equal: %v", err)
	}
	if !errors.Is(werr, err) {
		t.Errorf("should be equal: %v", err)
	}
	if !errors.Is(werr, err2) {
		t.Errorf("should be equal: %v", err)
	}

	if !errors.As(err, &base) {
		t.Errorf("should be matchs: %v", err)
	}
	if !IsBadRequest(err) {
		t.Errorf("should be matchs: %v", err)
	}

	if code := Code(err); code != err2.Code {
		t.Errorf("got %d want: %s", code, err)
	}
	if reason := Reason(err); reason != err3.Reason {
		t.Errorf("got %s want: %s", reason, err)
	}

	if err3.Metadata["foo"] != "bar" {
		t.Error("not expected metadata")
	}
}
