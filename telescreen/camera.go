package telescreen

import (
	"fmt"
	"image"
	"sort"

	"gocv.io/x/gocv"
)

type state struct {
	Working bool
	Img     *gocv.Mat
}

func personState(c *Conf) (*state, error) {
	cap, err := capture(c.DeviceID)
	defer cap.Close()
	if err != nil {
		return nil, err
	}

	faces, err := detectFaces(&cap, c.CascadeClassifierPath)
	if err != nil {
		return nil, err
	}

	if len(faces) == 0 {
		return &state{Working: false}, nil
	}

	// Choose the most human face.
	sort.Slice(faces, func(i, j int) bool {
		return (faces[i].Dx() + faces[i].Dy()) > (faces[j].Dx() + faces[j].Dy())
	})

	cropped := cap.Region(faces[0])
	img := cropped.Clone()

	return &state{
		Working: true,
		Img:     &img,
	}, nil
}

func capture(devideID int) (gocv.Mat, error) {
	img := gocv.NewMat()

	webcam, err := gocv.OpenVideoCapture(1)
	if err != nil {
		return img, err
	}
	defer webcam.Close()

	if ok := webcam.Read(&img); !ok {
		return img, fmt.Errorf("cannot read device %v", devideID)
	}

	if img.Empty() {
		return img, fmt.Errorf("the surveilled person violated")
	}

	return img, nil
}

func detectFaces(img *gocv.Mat, cpath string) ([]image.Rectangle, error) {
	c := gocv.NewCascadeClassifier()
	defer c.Close()

	if !c.Load(cpath) {
		return []image.Rectangle{}, fmt.Errorf("error reading cascade file: %s", cpath)
	}

	faces := c.DetectMultiScale(*img)
	return faces, nil
}
