package grammar

import "testing"

func TestEmptyVisitMap(t *testing.T) {
	var mapper = newVisitMap()
	var input = "not present"
	if expected, observed := false, mapper.HasKey(input); expected != observed {
		t.Fatalf("Expected %v, observed %v", expected, observed)
	}
}

func TestVisitMapSuccess(t *testing.T) {
	var mapper = newVisitMap()
	var key = "present"
	mapper.AddKey(key)
	if expected, observed := true, mapper.HasKey(key); expected != observed {
		t.Fatalf("Expected %v, observed %v", expected, observed)
	}
}
func TestVisitMapFail(t *testing.T) {
	var mapper = newVisitMap()
	var key = "present"
	var antikey = "not present"
	mapper.AddKey(key)
	if expected, observed := false, mapper.HasKey(antikey); expected != observed {
		t.Fatalf("Expected %v, observed %v", expected, observed)
	}
}

func TestStrIndexPresent(t *testing.T) {
	in := []string{"foo", "bar", "baz"}
	elem := "bar"
	if expected, observed := 1, strIndex(in, elem); expected != observed {
		t.Fatalf("Expected %v, observed %v", expected, observed)
	}
}

func TestStrIndexNotPresent(t *testing.T) {
	in := []string{"foo", "bar", "baz"}
	elem := "zillo"
	if expected, observed := -1, strIndex(in, elem); expected != observed {
		t.Fatalf("Expected %v, observed %v", expected, observed)
	}
}

func TestVistMap(t *testing.T) {
	var mapper = newVisitMap()
	var source, target = "good boy", "bad man"
	if expected, observed := false, mapper.HasVisited(source, target); expected != observed {
		t.Fatalf("Expected %v, observed %v", expected, observed)
	}
	mapper.Visit(source, target)
	if expected, observed := true, mapper.HasVisited(source, target); expected != observed {
		t.Fatalf("Expected %v, observed %v", expected, observed)
	}
}
