package command

type ForgetScheduledWorkloadParams struct {
}

type ForgetScheduledWorkload struct {
}

func NewForgetScheduleWorkload() *ForgetScheduledWorkload {
	return &ForgetScheduledWorkload{}
}

func (cmd *ForgetScheduledWorkload) Invoke(params ForgetScheduledWorkloadParams) error {
	panic("not implemented")
}
