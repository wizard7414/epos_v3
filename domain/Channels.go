package domain

import "epos_v3/domain/data"

type AppChan struct {
	ProcessChan chan string

	SuccessChan chan string

	FailedChan chan string

	EntryChan chan data.Entry

	ObjectChan chan data.Object
}
