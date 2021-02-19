package command

type SetReplicaCountParams struct {

}

type SetReplicaCount struct {

}

func NewSetReplicaCount() *SetReplicaCount {
	return &SetReplicaCount{

	}
}

func (cmd *SetReplicaCount) Invoke(params SetReplicaCountParams) error {
	panic("not implemented");
}
