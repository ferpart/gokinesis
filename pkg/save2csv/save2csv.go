package common2csv

import (
	"bytes"
	"os"

	"github.com/yukithm/json2csv"
)

const (
	trackedDir = "tracked"
	ext        = ".csv"

	timeFmt = "YYYY-MM-DDThh:mm:ss"
)

func Save2CSV(items []map[string]interface{}) error {
	if err := os.MkdirAll(trackedDir, os.ModePerm); err != nil {
		return err
	}
	b := &bytes.Buffer{}
	wr := json2csv.NewCSVWriter(b)
	csv, err := json2csv.JSON2CSV(items)
	if err != nil {
		return err
	}
	if err = wr.WriteCSV(csv); err != nil {
		return err
	}
	wr.Flush()

	for _, record := range c.GetRecordMap() {
		b := &bytes.Buffer{}
		wr := json2csv.NewCSVWriter(b)
		csv, err := json2csv.JSON2CSV(record)
		if err != nil {
			return err
		}
		if err = wr.WriteCSV(csv); err != nil {
			return err
		}
		wr.Flush()

		recordKey := record[0][c.GetRecordKey()].(string)
		if err = os.WriteFile(trackedDir+"/"+recordKey+ext, b.Bytes(), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func getFileName() string {
}
