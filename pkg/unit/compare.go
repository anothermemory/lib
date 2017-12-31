package unit

// Equal compares two units and returns true if they are equal and false otherwise
func Equal(u1, u2 Unit) bool {
	if u1 == nil || u2 == nil {
		return false
	}
	if !equalUnit(u1, u2) {
		return false
	}

	switch u1.Type() {
	case TypeUnit:
		return true
	case TypeTextPlain:
		return equalTextPlain(u1.(TextPlain), u2.(TextPlain))
	case TypeTextMarkdown:
		return equalTextMarkdown(u1.(TextMarkdown), u2.(TextMarkdown))
	case TypeTextCode:
		return equalTextCode(u1.(TextCode), u2.(TextCode))
	case TypeTodo:
		return equalTodo(u1.(Todo), u2.(Todo))
	case TypeList:
		return equalList(u1.(List), u2.(List))
	}

	// Not supported units are not equal
	return false
}

// NotEqual two units and returns true if they are not equal and false otherwise
func NotEqual(u1, u2 Unit) bool {
	return !Equal(u1, u2)
}

func equalUnit(u1, u2 Unit) bool {
	if u1.ID() != u2.ID() {
		return false
	}
	if u1.Type() != u2.Type() {
		return false
	}
	if u1.Title() != u2.Title() {
		return false
	}
	if !u1.Created().Equal(u2.Created()) {
		return false
	}
	if !u1.Updated().Equal(u2.Updated()) {
		return false
	}
	return true
}
func equalTextPlain(u1, u2 TextPlain) bool {
	return u1.Data() == u2.Data()
}
func equalTextMarkdown(u1, u2 TextMarkdown) bool {
	return u1.Data() == u2.Data()
}
func equalTextCode(u1, u2 TextCode) bool {
	if u1.Language() != u2.Language() {
		return false
	}
	return u1.Data() == u2.Data()
}
func equalTodo(u1, u2 Todo) bool {
	if len(u1.Items()) != len(u2.Items()) {
		return false
	}

	for i := 0; i < len(u1.Items()); i++ {
		i1 := u1.GetItem(i)
		i2 := u2.GetItem(i)
		if i1.Done() != i2.Done() {
			return false
		}
		if i1.Data() != i2.Data() {
			return false
		}
	}

	return true
}
func equalList(u1, u2 List) bool {
	if len(u1.Items()) != len(u2.Items()) {
		return false
	}

	for i := 0; i < len(u1.Items()); i++ {
		i1 := u1.GetItem(i)
		i2 := u2.GetItem(i)

		if NotEqual(i1, i2) {
			return false
		}
	}

	return true
}
