package outpututil

import (
	"sync"
)

type ResultListener struct {
	resultChannel chan ResultList
	doneChannel   chan struct{}
	initOnce      sync.Once
	printer       ResultPrinter
	wg            sync.WaitGroup
}

func (l *ResultListener) Init() {
	l.resultChannel = make(chan ResultList, 3)
	l.doneChannel = make(chan struct{})
}

func (l *ResultListener) GetResultChannel() chan<- ResultList {
	return l.resultChannel
}

func (l *ResultListener) Done() {
	l.doneChannel <- struct{}{}
}

func (l *ResultListener) WaitForClose() {
	l.wg.Wait()
}

func (l *ResultListener) Listen() {
	l.printer.initFile()
	l.wg.Add(1)
	for {
		select {
		case resultList := <-l.resultChannel:
			l.printer.PrintResultList(resultList)
		case <-l.doneChannel:
			close(l.resultChannel)
			l.printer.Close()
			l.wg.Done()
			return
		}
	}
}
