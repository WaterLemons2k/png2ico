package main

import (
	"fmt"
	"os"
)

type writer struct{ w *os.File }
type png struct {
	writer
	size uint32
}

func e(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func newWriter(file string) (w writer, err error) {
	w.w, err = os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0644)
	return
}
func (w writer) write(b []byte) {
	_, err := w.w.Write(b)
	e(err)
}

func newPNG(file string) (p png, err error) {
	p.w, err = os.Open(file)
	if err != nil {
		return
	}
	// ignore the error here, as we've already checked when we opening the file
	stat, _ := os.Stat(file)
	p.size = uint32(stat.Size())

	if !p.valid() {
		err = fmt.Errorf("%s: not a valid PNG", file)
	}
	return
}

// A valid PNG datastream shall begin with a [PNG signature],
// immediately followed by an [IHDR chunk].
//
// https://evanhahn.com/worlds-smallest-png/
//
// [PNG signature]: https://www.w3.org/TR/png/#3PNGsignature
// [IHDR chunk]: https://www.w3.org/TR/png/#11CcGen
func (p png) valid() bool {
	header := p.read(25)
	return string(header[:16]) == "\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR"
}

func (p png) read(size uint32) []byte {
	buf := make([]byte, size)
	_, err := p.w.ReadAt(buf, 0)
	e(err)
	return buf
}

// https://devblogs.microsoft.com/oldnewthing/20101018-00/?p=12513
func (p png) iconHeader(cnt byte) []byte {
	// ICONDIR
	off := byte(22)
	b := make([]byte, off)
	b[2], b[4] = 1, cnt

	// ICONDIRENTRY
	putUint32(b[14:], p.size)
	b[18] = off
	return b
}

func main() {
	if len(os.Args) != 2 {
		e(fmt.Errorf("usage: png2ico <png>"))
	}

	png, err := newPNG(os.Args[1])
	e(err)
	icoName := os.Args[1] + ".ico"
	ico, err := newWriter(icoName)
	e(err)

	ico.write(png.iconHeader(1))
	ico.write(png.read(png.size))
	fmt.Println("ICO:", icoName)
}
