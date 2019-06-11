package findMax

func FindMax(comp func(el1, el2 interface{}) bool, elements ...interface{}) (i interface{}, e error) {
	max := elements[0]
	defer func() {
		if r := recover(); r != nil {
			i = nil
			e = r.(error)
			return
		}
	}()
	for _, element := range elements {
		if comp(element, max) {
			max = element
		}
	}
	return max, nil
}
