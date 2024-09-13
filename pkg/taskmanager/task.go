package taskmanager

type Task struct {
	option string
	value  string
	action func(...string)
}

// getters and setters for Task type
func (t *Task) SetOption(opt string) {
	t.option = opt
}

func (t *Task) SetValue(val string) {
	t.value = val
}

func (t *Task) SetAction(action func(...string)) {
	t.action = action
}

func (t *Task) GetOption() string {
	return t.option
}

func (t *Task) GetValue() string {
	return t.value
}

func (t *Task) GetAction() func(...string) {
	return t.action
}
