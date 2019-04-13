package gojq

// Query ...
type Query struct {
	Pipe *Pipe `@@`
}

// Pipe ...
type Pipe struct {
	Terms []*Term `@@ ("|" @@)*`
}

// Term ...
type Term struct {
	ObjectIndex *ObjectIndex `@@ |`
	ArrayIndex  *ArrayIndex  `@@ |`
	Identity    *Identity    `@@`
}

// Identity ...
type Identity struct {
	_ string `"."`
}

// ObjectIndex ...
type ObjectIndex struct {
	Name     string `"." ( @Ident | "[" @String "]" )`
	Optional bool   `@"?"?`
}

// ArrayIndex ...
type ArrayIndex struct {
	Start *int `"." "[" ( @Integer?`
	End   *int `":" @Integer? |`
	Index *int `@Integer ) "]"`
}
