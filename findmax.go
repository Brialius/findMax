package findMax

import "errors"

func FindMax(comp func(el1, el2 interface{}) bool, elements []interface{}) (i interface{}, e error) {
	if len(elements) < 1 {
		return nil, errors.New("not enough parameters")
	}
	max := elements[0]
	defer func() {
		if r := recover(); r != nil {
			i = nil
			e = r.(error)
			return
		}
	}()
	for i, element := range elements {
		if i > 0 && comp(element, max) {
			max = element
		}
	}
	return max, nil
}
