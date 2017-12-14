package main

import "fmt"

func main() {
	var r Recommender
	r1Items := make(map[string]RecommenderItem, 6)
	r2Items := make(map[string]RecommenderItem, 8)
	r3Items := make(map[string]RecommenderItem, 7)
	r1Items["a"] = RecommenderItem{2, "notes a"}
	r1Items["b"] = RecommenderItem{3, "notes b"}
	r1Items["c"] = RecommenderItem{5, "notes c"}
	r1Items["d"] = RecommenderItem{7, "notes d"}
	r1Items["e"] = RecommenderItem{11, "notes e"}
	r1Items["f"] = RecommenderItem{13, "notes f"}
	r.Add(RecommenderEntity{"r1", r1Items})
	r2Items["a"] = RecommenderItem{2, "notes a"}
	r2Items["b"] = RecommenderItem{4, "notes b"}
	r2Items["c"] = RecommenderItem{5, "notes c"}
	r2Items["d"] = RecommenderItem{8, "notes d"}
	r2Items["e"] = RecommenderItem{10, "notes e"}
	r2Items["f"] = RecommenderItem{14, "notes f"}
	r2Items["g"] = RecommenderItem{16, "notes g"}
	r2Items["h"] = RecommenderItem{18, "notes h"}
	r.Add(RecommenderEntity{"r2", r2Items})
	r3Items["a"] = RecommenderItem{1, "notes a"}
	r3Items["b"] = RecommenderItem{14, "notes b"}
	r3Items["c"] = RecommenderItem{50, "notes c"}
	r3Items["d"] = RecommenderItem{48, "notes d"}
	r3Items["e"] = RecommenderItem{100, "notes e"}
	r3Items["f"] = RecommenderItem{19, "notes f"}
	r3Items["g"] = RecommenderItem{7, "notes g"}
	r.Add(RecommenderEntity{"r3", r3Items})
	neighbors := r.Recommend("r1")
	for key, val := range neighbors {
		fmt.Printf("Pearson correlation coefficient for %s is %f \n", key, val)
	}
}
