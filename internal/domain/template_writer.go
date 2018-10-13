package domain

type (
	// TplWriter abstracts to write generated code to file.
	TplWriter interface {
		Write(dst, tpl string, data interface{}) error
	}
)
