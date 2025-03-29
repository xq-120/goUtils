package tests

import (
	"fmt"
	"testing"

	xqkit "github.com/xq-120/goUtils/pkg"
)

func TestSortedMap(t *testing.T) {
	sm := xqkit.NewSortedMap()
	sm.SetObjectForKey("obj2", "2")
	sm.SetObjectForKey("obj1", "1")
	sm.SetObjectForKey("obj3", "3")
	for item := range sm.Iterator() {
		fmt.Printf("key:%s, val:%s\n", item.Key, item.Val)
	}

	if v, ok := sm.ObjectForKey("2"); ok {
		if v != "obj2" {
			t.Errorf("没找到")
		}
	}
}
