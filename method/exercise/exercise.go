package main

import "fmt"

type baseball struct{
	name string
	atBats int
	hits int
}

func(b baseball)calculate() float64{
	return float64(b.hits)/float64(b.atBats)
}

func main(){
	baseballPlayers := []baseball{
		{
			name:   "abc",
			atBats: 11,
			hits:   2,
		},
		{
			name:   "cba",
			atBats: 11,
			hits:   5,
		},
	}
	for _,v := range baseballPlayers{
		fmt.Printf("name %s average %f\n",v.name,v.calculate())
	}
}
