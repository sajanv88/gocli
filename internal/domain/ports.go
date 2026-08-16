package domain

type RouterGenerator interface {
	Name() RouterOption
	Generate(spec ProjectSpec) error
}

type DBGenerator interface {
	Name() DatabaseOption
	Generate(spec ProjectSpec) error
}

type FrontendGenerator interface {
	Name() FrontendOption
	Generate(spec ProjectSpec) error
}

type BackendGenerator interface {
	Name() AgentOption
	Generate(spec ProjectSpec) error
}
