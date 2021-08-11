package zstd

import (
	"bytes"
	"crypto/sha1"
	"os"
	"testing"

	"github.com/test-go/testify/require"
)

func TestCompressDecompress(t *testing.T) {
	type testCase struct {
		name   string
		input  string
		inFile string
	}

	testcases := []*testCase{
		{
			name:  "simple input",
			input: "How much wood could a woodchuck chuck if a woodchuck could chuck wood ?",
		},
		{
			name:   "zstd test case #414",
			inFile: "testdata/failer.dat",
		},
	}

	for _, tc := range testcases {
		var in []byte

		if tc.inFile != "" {
			var err error
			in, err = os.ReadFile(tc.inFile)
			require.NoError(t, err, tc.name)

		} else {
			in = []byte(tc.input)
		}

		var b bytes.Buffer
		wc := Compressor{}.NewWriter(&b)

		var err error
		_, err = wc.Write(in)
		require.NoError(t, err, tc.name)

		err = wc.Close()
		require.NoError(t, err, tc.name)

		var result bytes.Buffer
		err = Decompressor{}.Decompress(&result, &b)
		require.NoError(t, err, tc.name)

		// check that sha1 of result matches sha1 of input :
		got := sha1.Sum(result.Bytes())
		expected := sha1.Sum(in)

		require.Equal(t, expected, got, tc.name)
	}
}
