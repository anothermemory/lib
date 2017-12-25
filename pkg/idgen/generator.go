package idgen

// Generator represents interface which can be used to generate random ID values. Actual generated values depends
// on used implementation
type Generator interface {
	Generate() string
}
