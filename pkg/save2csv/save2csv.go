package save2csv

import (
	"bytes"
	"os"
	"time"

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

	if err = os.WriteFile(getFileName(), b.Bytes(), os.ModePerm); err != nil {
		return err
	}

	return nil
}

func getFileName() string {
	t := time.Now()
	return trackedDir + "/" + t.Format(timeFmt) + ext
}
