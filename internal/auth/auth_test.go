package auth

import (
    "reflect"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
    t.Run("no auth header", func(t *testing.T) {
        headers := make(map[string][]string)
        _, err := GetAPIKey(headers)
        if !reflect.DeepEqual(err, ErrNoAuthHeaderIncluded) {
            t.Errorf("expected ErrNoAuthHeaderIncluded, got %v", err)
        }
    })
    t.Run("malformed auth header", func(t *testing.T) {
        headers := map[string][]string{
            "Authorization": {"Bearer"},
        }
        _, err := GetAPIKey(headers)
        if err == nil {
            t.Errorf("expected error, got nil")
        }
    })
    t.Run("correct auth header", func(t *testing.T) {
        headers := map[string][]string{
            "Authorization": {"ApiKey test"},
        }
        key, err := GetAPIKey(headers)
        if err != nil {
            t.Errorf("expected nil, got %v", err)
        }
        if key != "test" {
            t.Errorf("expected test, got %s", key)
        }
    })
}
