package nerosql

import "testing"

func TestLexCreateModel(t *testing.T) {
	source := "create model iris_knn as knn(sepallength, sepalwidth, petallength, petalwidth, species, 3, euclidean) from iris;"
	cursor := cursor{pointer: 0, loc: Location{Line: 1, Col: 1}}
	expectedToken := &Token{
		Value: "create model iris_knn as knn(sepallength, sepalwidth, petallength, petalwidth, species, 3, euclidean) from iris;",
		Kind:  KeywordKind,
		Loc:   Location{Line: 1, Col: 1},
	}

	token, newCursor, ok := lexCreateModel(source, cursor)
	if !ok {
		t.Errorf("lexCreateModel returned false, expected true")
	}
	if token.Value != expectedToken.Value {
		t.Errorf("lexCreateModel returned token value %q, expected %q", token.Value, expectedToken.Value)
	}
	if token.Kind != expectedToken.Kind {
		t.Errorf("lexCreateModel returned token kind %v, expected %v", token.Kind, expectedToken.Kind)
	}
	if token.Loc != expectedToken.Loc {
		t.Errorf("lexCreateModel returned token location %v, expected %v", token.Loc, expectedToken.Loc)
	}
	if newCursor.pointer != uint(len(source)) {
		t.Errorf("lexCreateModel returned new cursor pointer %v, expected %v", newCursor.pointer, len(source))
	}
}
