package main

import (
	"encoding/json"
	"fmt"
)

// Movie m
type Movie struct {
	Title  string
	Year   int  "json:\"release\""
	Color  bool `json:"color,omitempty"`
	Actors []string
	MyName string
	Myname string
}

func jsonMovie() {
	movies := []Movie{
		{
			Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
			MyName: "dingchao",
			Myname: "chao",
		},
		{
			Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"},
		},
		{
			Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen"},
		},
	}
	data, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", data)

	var movies2 []Movie

	err2 := json.Unmarshal(data, &movies2)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(movies2)

	fmt.Println("__________")

	type Title struct {
		Title  string
		Color  bool `json:"color,omitempty"`
		// MyName string
		Myname string
	}

	var titles []Title

	data2 := []byte(`[{"Title":"Casablanca","release":1942,"Actors":["Humphrey Bogart","Ingrid Bergman"],"MyName":"dingchao"},{"Title":"Cool Hand Luke","release":1967,"color":true,"Actors":["Paul Newman"],"MyName":"","Myname":""},{"Title":"Bullitt","release":1968,"color":true,"Actors":["Steve McQueen"],"MyName":"","Myname":""}]`)

	if err := json.Unmarshal(data2, &titles); err != nil {
		fmt.Println(err)
	}
	fmt.Println(titles)
}
