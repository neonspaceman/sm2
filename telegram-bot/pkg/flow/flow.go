package flow

import (
	"context"
)

const ContextStateKey = "bot_handler/state-key"

type FlowFunction func(ctx context.Context, step, prevStep string, args []any)

type FlowState struct {
	Args        []any
	Step        string
	PrevStep    string
	StepChanged bool
}

type Flow struct {
	beforeEachAction FlowFunction
	activeAction     map[string]FlowFunction
	afterEachAction  FlowFunction
	unhandledState   FlowFunction
}

func (f *Flow) Active(ctx context.Context, step string, args ...any) *FlowState {
	state := &FlowState{
		Args:        args,
		Step:        step,
		PrevStep:    "",
		StepChanged: false,
	}

	ctx = context.WithValue(ctx, ContextStateKey, state)

	if f.beforeEachAction != nil {
		f.beforeEachAction(ctx, state.Step, state.PrevStep, state.Args)
	}

	if !state.StepChanged {
		f.callActiveAction(ctx, state)
	}

	if f.afterEachAction != nil {
		f.afterEachAction(ctx, state.Step, state.PrevStep, state.Args)
	}

	return state
}

func (f *Flow) Goto(ctx context.Context, gotoStep string) *FlowState {
	state := f.GetState(ctx)
	state.PrevStep = state.Step
	state.StepChanged = true
	state.Step = gotoStep

	f.callActiveAction(ctx, state)

	return state
}

func (f *Flow) GetState(ctx context.Context) *FlowState {
	return ctx.Value(ContextStateKey).(*FlowState)
}

func (f *Flow) callActiveAction(ctx context.Context, state *FlowState) {
	if action, ok := f.activeAction[state.Step]; ok {
		action(ctx, state.Step, state.PrevStep, state.Args)
	} else if f.unhandledState != nil {
		f.unhandledState(ctx, state.Step, state.PrevStep, state.Args)
	}
}
