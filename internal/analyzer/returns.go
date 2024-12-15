package analyzer

import "strconv"

// To prevent other types being set to Returns.
type VaryLike interface {
	getAt(i int) []any
	asVary() Vary
}

// Defines possible return values for a specific function. Decision tree will explore all possible combination.
type Vary []VaryElement

func (r Vary) getAt(i int) []any {
	return r[i].get()
}

func (r Vary) asVary() Vary {
	return r
}

// To prevent other types being inserted as an element to Vary.
type VaryElement interface {
	get() []any
	passable() bool
}

// Defines return values for a specific function. Allows the decision tree to continue deciding paths.
type Pass []any

func (r Pass) getAt(i int) []any {
	if i != 0 {
		panic("unsupported index: " + strconv.Itoa(i))
	}
	return r
}

func (r Pass) get() []any {
	return r
}

func (r Pass) asVary() Vary {
	return Vary{r}
}

func (r Pass) passable() bool {
	return true
}

// Defines return values for a specific function. Prevents the decision tree to continue deciding paths.
type Stop []any

func (r Stop) getAt(i int) []any {
	if i != 0 {
		panic("unsupported index: " + strconv.Itoa(i))
	}
	return r
}

func (r Stop) get() []any {
	return r
}

func (r Stop) asVary() Vary {
	return Vary{r}
}

func (r Stop) passable() bool {
	return false
}
