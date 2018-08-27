package models

// Job queue interface
type Job interface {
	Perform()
}
