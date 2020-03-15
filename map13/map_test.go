package map13

import "testing"

func TestMap(t *testing.T) {
	mymap := map[string]string{}
	mymap2 := make(map[string]string)
	mymap2["name"] = "张三"
	mymap2["age"] = "14"
	t.Log("map为", mymap)
	t.Log("map为", &mymap2)
	t.Log(len(mymap2))
	delete(mymap2, "age")
	t.Log("map为", mymap2)

}
