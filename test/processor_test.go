package processor

import (
    "strings"
    "testing"
)

func TestProcessURLs(t *testing.T) {
    input := strings.NewReader("https://example.com/path?query=1&action=edit\nhttps://example.com/path?query=2&action=edit")
    err := ProcessURLs(input, "")
    if err != nil {
        t.Errorf("ProcessURLs failed: %v", err)
    }
}
