package routetemplate

import (
	_ "fmt"
	"regexp"
	"strings"
)

type RouteTemplate struct {
	initialPathChunk          string
	TemplatePath              string
	RegExPath                 string
	PathSegmentVariableNames  []string
	QuerySegmentVariableNames []string
	Values                    map[string]string
}

type RouteTemplateMatch struct {
	BoundVariables map[string]string
	TemplatePath   string
	CanidatePath   string
}

func Parse(rawurl string) (routeTemplate RouteTemplate, err error) {
	var pathTemplate = rawurl
	var qsTemplate = ""
	var firstPathChunk = rawurl

	if strings.Index(rawurl, "?") > -1 {
		var pathParts = strings.Split(rawurl, "?")
		pathTemplate = pathParts[0]
		qsTemplate = pathParts[1]
	}
	var firstSquiggly = strings.Index(pathTemplate, "{")
	if firstSquiggly > -1 {
		firstPathChunk = pathTemplate[0:firstSquiggly]
	}

	re := regexp.MustCompile("{(.*?)}")
	var qsSegments = re.FindAllString(qsTemplate, -1)
	for i := 0; i < len(qsSegments); i++ {
		qsSegments[i] = strings.Trim(qsSegments[i], "{}")
	}

	var pathSegments = re.FindAllString(pathTemplate, -1)
	for i := 0; i < len(pathSegments); i++ {
		pathSegments[i] = strings.Trim(pathSegments[i], "{}")
	}

	for _, value := range strings.Split(pathTemplate, "/") {
		if strings.Index(value, "{") > -1 {
			value = strings.Trim(value, "{}")
		}
	}

	routeTemplate = RouteTemplate{}
	routeTemplate.initialPathChunk = firstPathChunk
	routeTemplate.TemplatePath = rawurl
	routeTemplate.RegExPath = convertToRegexp(rawurl)
	routeTemplate.PathSegmentVariableNames = pathSegments
	routeTemplate.QuerySegmentVariableNames = qsSegments
	return routeTemplate, nil
}

func IsMatch(candidate string, routeTemplate RouteTemplate) (matched bool, err error) {

	if strings.HasPrefix(candidate, routeTemplate.initialPathChunk) {
		re := regexp.MustCompile(routeTemplate.RegExPath)
		matched = false
		if re.FindAllString(candidate, -1) != nil {
			matched = true
		}
		return matched, nil
	} else {
		return false, nil
	}

}

func BindVariables(candidate string, routeTemplate RouteTemplate) (matched RouteTemplateMatch, err error) {
	var routeTemplateMatch = RouteTemplateMatch{}
	routeTemplateMatch.CanidatePath = candidate
	routeTemplateMatch.TemplatePath = routeTemplate.RegExPath

	routeTemplateMatch.BoundVariables = make(map[string]string)
	//if match, _ := IsMatch(candidate, routeTemplate); match == true {
	re := regexp.MustCompile(routeTemplate.RegExPath)

	var matchedValues = re.FindStringSubmatch(candidate)
	if matchedValues != nil {
		matchedValues = matchedValues[1:len(matchedValues)]
		for i := 0; i < len(matchedValues); i++ {
			routeTemplateMatch.BoundVariables[routeTemplate.PathSegmentVariableNames[i]] = matchedValues[i]
		}
	}

	//}
	return routeTemplateMatch, nil
}

func convertToRegexp(templatedItem string) (converted string) {
	re := regexp.MustCompile("{.*?}")
	converted = re.ReplaceAllStringFunc(templatedItem, replaceWithRegex)
	return
}
func replaceWithRegex(s string) string {
	var template = strings.TrimLeft(s, "{")
	template = strings.TrimRight(template, "}")

	var expressionType = ""
	if strings.Contains(template, ":") {
		expressionType = strings.Split(template, ":")[1]
	}

	switch expressionType {
	case "numeric":
		return "([0-9]+)"
	case "alpha":
		return "([A-Za-z]+)"
	case "alphanumeric":
		return "([A-Za-z0-9]+)"
	case "guid":
		return "(\\{?[a-fA-F0-9]{8}(?:-(?:[a-fA-F0-9]){4}){3}-[a-fA-F0-9]{12}\\}?)"

	}

	return "([^/]+)"
}
