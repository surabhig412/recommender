package main

type Recommender struct {
	entities []RecommenderEntity
}

type RecommenderEntity struct {
	name  string
	items map[string]RecommenderItem
}

type RecommenderItem struct {
	rating float64
	notes  string
}

func (r *Recommender) Add(re RecommenderEntity) {
	r.entities = append(r.entities, re)
}

func (r Recommender) Recommend(name string) map[string]float64 {
	var baseRe RecommenderEntity
	rNeighbor := make(map[string]float64, len(r.entities)-1)
	for _, re := range r.entities {
		if re.name == name {
			baseRe = re
			break
		}
	}
	for _, re := range r.entities {
		if re.name != name {
			rNeighbor[re.name] = pearsonCorrelation(baseRe.items, re.items)
		}
	}
	return rNeighbor
}
