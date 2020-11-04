package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m1)
	t.Log(m1["two"])
	m2 := map[string]int{}
	m2["one"] = 1
	m2["two"] = 2
	t.Log(m2)
	m3 := make(map[string]int, 10)
	t.Logf("len m3 = %d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m := map[int]int{}
	t.Log(m[1])
	m[2] = 0
	t.Log(m[2])
	m[3] = 0
	if val, ok := m[3]; ok {
		t.Logf("存在key为3的值为: %d", val)
	} else {
		t.Log("map中不存在key为3的")
	}

}

func TestTravelMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}

	for key, val := range m1 {
		t.Logf("key: %v, val: %d", key, val)
	}
}
