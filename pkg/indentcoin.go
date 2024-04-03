package pkg

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"math/rand"

	svg "github.com/ajstarks/svgo"
	"github.com/lucasb-eyer/go-colorful"
)

type Options struct {
	Squares int
	Size    int
	Word    string
}

func getHashcode(str string) int64 {
	hashSlice := sha1.Sum([]byte(str))

	return int64(binary.BigEndian.Uint64(hashSlice[:]))
}

func getColor(hash int64) string {
	h := (float64(hash&0xFF) / 255.0) * 360
	s := 0.5
	v := 0.9

	return colorful.Hsv(h, s, v).Hex()
}

func GetIndentcoin(options Options) (*bytes.Buffer, error) {
	buffer := &bytes.Buffer{}

	svgSize := options.Squares * options.Size

	svgo := svg.New(buffer)

	svgo.Start(svgSize, svgSize, "shape-rendering=\"crispEdges\"")

	mirror := options.Squares / 2

	hash := getHashcode(options.Word)

	fillRect := "fill: " + getColor(hash)

	r := rand.New(rand.NewSource(hash))

	for i := 0; i < mirror; i++ {
		randomBoolVector := make([]bool, options.Squares)

		for i := 0; i < options.Squares; i++ {
			randomBoolVector[i] = r.Float32() > .5
		}

		if !randomBoolVector[len(randomBoolVector)-1] {
			randomBoolVector[len(randomBoolVector)-1] = (i % 2) == 0
		}

		if !randomBoolVector[0] {
			randomBoolVector[0] = (i % 2) != 0
		}

		for j := 0; j < options.Squares; j++ {
			if randomBoolVector[j] {
				xCoord := i * options.Size
				mirrorXCoord := (options.Squares - i - 1) * options.Size
				yCoord := j * options.Size

				svgo.Rect(xCoord, yCoord, options.Size, options.Size, fillRect)
				svgo.Rect(mirrorXCoord, yCoord, options.Size, options.Size, fillRect)
			}
		}
	}

	svgo.End()

	return buffer, nil
}
