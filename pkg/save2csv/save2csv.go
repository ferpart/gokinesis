package save2csv

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/yukithm/json2csv"
)

const (
	trackedDir = "tracked"

	fileFmt = trackedDir + "/%s_%s.csv"

	timeFmt = "15:04:05"
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
	return fmt.Sprintf(fileFmt, uuid.New(), t.Format(timeFmt))
}
