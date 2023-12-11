package runner

import (
	"io"
	"os"
)

func RunPart(path string, part func(io.Reader) (int, error)) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	sum, err := part(file)
	if err != nil {
		return -1, nil
	}
	return sum, nil
}
