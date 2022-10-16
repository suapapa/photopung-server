package main

import "time"

type PTPImage struct {
	OrigName string
	Name     string
	Data     []byte
	BornAt   time.Time
}

func NewPTPImage(origName string, data []byte) *PTPImage {
	return &PTPImage{
		OrigName: origName,
		Name:     randSeq(8),
		Data:     data,
		BornAt:   time.Now(),
	}
}
