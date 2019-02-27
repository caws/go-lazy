package assets

import (
	"fmt"
	"strings"
)

type Attribute struct {
	AttributeName string
	AttributeType string
}

func (a *Attribute) composeAttributes(argument string) {
	attrAndType := strings.Split(argument, ":")
	a.AttributeName, a.AttributeType = attrAndType[0], attrAndType[1]
}

func (a *Attribute) generateThisLine(sizeAttributeName, sizeAttributeType int) string {
	var sb strings.Builder

	//Adding initial indentation
	sb.WriteString("\t")

	//Writing attribute name
	sb.WriteString(fmt.Sprintf("%s", a.AttributeName))
	sb.WriteString(" ")
	//Writing attribute indentation based on the biggest attribute name
	sb.WriteString(fmt.Sprintf("%s", a.generateIdentationName(sizeAttributeName)))

	//Writing attribute type
	sb.WriteString(fmt.Sprintf("%s", a.AttributeType))
	//Writing attribute type indentation based on the biggest attribute name
	sb.WriteString(fmt.Sprintf("%s", a.generateIdentationType(sizeAttributeType)))

	//Writing JSON stuff
	sb.WriteString(fmt.Sprintf("\t`json:\"%s\"`\n", a.AttributeName))

	return sb.String()
}

func (a *Attribute) generateIdentationName(sizeAttributeName int) string {
	return a.calculateIndentation(sizeAttributeName, len(a.AttributeName))
}

func (a *Attribute) generateIdentationType(sizeAttributeType int) string {
	return a.calculateIndentation(sizeAttributeType, len(a.AttributeType))
}

func (a *Attribute) calculateIndentation(longerAttribute, shorterAttribute int) string {
	var sb strings.Builder

	diff := longerAttribute - shorterAttribute

	for i := 1; i <= diff; i++ {
		sb.WriteString(" ")
	}

	return sb.String()
}
