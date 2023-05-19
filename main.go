package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Sum(fileName string) (ret int64, _ error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return 0, err
	}

	for _, line := range bytes.Split(b, []byte("\n")) {
		num, err := strconv.ParseInt(string(line), 10, 64)
		if err != nil {
			return 0, err
		}

		ret += num
	}

	return ret, nil
}

func SumByteScanner(fileName string) (ret int64, _ error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var fbuf [8 * 1024]byte
	var acc int64
	for {
		n, err := f.Read(fbuf[:])
		if err != nil && err != io.EOF {
			return 0, err
		}

		if n == 0 {
			ret += acc
			break
		}

		for _, c := range fbuf[:n] {
			if c == '\n' {
				ret += acc
				acc = 0
				continue
			}

			if c < '0' || c > '9' {
				return 0, fmt.Errorf("invalid character: %c", c)
			}

			acc = acc*10 + int64(c-'0')
		}
	}
	return ret, nil
}

func SumStream(fileName string) (ret int64, _ error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	for {
		var num int64
		_, err := fmt.Fscanf(f, "%d\n", &num)
		if err != nil {
			if err == io.EOF {
				return ret, nil
			}

			return 0, err
		}

		ret += num
	}
}

func GenFile(fileName string, n int) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	for i := 0; i < n; i++ {
		if _, err := f.WriteString("1\n"); err != nil {
			return err
		}
	}

	return nil
}
