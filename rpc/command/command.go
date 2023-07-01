package command

import "log"

type TaskServer bool

type Task uint8

// const (
// 	Load Task = iota
// 	Start
// 	Stop
// 	Remove
// )

type TaskData struct {
	UserId, ProjectName string
}

func (t *TaskServer) Load(data *TaskData, reply *bool) error {
	logger := log.Default()
	logger.Println("Loading new resource...")
	*reply = true
	return nil
}

func (t *TaskServer) Start(data *TaskData, reply *bool) error {
	logger := log.Default()
	logger.Println("Starting resource...")
	*reply = true
	return nil
}

func (t *TaskServer) Stop(data *TaskData, reply *bool) error {
	logger := log.Default()
	logger.Println("Stopping new resource...")
	*reply = true
	return nil
}

func (t *TaskServer) Remove(data *TaskData, reply *bool) error {
	logger := log.Default()
	logger.Println("Removing new resource...")
	*reply = true
	return nil
}
