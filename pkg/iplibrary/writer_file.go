// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://cdn.foyeseo.com .

package iplibrary

import (
	"compress/gzip"
	"os"
)

type FileWriter struct {
	fp       *os.File
	gzWriter *gzip.Writer
	password string

	rawWriter WriterInterface
}

func NewFileWriter(path string, meta *Meta, password string) (*FileWriter, error) {
	fp, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	gzWriter, err := gzip.NewWriterLevel(fp, gzip.BestCompression)
	if err != nil {
		return nil, err
	}

	var writer = &FileWriter{
		fp:        fp,
		gzWriter:  gzWriter,
		rawWriter: NewWriterV1(gzWriter, meta),
		password:  password,
	}
	return writer, nil
}

func (this *FileWriter) WriteMeta() error {
	return this.rawWriter.WriteMeta()
}

func (this *FileWriter) Write(ipFrom string, ipTo string, countryId int64, provinceId int64, cityId int64, townId int64, providerId int64) error {
	return this.rawWriter.Write(ipFrom, ipTo, countryId, provinceId, cityId, townId, providerId)
}

func (this *FileWriter) Sum() string {
	return this.rawWriter.Sum()
}

func (this *FileWriter) Close() error {
	err1 := this.gzWriter.Close()
	err2 := this.fp.Close()
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	// 加密内容
	if len(this.password) > 0 {
		var filePath = this.fp.Name()
		data, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		if len(data) > 0 {
			encodedData, encodeErr := NewEncrypt().Encode(data, this.password)
			if encodeErr != nil {
				return encodeErr
			}
			_ = os.Remove(filePath)
			err = os.WriteFile(filePath, encodedData, 0666)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
