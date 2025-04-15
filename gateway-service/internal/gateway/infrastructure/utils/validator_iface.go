package utils

type Validator interface {
	Struct(any) error
	Test(error) (map[string]string, error)
}
