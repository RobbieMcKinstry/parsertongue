package grammar

import "math"

// JSON is the struct which returns JSON
// representation of the *grammar.G
type JSON struct {
	Root        string        `json:"root"`
	Productions []*production `json:"grammar"`
}

// contains the information about a single production
type production struct {
	Name     string
	Children []string
	Level    int
}

// NewJSON returns a new JSON Obj
func NewJSON(gram *G) *JSON {
	jsonObj := &JSON{
		Root:        gram.sentenceRoot,
		Productions: []*production{},
	}
	for name, prod := range gram.gram {
		data := new(production)
		data.Name = name
		data.Children = Children(prod)
		jsonObj.Productions = append(jsonObj.Productions, data)
	}
	jsonObj.setLevels(gram)
	return jsonObj
}

// NewJSONFile builds the file from the JSON
func NewJSONFile(filename, root string) *JSON {
	g := New(filename, root)
	return NewJSON(g)
}

// this runs a topological sort on the productionData
// it sets the level of how deep each node should be in the graph
// calculating the level as the level of the deepest parent + 1
func (obj *JSON) setLevels(g *G) {
	var (
		root          = obj.Root
		rootIndex     = index(root, obj.Productions)
		level         = 1
		cycleDetector = newVisitMap()
	)
	setLevel(obj.Productions, rootIndex, level)
	obj.visitChildren(root, cycleDetector, g)
}

func (obj *JSON) visitChildren(parent string, cycleDetector visitMap, g *G) {
	children := Children(g.Prod(parent))
	for _, child := range children {
		if cycleDetector.HasVisited(parent, child) {
			continue
		}
		obj.setChildLevel(child, parent, cycleDetector)
		cycleDetector.Visit(parent, child)
		obj.visitChildren(child, cycleDetector, g)
	}
}

func (obj *JSON) setChildLevel(child, parent string, cycleDetector visitMap) {
	if cycleDetector.HasVisited(parent, child) {
		return
	}
	childData := obj.Productions[index(child, obj.Productions)]
	parentData := obj.Productions[index(parent, obj.Productions)]
	parentLvl := float64(parentData.Level)
	childLvl := float64(childData.Level)
	childData.Level = int(math.Max(parentLvl+1.0, childLvl))
}

func index(name string, productions []*production) int {
	for i, elem := range productions {
		if elem.Name == name {
			return i
		}
	}
	return -1
}

func setLevel(prods []*production, index, level int) {
	if index < 0 || index >= len(prods) {
		panic("index out of bounds")
	}
	prods[index].Level = level
}
