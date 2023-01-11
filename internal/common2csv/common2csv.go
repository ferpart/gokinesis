package common2csv

import (
	"bytes"
	"os"

	"github.com/fpartidabc/gokinesis/internal/consume"
	"github.com/yukithm/json2csv"
)

const (
	trackedDir = "tracked"
	ext        = ".csv"
)

func Common2CSV(store *consume.CommonMap) error {
	if err := os.MkdirAll(trackedDir, os.ModePerm); err != nil {
		return err
	}
	for _, common := range *store {
		b := &bytes.Buffer{}
		wr := json2csv.NewCSVWriter(b)
		csv, err := json2csv.JSON2CSV(common)
		if err != nil {
			return err
		}
		if err = wr.WriteCSV(csv); err != nil {
			return err
		}
		wr.Flush()

		sessionID := common[0]["session_id"].(string)
		if err = os.WriteFile(trackedDir+"/"+sessionID+ext, b.Bytes(), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
