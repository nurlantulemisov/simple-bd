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
func (q *Queue) Push(token query.Token) {
	q.dataQueue = append(q.dataQueue, &token)
	q.QSize++
	q.QStart = 0
}

// Pop remove
func (q *Queue) Pop() *query.Token {
	index := q.QStart
	q.QStart++
	return q.dataQueue[index]
}

// GetAll get all queue for debug
func (q *Queue) GetAll() []*query.Token {
	return q.dataQueue
}
