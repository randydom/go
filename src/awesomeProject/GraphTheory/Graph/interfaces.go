package Graph

type (

	Relationships struct{
		Id int
		Relation string
	}

	Element struct{
		Id int
		Name string
		Layer string
		Relations []Relationships
	}

	MyElement interface{
		Create (string, string)
		Add (Element, int, string)
}

)


