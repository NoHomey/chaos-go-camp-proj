package dict

//Keys return the list of dict map keys.
func Keys(m map[string]string) []string {
	keys := make([]string, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}
	return keys
}

//Data return lists of keys and coresponding values.
func Data(m map[string]string) ([]string, []string) {
	keys := make([]string, len(m))
	vals := make([]string, len(m))
	i := 0
	for key, val := range m {
		keys[i] = key
		vals[i] = val
		i++
	}
	return keys, vals
}
