package main

import "fmt"

type UserProfile struct {
	id   int
	name string
}

func main() {
	resultFollowers := []UserProfile{
		{
			id:   2,
			name: "shabana",
		},
		{
			id:   3,
			name: "spalzes",
		},
	}
	resultFollowing := []UserProfile{
		{
			id:   2,
			name: "shabana",
		},
		{
			id:   5,
			name: "shivam",
		},
	}

	resultFollowsyou := []UserProfile{}
	m := make(map[int]UserProfile)
	for _, j := range resultFollowing {
		//fmt.Println(j.id)
		m[j.id] = j
	}
	for _, k := range resultFollowers {
		_, ok := m[k.id]
		//fmt.Println(ansuser)
		if !ok {
			resultFollowsyou = append(resultFollowsyou, k)
		}
	}
	// fmt.Println(resultFollowing)
	// fmt.Println(resultFollowers)

	fmt.Println(resultFollowsyou)
}
