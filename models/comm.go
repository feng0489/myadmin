package models

func InArr(val string,arr []string) bool {
	var ok bool=false
	for _,v :=range arr {
		if v == val {
			ok=true
		}
	}
	return ok
}