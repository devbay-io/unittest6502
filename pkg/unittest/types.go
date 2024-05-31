package unittest

type UnitTestAssertionCondition struct {
	Name  string `yaml:"name" validate:"required"`
	Value string `yaml:"value" validate:"required"`
}

type UnitTestAssertion struct {
	Type      string                     `yaml:"type" validate:"required"`
	Location  string                     `yaml:"location" validate:"required"`
	Condition UnitTestAssertionCondition `yaml:"condition" validate:"required"`
}

type UnitTest struct {
	Platform       string              `yaml:"platform" validate:"required,oneof=simple nes"`
	ProgramCounter uint                `yaml:"program_counter" validate:"required"`
	Assert         []UnitTestAssertion `yaml:"assert" validate:"required"`
}
