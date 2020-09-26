package util

import (
	"bytes"
	"io"

	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"golang.org/x/xerrors"
)

func ToUTF8(data []byte) (string, error) {
	dr, err := chardet.NewTextDetector().DetectBest([]byte(data))
	if err != nil {
		return "", xerrors.Errorf("detect charset error: %w", err)
	}
	var trans transform.Transformer
	switch dr.Charset {
	case "UTF-8", "ISO-8859-1":
		return string(data), nil
	case "Shift-JIS":
		trans = japanese.ShiftJIS.NewDecoder()
	case "EUC-JP":
		trans = japanese.EUCJP.NewDecoder()
	case "ISO-2022-JP":
		trans = japanese.ISO2022JP.NewDecoder()
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, transform.NewReader(bytes.NewReader(data), trans))
	return buf.String(), nil
}
