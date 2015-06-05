package compressed

import "testing"

func TestExtension(t *testing.T) {
	curLen := 26
	if len(Extensions) != curLen {
		t.Fatalf("Length doesn't match. Expected %d, Got %d", curLen, len(Extensions))
	}
}

func TestIs(t *testing.T) {
	if Is("a/b/c/bar.css") {
		t.Fatal("Wrong detection. Must not be compressed")
	}

	if !Is("a/b/c/bar.lzma") {
		t.Fatal("Wrong detection. Must be compressed")
	}

	if !Is("a/b/c/bar.LZMA") {
		t.Fatal("Wrong detection. Must be compressed")
	}

	if Is("a/b/c/7zlzma") {
		t.Fatal("Wrong detection. Must not be compressed")
	}
}
