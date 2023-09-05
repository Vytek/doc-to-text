package doctotext

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestDocToText(t *testing.T) {
	path := filepath.Join("testdata", "file-sample_100kB.doc")
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	text, err := DocToText(file, true)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(text)
}

func TestDoc2ToText(t *testing.T) {
	path := filepath.Join("testdata", "test_file_doc_1000.doc")
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	text, err := DocToText(file, false)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(text)
}
