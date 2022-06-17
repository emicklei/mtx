package core

type Named interface{ Name() string }

func FindByName[T Named](elements []T, name string) (T, bool) {
	for _, each := range elements {
		if each.Name() == name {
			return each, true
		}
	}
	var t T
	return t, false
}
