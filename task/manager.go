package task

import (
	"ant/core"
	"ant/core/err"
	topic2 "ant/core/topic"
	"ant/core/types"
	"ant/core/v1/action"
	"fmt"
	"github.com/goinggo/mapstructure"
	log "github.com/sirupsen/logrus"
)

type Manager struct {
	tasks  map[string]*TaskRuntime
	remote core.Remote
}

func (m *Manager) ExecuteCommand(resource, cmd string, payload interface{}) error {
	if cmd == action.StatusSetCommand {
		n := new(types.TaskStatusSetCommand)
		err := mapstructure.Decode(payload, n)
		if err != nil {
			return err
		}
		tasks := make([]*TaskRuntime, 0)
		for _, id := range n.Ids {
			if task, ok := m.tasks[id]; ok {
				tasks = append(tasks, task)
			}
		}
		if len(tasks) == 0 {
			return nil
		}
		status := n.Status
		if status == types.Running {
			for _, task := range tasks {
				task.Start()
			}
		} else if status == types.Stopped {
			for _, task := range tasks {
				task.Stop()
			}
		} else {
			return fmt.Errorf("unsppored status %s", status)
		}
		return nil
	} else {
		return fmt.Errorf("unsupported command %s", cmd)
	}
}

func (m *Manager) SetTasksStatus(newStatus interface{}) error {
	if status, ok := newStatus.(*types.TaskStatusSetCommand); !ok {
		return err.NewTypeNotExpectedError(newStatus)
	} else {
		tasks := make([]*TaskRuntime, 0)
		for _, id := range status.Ids {
			if task, ok := m.tasks[id]; ok {
				tasks = append(tasks, task)
			}
		}
		if len(tasks) == 0 {
			return nil
		}
		status := status.Status
		if status == types.Running {
			for _, task := range tasks {
				task.Start()
			}
		} else if status == types.Stopped {
			for _, task := range tasks {
				task.Stop()
			}
		} else {
			return fmt.Errorf("unsppored status %s", status)
		}
		return nil
	}
}

func (m *Manager) CreatOrUpdateTask(task types.Task) {
	id := task.Id
	if oldTask, ok := m.tasks[id]; ok {
		oldTask.Close()
	}
	if runtime, error := NewTask(task, m); error != nil {
		log.WithField("err", error.Error()).Warn("create task runtime failed")
	} else {
		runtime.Go()
		m.tasks[id] = runtime
	}
}

func (m *Manager) OnRow(row *core.Row) {
	ids := make([]string, 0, len(m.tasks))
	for _, task := range m.tasks {
		if task.Table != row.Meta.Name {
			continue
		}
		if shouldForward := task.ShouldForward(row); shouldForward {
			ids = append(ids, task.Id)
		} else {
			task.AcceptRow(row)
		}
	}

	if len(ids) > 0 {
		m.Forward2Queen(row, ids)
	}
}

func (m *Manager) Forward2Queen(row *core.Row, ids []string) {

}

func (m *Manager) CommitTaskResult(task *types.Task, rows []*core.Row) {
	topic := topic2.TaskResultReportTopic + task.Id
	m.remote.SendAsync(topic, &types.TaskResult{Id: task.Id, Rows: rows})
}

func (m *Manager) OnMsg(topic string, msg interface{}) error {
	if topic == topic2.TaskTopicStartFix {
		task := new(types.Task)
		error := mapstructure.Decode(msg, task)
		if error != nil {
			return error
		}
		m.CreatOrUpdateTask(*task)
		return nil
	} else {
		return fmt.Errorf("unsupported topic %s", topic)
	}
}

func NewTaskManager(remote core.Remote) *Manager {
	manager := &Manager{
		tasks:  make(map[string]*TaskRuntime, 8),
		remote: remote,
	}
	remote.RegisterConsumer(topic2.TaskTopicStartFix, "v1", manager.OnMsg)
	return manager
}
