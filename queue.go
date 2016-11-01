package queue

import (
	"fmt"
	"time"
)

type Queue struct {
	size  int
	queue []Message
}

type Message struct {
	Id      string
	Time    time.Time
	Message string
}

func (this *Queue) Push(s Message) {
	this.queue = append(this.queue, s)
	this.size = this.size + 1
}

func (this *Queue) PushNext(s []Message, index int) {
	if this.size == 0 {
		this.queue = s
		this.size = this.size + 1
		return
	}
	before := this.queue[:index+1]
	after := this.queue[index+1:]
	fmt.Println(this.size, after)
	this.queue = append(before, append(s, after...)...)
	this.size = this.size + 1
}

//队列先进先出
func (this *Queue) Pop() Message {
	s := this.queue[0:1]
	this.queue = this.queue[1:this.size]
	this.size = this.size - 1
	//	logger.Debug(s[0].Message)
	return s[0]
}

//队列先进后出
func (this *Queue) Pull() Message {
	s := this.queue[this.size-1:]
	this.queue = this.queue[0 : this.size-1]
	this.size = this.size - 1
	return s[0]
}

func (this *Queue) GetSize() int {
	return this.size
}

func (this *Queue) Get(index int) Message {
	return this.queue[index : index+1][0]
}

func (this *Queue) List() []Message {
	return this.queue
}

func (this *Queue) FindById(id string) []int {
	var find []int
	if this.size <= 0 {
		return find
	}
	for i, _ := range this.queue {
		if msg := this.queue[i : i+1][0]; id == msg.Id {
			find = append(find, i)
		}
	}
	return find
}

func (this *Queue) DelById(id string) bool {
	if this.size <= 0 {
		return false
	}
	var findIndex []int = this.FindById(id)
	if len(findIndex) <= 0 {
		return false
	}

	var firstFind int = findIndex[0]
	this.queue = append(this.queue[:firstFind], this.queue[firstFind+1:]...)
	this.size = this.size - 1
	return true
}
