package models

import (
	"fmt"
	"strings"
)

type Module struct {
	Name         string
	Version      string
	ChildModules []*Module
}

func (m *Module) AddModule(parentName, childName, childVersion string) bool {
	if m.Name == parentName {
		m.ChildModules = append(m.ChildModules, &Module{
			Name:    childName,
			Version: childVersion,
		})
		return true
	}

	for _, child := range m.ChildModules {
		if child.AddModule(parentName, childName, childVersion) {
			return true
		}
	}

	return false
}

func (m *Module) String(depth int, prefix string) string {
	var sb strings.Builder

	// Current module
	sb.WriteString(prefix)
	if depth > 0 {
		sb.WriteString("├── ")
	}
	sb.WriteString(fmt.Sprintf("%s@%s\n", m.Name, m.Version))

	// Child modules
	for i, child := range m.ChildModules {
		newPrefix := prefix
		if depth > 0 {
			if i == len(m.ChildModules) {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}
		}
		sb.WriteString(child.String(depth+1, newPrefix))
	}

	return sb.String()
}

func parseModuleInfo(s string) (name, version string) {
	if strings.Contains(s, "@") {
		parts := strings.Split(s, "@")
		return parts[0], parts[1]
	}
	return s, "1.0.0" // default version
}

func (m *Module) GetWithDepth(depth int) *Module {
	if depth == 0 {
		return nil
	}

	newM := Module{
		Name:         m.Name,
		Version:      m.Version,
		ChildModules: make([]*Module, 0),
	}

	for _, cm := range m.ChildModules {
		module := cm.GetWithDepth(depth - 1)
		if module != nil {
			newM.ChildModules = append(newM.ChildModules, module)
		}
	}
	return &newM
}
