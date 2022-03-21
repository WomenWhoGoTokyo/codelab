package slender

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/image/draw"
)

type SlenderImage string

func (s SlenderImage) Path() (string, string, error) {
	path, err := filepath.Abs(filepath.Clean(string(s)))
	return path, filepath.Dir(path), err
}

func (s SlenderImage) Ext() string {
	return strings.ToLower(filepath.Ext(string(s)))
}

func (s SlenderImage) Name() string {
	file := string(s)
	return filepath.Base(file[:len(file)-len(s.Ext())])
}

func (s *SlenderImage) Make() error {
	// 3秒
	time.Sleep(3 * time.Second)

	path, dir, err := s.Path()
	if err != nil {
		return err
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	src, format, err := image.Decode(file)
	if err != nil {
		return err
	}

	rct := src.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, rct.Dx()/2, rct.Dy()))
	draw.CatmullRom.Scale(
		dst,
		dst.Bounds(),
		src,
		rct,
		draw.Over,
		nil,
	)

	slenderFile, err := os.Create(fmt.Sprintf("%s/%s-slender%s", dir, s.Name(), s.Ext()))
	if err != nil {
		return err
	}
	defer slenderFile.Close()

	switch format {
	case "jpeg":
		if err := jpeg.Encode(slenderFile, dst, &jpeg.Options{Quality: 100}); err != nil {
			return err
		}
	case "gif":
		if err := gif.Encode(slenderFile, dst, nil); err != nil {
			return err
		}
	case "png":
		if err := png.Encode(slenderFile, dst); err != nil {
			return err
		}
	default:
		return errors.New("Unsupported format")
	}
	return nil
}
