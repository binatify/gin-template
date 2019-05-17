package runmode

type RunMode string

const (
	Production  = RunMode("production")
	Development = RunMode("development")
	Test        = RunMode("test")
)

func (mode RunMode) IsValid() bool {
	switch mode {
	case Production, Development, Test:
		return true
	default:
		return false
	}
}

func (mode RunMode) IsProduction() bool {
	return mode == Production
}

func (mode RunMode) IsDevelopment() bool {
	return mode == Development
}

func (mode RunMode) IsTest() bool {
	return mode == Test
}
