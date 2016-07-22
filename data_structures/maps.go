package data_structures

func findPeopleWithCommonInterest(data map[string][]string, interest string) []string {
	
	result := []string{}	//declare slice

	for idx,val := range data {
		if contains(val, interest) {
			result = append(result, idx)
		}
	}
	return result
}

func contains(src []string, elem string) bool {
	
	var res bool = false
	for _,value := range src {
		if value==elem {
			res = true
		} 
	}
	return res
}
