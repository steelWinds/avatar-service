package internal

import (
	"bytes"
	"hash/fnv"
	"math/rand"

	svgo "github.com/ajstarks/svgo"
	colorful "github.com/lucasb-eyer/go-colorful"
)

type Options struct {
	Squares int
	Size    int
	Word    string
}

func getHashcode(str string) (int64, error) {
	fnvHash := fnv.New32()

	_, err := fnvHash.Write([]byte(str))

	if err != nil {
		return 0, err
	}

	return int64(fnvHash.Sum32()), nil
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

	svgo := svgo.New(buffer)

	svgo.Start(svgSize, svgSize)

	mirror := options.Squares / 2

	hash, err := getHashcode(options.Word)

	fillRect := "fill: " + getColor(hash)

	if err != nil {
		return nil, err
	}

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
