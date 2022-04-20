package outpututil

import (
	"sync"
)

type ResultListener struct {
	resultChannel chan ResultList
	doneChannel   chan struct{}
	initOnce      sync.Once
	printer       ResultPrinter
}

func (l *ResultListener) Init() {
	l.resultChannel = make(chan ResultList, 1)
	l.doneChannel = make(chan struct{})
}

func (l *ResultListener) GetResultChannel() chan<- ResultList {
	return l.resultChannel
}

func (l *ResultListener) Done() {
	l.doneChannel <- struct{}{}
}

func (l *ResultListener) Listen() {
	l.printer.initFile()
	for {
		select {
		case resultList := <-l.resultChannel:
			l.printer.PrintResultList(resultList)
		case <-l.doneChannel:
			close(l.resultChannel)
			close(l.doneChannel)
			l.printer.Close()
			return
		}
	}
}
