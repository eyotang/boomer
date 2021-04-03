package boomer

var (
	StopUser = &stopUser{}
)

type stopUser struct{}

func (su *stopUser) Error() string {
	return "StopUser"
}
