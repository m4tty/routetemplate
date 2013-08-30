package routetemplate

import (
	_ "fmt"
)

var routeTemplates []RouteTemplate

func Add(template string) (err error) {
	routeTemplate, _ := Parse(template)
	routeTemplates = append(routeTemplates, routeTemplate)
	return nil
}

func GetMatchedTemplateString(url string) (template string, err error) {
	template = ""

	for _, value := range routeTemplates {

		wasMatched, _ := IsMatch(url, value)
		if wasMatched {
			template = value.TemplatePath
			break
		}
	}
	return template, nil
}

func GetMatchTemplate(url string) (routeTemplateMatch RouteTemplateMatch, err error) {

	routeTemplateMatch = RouteTemplateMatch{}
	for _, value := range routeTemplates {
		wasMatched, _ := IsMatch(url, value)
		if wasMatched {

			routeTemplateMatch, _ = BindVariables(url, value)
			break
		}
	}
	return routeTemplateMatch, nil
}

func GetMatchedTemplate(url string) (template string, err error) {
	template = ""

	for _, value := range routeTemplates {

		wasMatched, _ := IsMatch(url, value)
		if wasMatched {
			template = value.TemplatePath
			break
		}
	}
	return template, nil
}

func GetAllTemplates() (templates []RouteTemplate, err error) {
	return routeTemplates, nil
}

func ClearAllTemplates() (err error) {
	routeTemplates = make([]RouteTemplate, 0)
	return nil
}
