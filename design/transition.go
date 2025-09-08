package design

// Transition Define state transition rules
// Represents a state transition rule in the state machine, defining conditions, triggering events, and execution logic for transitioning between states
type Transition struct {
	EventType  string // Identifier for the event type that triggers this transition
	DstStateID *int64 // Pointer to the target state's ID (may be nil for dynamically determined target state)
	IsForce    bool   // Whether to force transition (ignores some preconditions) reserved
	DstState   State  // Target state instance

	Callbacks  []Action    // List of callback actions to execute during transition
	Conditions []Condition // List of conditions that must all be satisfied to execute transition
}

// TransitionConstructor Constructor for Transition
// Provides flexible configuration for transition rules, including allowed states, events, conditions, and callbacks
type TransitionConstructor struct {
	AllowedSourceStateIDs []int64 // List of allowed source state IDs (must not be empty)
	AllowedDstStateIDs    []int64 // List of allowed target state IDs (empty means allow all)
	IsForce               bool    // Whether to force transition (same as Transition.IsForce)
	EventType             string  // Event type identifier that triggers transition
	DstState              State   // Target state instance

	TransitionCallbacks  []Action    // Callback actions specific to this transition
	GlobalCallbacks      []Action    // Globally applicable callback actions (shared across transitions)
	TransitionConditions []Condition // Conditions specific to this transition
	GlobalConditions     []Condition // Globally applicable conditions (shared across transitions)
}
