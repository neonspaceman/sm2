package flow

type Builder struct {
	flow *Flow
}

func NewBuilder() *Builder {
	return &Builder{
		flow: &Flow{
			activeAction: make(map[string]FlowFunction),
		},
	}
}

func (b *Builder) Before(f FlowFunction) *Builder {
	b.flow.beforeEachAction = f

	return b
}

func (b *Builder) After(f FlowFunction) *Builder {
	b.flow.afterEachAction = f

	return b
}

func (b *Builder) Step(stepName string, f FlowFunction) *Builder {
	b.flow.activeAction[stepName] = f

	return b
}

func (b *Builder) UnhandledState(f FlowFunction) *Builder {
	b.flow.unhandledState = f

	return b
}

func (b *Builder) Flow() *Flow {
	return b.flow
}
