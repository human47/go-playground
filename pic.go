package main

import (
	"bytes"
	"image"
	"image/png"
	"math/cmplx"
	"os"
)

func ShowImage(m image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(buf.Bytes())
}

func Pic() {
	const (
		dx       = 1024
		dy       = 682
		max_iter = 4096
	)
	yScale := 3. / float64(dy)
	xScale := 2. / float64(dx)

	IterMap := make([][]uint32, dy)

	for yi := range IterMap {
		IterMap[yi] = make([]uint32, dx)

		for xi := range IterMap[yi] {
			iteration := uint32(0)
			c := complex(float64(xi)*xScale-1.5, float64(yi)*yScale-1.)
			z := 0 + 0i
			for ; InSet(z) && iteration < max_iter; iteration++ {
				z = z*z + c
			}
			IterMap[yi][xi] = iteration
		}
	}

	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := IterMap[y][x]
			i := y*m.Stride + x*4

			// If they don't escape, use black for "in the set"
			if v == max_iter {
				v = 0
			}

			m.Pix[i] = uint8(v << 4 & 0xff)
			m.Pix[i+1] = uint8(v & 0xf0)
			m.Pix[i+2] = uint8(v >> 8 & 0xf0)
			m.Pix[i+3] = 255
		}
	}
	ShowImage(m)
}

func InSet(z complex128) bool {
	r, _ := cmplx.Polar(z)
	return r < 4
}

func main() {
	Pic()
}

/* ft=go */
