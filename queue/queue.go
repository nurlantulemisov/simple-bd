package queue

import (
	"query"
)

// Queue for parcing
type Queue struct {
	QSize     int64
	QStart    int64
	dataQueue []*query.Token
}

// Push add
func (q *Queue) Push(token *query.Token) {
	q.dataQueue = append(q.dataQueue, token)
	q.QSize++
	q.QStart = 0
}

// Pop remove
func (q *Queue) Pop() *query.Token {
	index := q.QStart
	q.QStart++
	if index <= int64(len(q.dataQueue)) {
		elementRemove := q.dataQueue[index]
		q.dataQueue = q.dataQueue[index:]
		return elementRemove
	}
	return nil
}

// GetAll get all queue for debug
func (q *Queue) GetAll() []*query.Token {
	return q.dataQueue
}
