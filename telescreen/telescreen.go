package telescreen

import (
	"fmt"

	"github.com/jewel12/distopico/slack"
	"gocv.io/x/gocv"
)

type Conf struct {
	DeviceID              int
	CascadeClassifierPath string
	Slack                 *slack.SlackClient
}

func Surveil(c *Conf) error {
	st, err := personState(c)
	if err != nil {
		return fmt.Errorf("failed to monitor person: ", err)
	}

	var m *gocv.Mat
	if st.Working {
		m = genScreen(st.Img)
	} else {
		m = genSlogan()
	}
	defer m.Close()
	img, err := m.ToImage()
	if err != nil {
		return fmt.Errorf("failed to convert Mat to Image: ", err)
	}
	return c.Slack.SetUserPhoto(img)
}
