package impl

import "awesomeProject/GraphTheory/Graph"

type (
	BasicGraph struct {
		Elements []Graph.Element
	}
)

func NewBasicGraph(el *[]Graph.Element) Graph.MyElement{
	return &BasicGraph{Elements: *el}
}

func (b *BasicGraph) Create(n string, l string) {
	id := len(b.Elements) + 1
	// Creating a new element and add it to the Element list
	ele := Graph.Element{
		Id: id,
		Name: n,
		Layer: l,
	}

	b.Elements = append(b.Elements, ele)
}

func (b *BasicGraph) Add(el Graph.Element, id int, r string) {
	// I am just adding the relationship to the element
	el.Relations = append(el.Relations, Graph.Relationships{
		Id: id,
		Relation: r,
	})

}