package models

// Job queue interface
type Job interface {
	Exec()
}
