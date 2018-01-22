package grammar

// A visitMap is a map from a starting production to
// the list of all productions it has visited.
type visitMap map[string][]string

// NewVisitMap returns an empty map
func newVisitMap() visitMap {
	return make(visitMap)
}

// HasKey returns true if the key is in the map
func (mapper visitMap) HasKey(key string) bool {
	_, ok := mapper[key]
	return ok
}

// AddKey adds the key to the map, overwriting an existing structure
func (mapper visitMap) AddKey(key string) {
	mapper[key] = []string{}
}

func (mapper visitMap) HasVisited(source, target string) bool {
	if !mapper.HasKey(source) {
		return false
	}
	reached := mapper[source]
	return strIndex(reached, target) != -1
}

func (mapper visitMap) Visit(source, target string) {
	if mapper.HasVisited(source, target) {
		return
	}
	if !mapper.HasKey(source) {
		mapper.AddKey(source)
	}
	mapper[source] = append(mapper[source], target)
}

func strIndex(array []string, elem string) int {
	for i, val := range array {
		if val == elem {
			return i
		}
	}
	return -1
}
