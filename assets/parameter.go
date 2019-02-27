package assets

import (
	"strings"
)

type Parameter struct {
	PackageName       string
	StructName        string
	Attributes        []Attribute
	SizeAttributeName int
	SizeAttributeType int
}

func (p *Parameter) ComposeParameters(arguments []string) {
	p.SizeAttributeName = 0
	p.PackageName = arguments[1]
	p.StructName = arguments[2]
	p.ComposeAttributes(arguments[3:])
}

func (p *Parameter) ComposeAttributes(arguments []string) {
	for _, arg := range arguments {
		var attribute Attribute
		attribute.composeAttributes(arg)
		p.checkLongestAttribute(arg)
		p.Attributes = append(p.Attributes, attribute)
	}
}

func (p *Parameter) checkLongestAttribute(attribute string) {
	attrAndType := strings.Split(attribute, ":")
	p.checkLongestAttributeName(attrAndType[0])
	p.checkLongestAttributeType(attrAndType[1])
}

func (p *Parameter) checkLongestAttributeName(attributeName string) {
	if len(attributeName) > p.SizeAttributeName {
		p.SizeAttributeName = len(attributeName)
	}
}

func (p *Parameter) checkLongestAttributeType(attributeType string) {
	if len(attributeType) > p.SizeAttributeType {
		p.SizeAttributeType = len(attributeType)
	}
}

func (p *Parameter) generateContent() string {
	var sb strings.Builder

	for _, attribute := range p.Attributes {
		sb.WriteString(attribute.generateThisLine(p.SizeAttributeName, p.SizeAttributeType))
	}
	sb.WriteString("}")

	return sb.String()
}