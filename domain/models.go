package domain

const (
	Undefined_VehicleType    = ""
	Undefined_WorkflowFactor = ""
	Undefined_MovementOption = ""
)

type MovementActivity struct {
	Type   string
	Option string
}

type ContractCondition struct {
	Id                   string
	ContractorIdentifier string
	BranchIdentifier     string
	Name                 string
	VehicleType          string
	MovementActivities   []MovementActivity
	WorkflowType         string
	WorkflowFactor       string
}
