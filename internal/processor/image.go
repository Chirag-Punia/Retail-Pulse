package processor

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"net/http"
	"time"
)

type ImageProcessor struct{}

func NewImageProcessor() *ImageProcessor {
	return &ImageProcessor{}
}

func (p *ImageProcessor) ProcessImage(url string) (int, error) {

	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to download image: %v", err)
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to decode image: %v", err)
	}

	bounds := img.Bounds()
	perimeter := 2 * (bounds.Dx() + bounds.Dy())

	sleepTime := 100 + rand.Intn(301)
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)

	return perimeter, nil
}
