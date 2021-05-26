package main

import "errors"

var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL.")

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}

	return "", ErrNoAvatarURL
}

type Avatar interface {
	GetAvatarURL(c *client) (string, error)
}
