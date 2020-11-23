package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const titleMenuSize = 100

func newTitleScreen(renderer *sdl.Renderer) *element {
	titleScreen := &element{}

	titleScreen.position = vector{
		x: 1020,
		y: 550,
	}

	titleScreen.active = true

	sr := newSpriteRenderer(titleScreen, renderer, "Sprites/StartMenu.bmp")
	titleScreen.addComponent(sr)

	ts := newTScreen(titleScreen)
	titleScreen.addComponent(ts)

	return titleScreen

}
