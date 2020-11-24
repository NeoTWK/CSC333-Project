package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type tScreen struct {
	container *element

	sr *spriteRenderer
}

func newTScreen(container *element) *tScreen {
	return &tScreen{
		container: container,

		sr: container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (mover *tScreen) onCollision(other *element) error {
	return nil
}

func (mover *tScreen) onUpdate() error {
	return nil
}

func (mover *tScreen) onDraw(renderer *sdl.Renderer) error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_S] == 1 {

		elements = append(elements, newPlayer(renderer))
		elements = append(elements, newBasicEnemy(renderer))

		mover.container.active = false

		initPlayerBulletPool(renderer)
		initEnemyBulletPool(renderer)

		for {
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch event.(type) {
				case *sdl.QuitEvent:
					return nil
				}
			}

			renderer.SetDrawColor(0, 255, 255, 255)

			renderer.Clear()

			for _, elem := range elements {
				if elem.active {
					err := elem.update()
					if err != nil {
						fmt.Println("updating element:", err)
						return nil
					}
					err = elem.draw(renderer)
					if err != nil {
						fmt.Println("drawing element:", elem)
						return nil
					}
				}
			}

			if err := checkCollisions(); err != nil {
				fmt.Println("checking collisions:", err)
				return nil
			}

			renderer.Present()
		}

	}
	return nil
}
