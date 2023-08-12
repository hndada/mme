package mme

import (
	"reflect"
)

// Todo:Element -> Block, Box
type Block = Element

type Condition struct{}

func Or()
func And()

// These methods are for applying style for condition-meeting elements.
type Selector struct {
	Type     SelectorType
	Relation Relation
	Names    []string
	// Blocks   []string // names
	// Fields   []string // names
}

type SelectorType int

const (
	SelectorTypeBlock SelectorType = iota
	SelectorTypeLabel
	SelectorTypeID
	SelectorTypeField   // Attributes
	SelectorTypeElement // Pseudo-element
	SelectorTypeEvent   // Pseudo-class
)

type Relation int

const (
	RelationChild      Relation = iota // >
	RelationDescendant                 // " " (space)
	RelationSibling                    // ~
	RelationAdjacent                   // +
)

func blockNames(bs ...Block) []string {
	names := make([]string, len(bs))
	for i, b := range bs {
		name := reflect.TypeOf(b).Name()
		names[i] = name
	}
	return names
}

// func Selector(r Relation, bs ...Block) string {
// 	var op string
// 	switch r {
// 	case RelationDescendant:
// 		op = " "
// 	case RelationChild:
// 		op = " > "
// 	case RelationSibling:
// 		op = " ~ "
// 	case RelationAdjacent:
// 		op = " + "
// 	}
// 	names := blockNames(bs...)
// 	return strings.Join(names, op)
// }
