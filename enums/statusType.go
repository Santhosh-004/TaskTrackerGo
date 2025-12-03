package enums

type StatusType int

const (
	ToDo StatusType = iota
	Done
)

var statusName = map[StatusType]string {
	ToDo: "To-Do",
	Done: "Done",
}

func (st StatusType) ToString() string {
	return statusName[st];
}