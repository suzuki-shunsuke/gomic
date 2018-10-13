package domain

type (
	// MockTemplateArg is a template argument and represents a mock.
	MockTemplateArg interface {
		Version() string
		URL() string
		PackageName() string
		MockName() string
		Imports() [][]string
		Methods() []Method
	}

	// Method represents a mock's method.
	Method interface {
		Name() string
		Declaration() string
		Definition() string
		ParamsStr() string
		ResultValuesStr() string
		HasResultNames() bool
		Imports() map[string]ImportSpec
		Results() []Var
		IsEllipsis() bool
	}

	// Var represents a method's parameter and return value.
	Var interface {
		Name() string
		Type() string
		PkgName() string
		PkgPath() string
	}

	// ImportSpec represents an import statement.
	ImportSpec interface {
		Host() string
		Name() string
		Path() string
		String() string
	}

	// ImportSpecs represents a collection of ImportSpec.
	ImportSpecs interface {
		Add(ImportSpec) (ImportSpec, error)
		Names() map[string]ImportSpec
		Paths() map[string]ImportSpec
	}
)
