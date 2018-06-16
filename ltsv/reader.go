package ltsv

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Reader struct {
	reader *bufio.Reader
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		reader: bufio.NewReader(r),
	}
}

func (r *Reader) Read() (map[string]interface{}, error) {
	record := make(map[string]interface{})

	for {
		line, _, err := r.reader.ReadLine()
		if err != nil {
			return nil, err
		}

		strLine := strings.TrimSpace(string(line))
		splitedLine := strings.Split(strLine, "\t")

		for _, field := range splitedLine {
			if field == "" {
				continue
			}
			row := strings.SplitN(field, ":", 2)
			if len(row) != 2 {
				return record, fmt.Errorf("unexpected label name %s", row[0])
			}
			record[row[0]] = parseValue(string(row[1]))
		}
		return record, nil
	}
}

func (r *Reader) ReadAll() ([]map[string]interface{}, error) {
	records := []map[string]interface{}{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			return records, nil
		}
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
}

func parseValue(val string) interface{} {
	var (
		v   interface{}
		err error
	)

	v, err = strconv.Atoi(val)
	if err == nil {
		return v
	}

	v, err = strconv.ParseFloat(val, 64)
	if err == nil {
		return v
	}

	v, err = strconv.ParseBool(val)
	if err == nil {
		return v
	}

	return val
}
