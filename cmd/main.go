package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// Initialize SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("Initializing SDL: %s", err)
	}
	defer sdl.Quit()

	// Create a window
	window, err := sdl.CreateWindow(
		"Go Pong",               // title
		sdl.WINDOWPOS_UNDEFINED, // x position
		sdl.WINDOWPOS_UNDEFINED, // y position
		800,                     // width
		600,                     // height
		sdl.WINDOW_SHOWN)        // flags
	if err != nil {
		log.Fatalf("Creating window: %s", err)
	}
	defer window.Destroy()

	// // Main loop
	// running := true
	//
	//	for running {
	//		// Handle events
	//		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	//			switch event.(type) {
	//			case *sdl.QuitEvent:
	//				running = false
	//			}
	//		}
	//
	//		// Game rendering and logic goes here
	//	}
}
