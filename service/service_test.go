package service

import (
	"testing"
)

type MockDb struct {
	value map[string]int
}

func NewMockDb() *MockDb {
	return &MockDb{}
}

func (db *MockDb) Get(key string) (any, error) {
	return db.value[key], nil
}

func (db *MockDb) Set(key string, value any) error {
	v := value.(int)
	db.value[key] = v
	return nil
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestGetAge(t *testing.T) {
	db := NewMockDb()
	db.value = make(map[string]int)
	db.value["age"] = 2
	svc := NewService(db)
	v, err := svc.GetAge()
	if v != 2 {
		t.Fatalf("Expect %d but receive %d\n", db.value["age"], v)
	}
	if err != nil {
		t.Fatalf("GetAge got error!\n")
	}
}

func TestSetAge(t *testing.T) {
	db := NewMockDb()
	db.value = make(map[string]int)
	svc := NewService(db)
	var v int = 15
	err := svc.SetAge(v)
	if err != nil {
		t.Fatalf("SetAge returns error!\n")
	}
}

func TestSetGetAge(t *testing.T) {
	db := NewMockDb()
	db.value = make(map[string]int)
	svc := NewService(db)
	var ageSet int = 15
	err := svc.SetAge(ageSet)
	if err != nil {
		t.Fatalf("SetAge returns error!\n")
	}
	ageGet, err := svc.GetAge()
	if ageGet != ageSet {
		t.Fatalf("Expect %d but receive %d\n", ageGet, ageSet)
	}
	if err != nil {
		t.Fatalf("GetAge got error!\n")
	}
}
