package link

import (
	"strings"
	"testing"
)

func TestParseHTML(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected []Link
	}{
		{
			name: "single link",
			html: `<html><body><a href="/example">Example</a></body></html>`,
			expected: []Link{
				{Href: "/example", Text: "Example"},
			},
		},
		{
			name: "multiple links",
			html: `<html><body>
				<a href="/first">First</a>
				<a href="/second">Second</a>
			</body></html>`,
			expected: []Link{
				{Href: "/first", Text: "First"},
				{Href: "/second", Text: "Second"},
			},
		},
		{
			name: "link with nested elements",
			html: `<a href="/nested">Hello <span>world</span>!</a>`,
			expected: []Link{
				{Href: "/nested", Text: "Hello world !"}, // 注意空格处理
			},
		},
		{
			name:     "no links",
			html:     `<html><body><p>No links here</p></body></html>`,
			expected: []Link{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.html)
			got, err := ParseHTML(r)
			if err != nil {
				t.Fatalf("ParseHTML() error = %v", err)
			}
			if len(got) != len(tt.expected) {
				t.Fatalf("ParseHTML() got %d links, want %d", len(got), len(tt.expected))
			}
			for i := range got {
				if got[i].Href != tt.expected[i].Href {
					t.Errorf("ParseHTML() link %d Href = %v, want %v", i, got[i].Href, tt.expected[i].Href)
				}
				if got[i].Text != tt.expected[i].Text {
					t.Errorf("ParseHTML() link %d Text = %v, want %v", i, got[i].Text, tt.expected[i].Text)
				}
			}
		})
	}
}
