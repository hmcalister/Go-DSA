package linkedlistqueue

import "errors"

var (
	ErrorQueueEmpty   = errors.New("queue is empty")
	ErrorItemNotFound = errors.New("item not found in queue")
)
