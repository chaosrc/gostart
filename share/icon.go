package main

import (
	"image"
	"sync"
)

var icon map[string]image.Image

func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

var muOnce sync.Once

func Icon(name string) image.Image {
	muOnce.Do(loadIcons)
	return icon[name]
}

func loadIcon(name string) image.Image {
	return nil
}
