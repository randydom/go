package impl

func Divide(f int32, s int32) (int32, int32){

	var d,r int32

	d = f / s
	r = f % s
	return d ,r

}
