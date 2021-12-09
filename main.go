package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	screenWidth := int32(800)
	screenHeigth := int32(450)
	rl.InitWindow(screenWidth, screenHeigth, "flappy-game")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
