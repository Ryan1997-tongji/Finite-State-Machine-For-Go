package design

// Transition Define state transition rules
type Transition struct {
	EventType  string
	DstStateID *int64
	IsForce    bool
	DstState   State

	Callbacks  []Action
	Conditions []Condition
}

// TransitionConstructor Constructs for Transition
type TransitionConstructor struct {
	AllowedSourceStateIDs []int64 // not allow empty
	AllowedDstStateIDs    []int64 // empty means allow all
	IsForce               bool
	EventType             string
	DstState              State

	TransitionCallbacks  []Action
	GlobalCallbacks      []Action
	TransitionConditions []Condition
	GlobalConditions     []Condition
}
