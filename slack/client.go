package slack

import (
	"image"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/slack-go/slack"
)

type SlackClient struct {
	Client *slack.Client
}

func NewSlackClient(token string) *SlackClient {
	return &SlackClient{
		Client: slack.New(token),
	}
}

func (s *SlackClient) SetUserPhoto(img image.Image) error {
	tf, err := ioutil.TempFile(os.TempDir(), "profile.png")
	if err != nil {
		return err
	}
	png.Encode(tf, img)
	if err := s.Client.SetUserPhoto(tf.Name(), slack.UserSetPhotoParams{}); err != nil {
		return err
	}
	return nil
}
