package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"testing"
)

func getFS() fs.FS {
	return os.DirFS("tests")
}

func isInputFilename(filename string) bool {
	return !strings.HasSuffix(filename, ".a")
}

func getAnswerFilename(inputFilename string) string {
	return fmt.Sprintf("%s.a", inputFilename)
}

func Test_run(t *testing.T) {
	tests := loadTests(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.Close()

			wantOut, err := tt.WantOut()
			if err != nil {
				t.Fatal(err)
			}

			out := &bytes.Buffer{}
			run(tt.in, out)
			gotOut := out.String()

			if gotOut != wantOut {
				t.Errorf("got\t%#v\nexpected\t%#v", gotOut, wantOut)
			}
		})
	}
}

type testCase struct {
	name    string
	in      fs.File
	wantOut fs.File
}

func (c *testCase) WantOut() (string, error) {
	wantOutBuf := &bytes.Buffer{}
	_, err := wantOutBuf.ReadFrom(c.wantOut)
	if err != nil {
		return "", err
	}
	return wantOutBuf.String(), nil
}

func (c *testCase) Close() {
	err := c.in.Close()
	err2 := c.wantOut.Close()
	if err != nil {
		fmt.Printf("failed closing: %v", err)
	}
	if err2 != nil {
		fmt.Printf("failed closing: %v", err)
	}
}

func loadTests(t *testing.T) []*testCase {
	var tests []*testCase
	fsys := getFS()
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walking dir: %w", err)
		}
		if !d.IsDir() && isInputFilename(d.Name()) {
			test, err := loadTest(path, d, fsys)
			if err != nil {
				return fmt.Errorf("loading test: %w", err)
			}

			tests = append(tests, test)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return tests
}

func loadTest(path string, d fs.DirEntry, fsys fs.FS) (*testCase, error) {
	inputFile, err := fsys.Open(path)
	if err != nil {
		return nil, err
	}

	outputFile, err := fsys.Open(getAnswerFilename(path))
	if err != nil {
		return nil, err
	}

	return &testCase{
		name:    d.Name(),
		in:      inputFile,
		wantOut: outputFile,
	}, nil
}
