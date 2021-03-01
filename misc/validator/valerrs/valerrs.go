package valerrs

import lib "github.com/go-playground/validator/v10"

//Fields extracts Field from the given ValidationErrors list.
func Fields(errs lib.ValidationErrors) []string {
	res := make([]string, len(errs))
	i := 0
	for _, err := range errs {
		res[i] = err.Field()
		i++
	}
	return res
}
