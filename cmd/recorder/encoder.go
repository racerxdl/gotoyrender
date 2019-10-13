package main

import (
	"bytes"
	"github.com/icza/mjpeg"
	"image"
	"image/jpeg"
)

type VideoEncoder struct {
	width       int
	height      int
	fps         int
	w           mjpeg.AviWriter
	frameNumber int
}

func MakeVideoEncoder(output string, width, height, fps int) *VideoEncoder {
	enc := &VideoEncoder{
		width:       width,
		height:      height,
		fps:         fps,
		frameNumber: 0,
	}

	aw, err := mjpeg.New(output, int32(width), int32(height), int32(fps))

	if err != nil {
		panic(err)
	}

	enc.w = aw

	return enc
}

func (ve *VideoEncoder) PutFrame(img image.Image) {
	buf := &bytes.Buffer{}
	err := jpeg.Encode(buf, img, nil)
	if err != nil {
		panic(err)
	}

	ve.frameNumber++

	ve.w.AddFrame(buf.Bytes())
}

func (ve *VideoEncoder) Close() {
	ve.w.Close()
}
