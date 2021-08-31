package main

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar

	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErroNoAvatarURL {
		t.Error("AuthAvatar.GetAvatarURL should return ErrNoAvatarURL when no value present")
	}

	testUrl := "http://url-to-gravatar/"
	client.userData = map[string]interface{}{"avatar_url": testUrl}
	url, err = authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("AuthAvatar.GetAvatarURL should return ErrNoAvatarURL when no value present")
	}

	if url != testUrl {
		t.Error("AuthAvatar.GetAvatarURL should return correct URL")
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	client := new(client)
	client.userData = map[string]interface{}{"userid": "0bc83cb571cd1c50ba6f3e8a78ef1346"}
	url, err := gravatarAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("GravatarAvatar.GetAvatarURL should not return an error")
	}
	if url != "//www.gravatar.com/avatar/0bc83cb571cd1c50ba6f3e8a78ef1346" {
		t.Errorf("GravatarAvatar.GetAvatarURL wrongly returned %s", url)
	}
}

func TestFileSystemAvatar(t *testing.T) {

	// make a test avatar file
	filename := path.Join("avatars", "abc.jpg")
	if err := os.MkdirAll("avatars", 0777); err != nil {
		t.Errorf("couldn't make avatar dir: %s", err)
	}
	if err := ioutil.WriteFile(filename, []byte{}, 0777); err != nil {
		t.Errorf("couldn't make avatar: %s", err)
	}
	defer os.Remove(filename)

	var fileSystemAvatar FileSystemAvatar
	user := &chatUser{uniqueID: "abc"}

	url, err := fileSystemAvatar.GetAvatarURL(user)
	if err != nil {
		t.Errorf("FileSystemAvatar.GetAvatarURL should not return an error: %s", err)
	}
	if url != "/avatars/abc.jpg" {
		t.Errorf("FileSystemAvatar.GetAvatarURL wrongly returned %s", url)
	}

}
