package domain

import "log"

type AppLogger struct {
	GeneralLogger *log.Logger
	ErrorLogger   *log.Logger
}
