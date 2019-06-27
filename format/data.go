package format

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"image"
	"image/png"
	"strings"
)

func GetImageFromDataURL(data string) (image.Image, error) {
	if strings.HasPrefix(data, "data:image/png;base64,") {
		data := data[22:]
		dec, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			fmt.Println(err)
		}
		return png.Decode(bytes.NewReader(dec))
	}
	return nil, errors.New("unknown data URL")
}