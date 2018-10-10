package Calculator


type (

	MyCalculator interface {
		Operate() (int32, int32)
	}

	NumberSet struct {
		First int32
		Second int32
		Outcome int32
		Remainder int32
	}

)




