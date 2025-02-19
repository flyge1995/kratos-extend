package crawler

type Metadata struct {
	TaskID      int    // 任务ID
	TaskName    string // 任务名称
	DetailState int    // 1. 暂无详情需要抓取  2.已有详情
	//URLString   string   // 任务URL
	//URL         *url.URL // init自动解析
	Crontab string // 抓取周期
}

type Context struct {
}

type Task interface {
	Metadata() Metadata
	OnHandler(ctx *Context) error
}

func NewTaskDecorator(task Task) (*TaskDecorator, error) {
	metadata := task.Metadata()
	//if metadata.URL == nil {
	//	parse, err := url.Parse(metadata.URLString)
	//	if err != nil {
	//		return nil, err
	//	}
	//	metadata.URL = parse
	//}
	return &TaskDecorator{task: task, metadata: metadata}, nil
}

var _ Task = new(TaskDecorator)

type TaskDecorator struct {
	task     Task
	metadata Metadata
}

func (t *TaskDecorator) Metadata() Metadata {
	return t.metadata
}

func (t *TaskDecorator) OnHandler(ctx *Context) error {
	return t.task.OnHandler(ctx)
}

//var ErrorImplementMe = errors.New("implement me")
