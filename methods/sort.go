package main

import (
	"sort"
	"time"
)

var _ sort.Interface = (StringList)(nil)

type StringList []string

func (s StringList) Len() int {
	return len(s)
}

func (s StringList) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s StringList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type byArtist []*Track

func (b byArtist) Len() int           { return len(b) }
func (b byArtist) Less(i, j int) bool { return b[i].Artist < b[j].Artist }
func (b byArtist) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

type customSort struct {
	tracks []*Track
	comp   func(a, b *Track) bool
}

func (c customSort) Len() int {
	return len(c.tracks)
}

func (c customSort) Less(i, j int) bool {
	return c.comp(c.tracks[i], c.tracks[j])
}

func (c customSort) Swap(i, j int) {
	c.tracks[i], c.tracks[j] = c.tracks[j], c.tracks[i]
}

