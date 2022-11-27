package journal

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/viper"
	"pad/services/cache/internal/api/model"
)

type Journal interface {
	Start(ctx context.Context)
	Add([]model.Command)
	Restore() ([]model.Command, error)
}

type journal struct {
	queue chan []model.Command
	file  *os.File
}

func New() Journal {
	return &journal{
		queue: make(chan []model.Command, 50),
	}
}

func (j *journal) Start(ctx context.Context) {
	var (
		journalPath = viper.GetString("journal.path")
		err         error
	)
	if journalPath == "" {
		log.Println("Journal path is not defined. Using .journal.memstore")
		journalPath = ".journal.memstore"
	}

	j.file, err = os.OpenFile(journalPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Panicf("Could not start journal: %v", err)
	}

	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Panicf("Could not close journal file")
		}
	}(j.file)

	log.Println("Starting journal listener")
	for {
		select {
		case trns := <-j.queue: // Get transaction from queue
			for _, cmd := range trns {
				log.Println("Writing in journal")
				//if cmd.CmdType == "SET" {
				bin, _ := json.MarshalIndent(cmd, "", " ")
				if _, err := j.file.Write(bin); err != nil {
					log.Panicf("Journaling failed: %v", err)
				}
				//}
			}
		case <-ctx.Done():
			log.Println("Stopping queue listener")
			return
		}
	}
}

func (j *journal) Add(transaction []model.Command) {
	if j.file == nil {
		log.Println("Ignoring")
		return
	}

	j.queue <- transaction
}

func (j *journal) Restore() ([]model.Command, error) {
	if j.file != nil {
		log.Panic("Trying to restore while journal is already started")
	}
	journalPath := viper.GetString("journal.path")

	if journalPath == "" {
		log.Println("Journal path is not defined. Using .journal.memstore")
		journalPath = ".journal.memstore"
	}

	var err error
	j.file, err = os.OpenFile(journalPath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}

	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Panicf("Could not close journal file")
		}
	}(j.file)

	var buf []model.Command

	bytes, _ := ioutil.ReadAll(j.file)

	if err := json.Unmarshal(bytes, &buf); err != nil {
		return nil, err
	}
	return buf, nil
}
