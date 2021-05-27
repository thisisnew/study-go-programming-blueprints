package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)

	if tracer == nil {
		t.Error("Return from New should not be nil")
	} else {
		tracer.Trace("Hello trace package")
		if buf.String() != "Hello trace package.\n" {
			t.Errorf("Trace should not write '%s'.", buf.String())
		}
	}
}

//
//func TestFileSystemAvatar(t *testing.T) {
//	filename := path.Join("avatars", "abc.jpg")
//	if err := os.MkdirAll("avatars", 0777); err != nil {
//		t.Errorf("couldn't make avatar dir: %s", err)
//	}
//	if err := ioutil.WriteFile(filename, []byte{}, 0777); err != nil {
//		t.Errorf("couldn't make avatar: %s", err)
//	}
//	defer os.Remove(filename)
//
//	var fileSystemAvatar FileSystemAvatar
//	user := &chatUser{uniqueID: "abc"}
//
//	url, err := fileSystemAvatar.GetAvatarURL(user)
//	if err != nil {
//		t.Errorf("FileSystemAvatar.GetAvatarURL should not return an error: %s", err)
//	}
//	if url != "/avatars/abc.jpg" {
//		t.Errorf("FileSystemAvatar.GetAvatarURL wrongly returned %s", url)
//	}
//}
