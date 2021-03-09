package valerrs

import lib "github.com/go-playground/validator/v10"

//Collect extracts data for pairs of invalid field and validation from the given ValidationErrors list.
func Collect(errs lib.ValidationErrors) map[string]string {
	res := make(map[string]string, len(errs))
	for _, err := range errs {
		res[err.Field()] = err.Tag()
	}
	return res
}
