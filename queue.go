package gmq

import (
	"time"
)

type Queue struct {
	size  int
	queue []Message
}

type Message struct {
	Id      string
	Name    string
	Time    time.Time
	Message string
}

func (this *Queue) Push(s Message) {
	this.queue = append(this.queue, s)
	this.size = this.size + 1
}

func (this *Queue) PushNext(msg Message, index int) {
	var arry []Message = make([]Message, 0)
	arry = append(arry, msg)
	if this.size == 0 {
		this.queue = arry
		this.size = this.size + 1
		return
	}
	before := this.queue[:index+1]
	after := this.queue[index+1:]
	this.queue = append(before, append(arry, after...)...)
	this.size = this.size + 1
}

//队列先进先出
func (this *Queue) Pop() Message {
	s := this.queue[0:1]
	this.queue = this.queue[1:]
	this.size = this.size - 1
	// log.Println(s[0])
	return s[0]
}

//队列先进后出
func (this *Queue) Pull() Message {
	s := this.queue[this.size-1:]
	this.queue = this.queue[0 : this.size-1]
	this.size = this.size - 1
	return s[0]
}

func (this *Queue) Size() int {
	return this.size
}

func (this *Queue) Get(index int) Message {
	return this.queue[index : index+1][0]
}

func (this *Queue) List() []Message {
	return this.queue
}

func (this *Queue) FindById(id string) (index int, msg Message) {
	if this.size <= 0 {
		return -1, msg
	}
	for k, v := range this.queue {
		if v.Id == id {
			return k, this.queue[k : k+1][0]
		}
	}
	return -1, msg
}

func (this *Queue) FindByName(name string) (index int, msg Message) {
	if this.size <= 0 {
		return -1, msg
	}
	for k, v := range this.queue {
		if v.Name == name {
			return k, this.queue[k : k+1][0]
		}
	}
	return -1, msg
}

func (this *Queue) DeleteById(id string) int {
	if this.size <= 0 {
		return -1
	}
	var index, _ = this.FindById(id)
	if index < 0 {
		return -1
	}
	this.queue = append(this.queue[:index], this.queue[index+1:]...)
	this.size = this.size - 1
	return index
}
