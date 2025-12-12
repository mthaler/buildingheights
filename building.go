package main

type Building struct {
	Name               string
	Type               string
	YearOfConstruction string
	Height             string
	Latitude           string
	Longitude          string
	Remark             string
}

func (b Building) toSlice() []string {
	var result []string
	result = make([]string, 0)
	result = append(result, b.Name)
	result = append(result, b.Type)
	result = append(result, b.YearOfConstruction)
	result = append(result, b.Height)
	result = append(result, b.Latitude)
	result = append(result, b.Longitude)
	result = append(result, b.Remark)
	return result
}
