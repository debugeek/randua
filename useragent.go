package randua

import (
	"strings"
)

type UserAgentBuilder struct {
	Elements []UserAgentElement
}

func NewUserAgentBuilder() UserAgentBuilder {
	return UserAgentBuilder{
		Elements: make([]UserAgentElement, 0),
	}
}

func (b *UserAgentBuilder) AddElement(element UserAgentElement) {
	b.Elements = append(b.Elements, element)
}

func (b UserAgentBuilder) Build() string {
	var components []string
	for _, element := range b.Elements {
		var component = element.Name
		if len(element.Comments) > 0 {
			component += " (" + strings.Join(element.Comments, "; ") + ")"
		}
		components = append(components, component)
	}
	return strings.Join(components, " ")
}

type UserAgentElement struct {
	Name     string
	Comments []string
}

func NewUserAgentElement(name string) UserAgentElement {
	return UserAgentElement{
		Name: name,
	}
}

func (e *UserAgentElement) AddComment(comment string) {
	e.Comments = append(e.Comments, comment)
}
