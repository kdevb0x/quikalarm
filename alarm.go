package main // import "github.com/kidoda/quikalarm"

import (
	"io"
)

type OutputWriter interface {
	io.Writer
}
