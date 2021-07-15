package main

import "fmt"

func main() {
	t := template{ID: 3,
		OriginCompany:      "Linux",
		DestinationCompany: "Windows"}

	var ts templates = []template{
		{ID: 1,
			OriginCompany:      "Nike",
			DestinationCompany: "Adidas"},
		{ID: 2,
			OriginCompany:      "Nintendo",
			DestinationCompany: "Sony"},
		{ID: 3,
			OriginCompany:      "Linux",
			DestinationCompany: "Windows"},
	}

	ft := getTemplate(ts, t)
	fmt.Println(ft)
}

type template struct {
	ID                 int
	OriginCompany      string
	DestinationCompany string
}

type templates []template

type templatesFilter func(template) bool

func getTemplate(ts templates, t template) templates {

	var fnFilter templatesFilter = func(match template) bool {
		if t.OriginCompany != match.OriginCompany {
			return false
		}
		if t.DestinationCompany != match.DestinationCompany {
			return false
		}
		return true
	}

	return templates.Filter(ts, fnFilter)
}

func (ts templates) Filter(filter templatesFilter) templates {
	filtered := templates{}
	for _, t := range ts {
		if !filter(t) {
			continue
		}
		filtered = append(filtered, t)
	}
	return filtered
}
