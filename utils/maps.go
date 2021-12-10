package utils

func GetKeys(values map[string]string) []string{
	var keys []string
	incr := 0
	for k := range values {
		keys[incr] = k
		incr++
	}
	return keys
}
