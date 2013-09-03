package routetemplate //same package name as source file

import (
	"fmt"
	"testing"
)

func Test_RegExPath_Numeric(t *testing.T) { //test function starts with "Test" and takes a pointer to type testing.T
	var shouldMatch = "/test/([0-9]+)/whats/([^/]+)"
	if url, err := Parse("/test/{id:numeric}/whats/{what}"); err != nil {
		t.Error("Test_Numeric did not work as expected.")
	} else {
		if url.RegExPath != shouldMatch {
			t.Error("Test_Numeric did not work as expected.")
		}
	}

}

func Test_RegExPath_Default(t *testing.T) { //test function starts with "Test" and takes a pointer to type testing.T
	var shouldMatch = "/test/([^/]+)/whats/([^/]+)"
	if url, err := Parse("/test/{id}/whats/{what}"); err != nil {
		t.Error("Test_Default did not work as expected.")
	} else {
		if url.RegExPath != shouldMatch {
			t.Error("Test_Default did not work as expected.")
		}
	}

}

func Test_RegExPath_Alpha(t *testing.T) { //test function starts with "Test" and takes a pointer to type testing.T
	var shouldMatch = "/test/([A-Za-z]+)/whats/([^/]+)"
	if url, err := Parse("/test/{id:alpha}/whats/{what}"); err != nil {
		t.Error("Test_Alpha did not work as expected.")
	} else {
		if url.RegExPath != shouldMatch {
			t.Error("Test_Alpha did not work as expected.")
		}
	}

}

func Test_RegExPath_AlphaNumeric(t *testing.T) { //test function starts with "Test" and takes a pointer to type testing.T
	var shouldMatch = "/test/([A-Za-z0-9]+)/whats/([^/]+)"

	if url, err := Parse("/test/{id:alphanumeric}/whats/{what}"); err != nil {
		t.Error("Test_AlphaNumeric did return the correct expression.")
	} else {
		if url.RegExPath != shouldMatch {
			t.Error("Test_AlphaNumeric did not work as expected.")
		}
	}

}

func Test_PathSegmentParsing(t *testing.T) { //test function starts with "Test" and takes a pointer to type testing.T
	var template1 = "/test1/blah/{id:numeric}/whats/{what}"

	if url, err := Parse(template1); err != nil {
		t.Error("Test_PathSegmentParsing did not work as expected.")
	} else {
		fmt.Printf("%q\n", url.PathSegmentVariableNames[1])
		if url.PathSegmentVariableNames[1] != "what" {
			t.Error("Test_PathSegmentParsing did not work as expected.")
		}
	}
	ClearAllTemplates()

}

func Test_Add_Routes(t *testing.T) { //test function starts with "Test" and takes a pointer to type testing.T
	var template1 = "/test1/blah/{id:numeric}/whats/{what}"
	var template1a = "/test1/blah/{id}/whats/{what}"
	var template2 = "/test2/blah/{id}/whats/{what}"
	var template3 = "/test3/blah/{id}/whats/{what}"
	var template3a = "/test3/blah/{id}/whats/{what:numeric}"
	var template4 = "/test4/blah/{id:alphanumeric}/whats/{what}"

	Add(template1)
	Add(template1a)
	Add(template2)
	Add(template3)
	Add(template3a)
	Add(template4)

	if templates, err := GetAllTemplates(); err != nil {
		t.Error("Test_Add_Routes did return the correct expression.")
	} else {
		if len(templates) != 6 {
			t.Error("Test_Add_Routes did not work as expected.")
		}
	}
	ClearAllTemplates()

}

func Test_Match_Template(t *testing.T) { //test function starts with "Test" and takes a pointer to type testing.T
	var template1 = "/test1/blah/{id:numeric}/whats/{what}"
	var template1a = "/test1/blah/{id}/whats/{what}"
	var template2 = "/test2/blah/{id}/whats/{what}"
	var template3 = "/test3/blah/{id}/whats/{what}"
	var template3a = "/test3/blah/{id}/whats3a/{what:numeric}"
	var template4 = "/test4/blah/{id:alphanumeric}/whats/{what}"

	var testUrl = "/test3/blah/1324/whats3a/12344"

	Add(template1)
	Add(template1a)
	Add(template2)
	Add(template3)
	Add(template3a)
	Add(template4)

	if template, err := GetMatchedTemplateString(testUrl); err != nil {
		t.Error("Test_Match_Template did return the correct expression.")
	} else {
		//fmt.Printf("%q\n", template)

		if template != template3a {
			t.Error("Test_Match_Template did not work as expected.")
		}
	}
	ClearAllTemplates()
}

func Test_VariableBinding(t *testing.T) { //test function starts with "Test" and takes a pointer to type testing.T
	var template1 = "/test1/blah/{id:numeric}/whats/{what}"
	var testUrl = "/test1/blah/1324/whats/12344"

	if url, err := Parse(template1); err != nil {
		t.Error("Parse call in Test_VariableBinding did not work as expected.")
	} else {
		if templateMatch, err := BindVariables(testUrl, url); err != nil {
			t.Error("BindVariables call in Test_VariableBinding did not work as expected.")
		} else {
			fmt.Printf("%q\n", templateMatch.BoundVariables)
			if templateMatch.BoundVariables["what"] != "12344" {
				t.Error("Test_VariableBinding did not work as expected.")
			}
		}
	}
}

type TestController struct {
}

func (testController TestController) SomeActionHandler() {

}
func Test_AddRoute(t *testing.T) {
	var testController = TestController{}
	AddRoute("name", "method", "urlTemplate", testController.SomeActionHandler)
}

func Test_MatchTemplateObject(t *testing.T) { //test function starts with "Test" and takes a pointer to type testing.T
	var template1 = "/test1/blah/{id:numeric}/whats/{what}"
	var template1a = "/test1/blah/{id}/whats/{what}"
	var template2 = "/test2/blah/{id}/whats/{what}"
	var template3 = "/test3/blah/{id}/whats/{what}"
	var template3a = "/test3/blah/{id}/whats3a/{what:numeric}"
	var template4 = "/test4/blah/{id:alphanumeric}/whats/{what}"

	var testUrl = "/test3/blah/1324/whats3a/12344"

	Add(template1)
	Add(template1a)
	Add(template2)
	Add(template3)
	Add(template3a)
	Add(template4)

	if matchTemplateObject, err := GetMatchTemplate(testUrl); err != nil {
		t.Error("Test_MatchTemplateObject did return the correct expression.")
	} else {
		//fmt.Printf("%q\n", template)
		fmt.Printf("%q\n", matchTemplateObject.TemplatePath)

		if matchTemplateObject.TemplatePath != template3a {
			t.Error("Test_MatchTemplateObject did not work as expected.")
		}
	}
	ClearAllTemplates()
}

// unless you can pass a pointer to a function in an Interface type
//routes.MapRoute("some other name", "GET", "users/{id}", &structWithFunctions, "functionName")
//routes.MapHttpRoute("some name", "routeTemplate/{routeTemplate}/etc/{etc}", struct with functions? or function, since function signature needs to match)
//
