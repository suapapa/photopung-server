package main

import (
	"bytes"
	"context"
	"image"
	"image/jpeg"
	"io"
	"log"
	"mime/multipart"
	"time"

	"github.com/nfnt/resize"
	"github.com/pkg/errors"
)

var (
	imgCache = make(map[string]*PTPImage)
)

func cacheImageFromFile(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", errors.Wrap(err, "failed to cache image")
	}
	defer src.Close()

	img, err := imgFromReader(src)
	if err != nil {
		return "", errors.Wrap(err, "failed to cache image")
	}

	// resize image into 800x800 box
	if img.Bounds().Dx() > 800 || img.Bounds().Dy() > 800 {
		img = resize.Resize(800, 800, img, resize.Lanczos3)
	}
	ptpi := NewPTPImage(file.Filename, imgToBytes(img))
	imgCache[ptpi.Name] = ptpi

	return ptpi.Name, nil
}

func cleanGarbageImageCache(ctx context.Context) {
	tkr := time.NewTicker(1 * time.Minute)
	for {
		select {
		case <-ctx.Done():
			return
		case <-tkr.C:
			for k, v := range imgCache {
				// TODO: lifetime should be configurable
				if time.Since(v.BornAt) > 10*time.Minute {
					log.Printf("delete image from cache: %s", k)
					delete(imgCache, k)
				}
			}
		}
	}
}

func imgFromReader(r io.Reader) (image.Image, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func imgToBytes(img image.Image) []byte {
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)
	return buf.Bytes()
}
