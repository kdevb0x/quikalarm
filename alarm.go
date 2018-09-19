package main // import "github.com/kidoda/quikalarm"

import (
	"io"
	"time"
)

type OutputWriter interface {
	io.Writer
}
