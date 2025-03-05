package constant

type AppointmentStatus string

const (
	ToDo       AppointmentStatus = "To Do"
	InProgress AppointmentStatus = "In Progress"
	Done       AppointmentStatus = "Done"
)

func AppointmentStatuses() []AppointmentStatus {
	return []AppointmentStatus{ToDo, InProgress, Done}
}

func (status AppointmentStatus) IsValid() bool {
	switch status {
	case ToDo, InProgress, Done:
		return true
	}
	return false
}

func (status AppointmentStatus) String() string {
	return string(status)
}
