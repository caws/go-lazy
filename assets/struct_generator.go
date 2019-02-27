package assets

import (
	"fmt"
	"os"
)

type StructGenerator struct {
	File *os.File
	Parameter Parameter
}

func (s *StructGenerator) GenerateClass(arguments []string) {
	s.Parameter.ComposeParameters(arguments)
	s.createFile()
	s.writePackage()
	s.writeStructSignature()
	s.writeStructAttributes()
}

func (s *StructGenerator) createFile() {
	f, err := os.Create(fmt.Sprintf("%s.go",s.Parameter.StructName))
	if err != nil {
		fmt.Println(err)
		return
	}

	s.File = f
}

func (s *StructGenerator) writePackage() {
	_, err := s.File.WriteString(fmt.Sprintf("package %s\n\n", s.Parameter.PackageName))
	s.someError(err)
}

func (s *StructGenerator) writeStructSignature() {
	_, err := s.File.WriteString(fmt.Sprintf("type %s struct {\n", s.Parameter.StructName))
	s.someError(err)
}

func (s *StructGenerator) writeStructAttributes() {
	content := s.Parameter.generateContent()
	_, err := s.File.WriteString(content)
	s.someError(err)
}

func (s *StructGenerator) someError(err error) {
	if err != nil {
		fmt.Println(err)
		s.File.Close()
		return
	}
}