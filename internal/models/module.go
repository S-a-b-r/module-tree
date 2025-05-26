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

func (m *Module) GetWithSubstr(substr string) *Module {
	if !strings.Contains(m.Name, substr) {
		return nil
	}

	newM := Module{
		Name:         m.Name,
		Version:      m.Version,
		ChildModules: make([]*Module, 0),
	}

	for _, cm := range m.ChildModules {
		module := cm.GetWithSubstr(substr)
		if module != nil {
			newM.ChildModules = append(newM.ChildModules, module)
		}
	}

	return &newM
}

func (m *Module) DrawIONode(sb *strings.Builder, parentId int, nextId *int, x, y int) {
	currentId := *nextId
	*nextId++

	// Добавляем узел модуля
	sb.WriteString(fmt.Sprintf(`
		<mxCell id="%d" value="%s" style="rounded=1;whiteSpace=wrap;html=1;" parent="1" vertex="1">
			<mxGeometry x="%d" y="%d" width="120" height="60" as="geometry" />
		</mxCell>`,
		currentId, m.Name, x, y))

	// Добавляем связь с родителем
	if parentId != 1 {
		sb.WriteString(fmt.Sprintf(`
		<mxCell id="%d" value="%s" source="%d" target="%d" parent="1" edge="1">
			<mxGeometry relative="1" as="geometry" />
		</mxCell>`,
			*nextId, m.Version, parentId, currentId))
		*nextId++
	}

	// Отрисовываем дочерние элементы
	childX := x - 200 + len(m.ChildModules)*20
	childY := y + 100
	for i, dep := range m.ChildModules {
		dep.DrawIONode(sb, currentId, nextId, childX+i*150, childY)
	}
}

func parseModuleInfo(s string) (name, version string) {
	if strings.Contains(s, "@") {
		parts := strings.Split(s, "@")
		return parts[0], parts[1]
	}
	return s, "1.0.0" // default version
}
