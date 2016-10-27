package queue

type Queue struct {
	size    int      `队列大小`
	maxSize int      `队列最大值`
	queue   []string `队列数据`
}

func (this *Queue) Push(s string) {
	this.queue = append(this.queue, s)
	this.size = this.size + 1
}

//队列先进先出
func (this *Queue) Pop() string {
	s := this.queue[0:1]
	this.queue = this.queue[1:len(this.queue)]
	this.size = this.size - 1
	return s[0]
}

//队列先进后出
func (this *Queue) Pull() string {
	s := this.queue[len(this.queue)-1:]
	this.queue = this.queue[0 : len(this.queue)-1]
	this.size = this.size - 1
	return s[0]
}

func (this *Queue) GetSize() int {
	return this.size
}

func (this *Queue) GetMaxSize() int {
	return this.maxSize
}

func (this *Queue) SetMaxSize(maxSize int) {
	this.maxSize = maxSize
}

func (this *Queue) Get(index int) string {
	return this.queue[index-1 : index][0]
}
