package domain

type RouterGenerator interface {
	Name() string
	Generate(spec ProjectSpec) error
}

type DBGenerator interface {
	Name() string
	Generate(spec ProjectSpec) error
}

type FrontendGenerator interface {
	Name() string
	Generate(spec ProjectSpec) error
}
