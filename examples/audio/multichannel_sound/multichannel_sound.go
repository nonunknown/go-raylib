package main

import (
	"fmt"
	rl "goray/raylib"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [audio] example - Multichannel sound playing")
	rl.InitAudioDevice()

	fxWav := rl.LoadSound("../audio/resources/sound.wav")
	fxOgg := rl.LoadSound("../audio/resources/target.ogg")

	rl.SetSoundVolume(fxWav, 0.2)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(int32(rl.KEY_ENTER)) {
			rl.PlaySoundMulti(fxWav)
		}
		if rl.IsKeyPressed(int32(rl.KEY_SPACE)) {
			rl.PlaySoundMulti(fxOgg)
		}

		rl.BeginDrawing()

		rl.ClearBackground(*rl.RayWhite)

		rl.DrawText("MULTICHANNEL SOUND PLAYING", 20, 20, 20, *rl.Gray)
		rl.DrawText("Press SPACE to play new ogg instance!", 200, 120, 20, *rl.LightGray)
		rl.DrawText("Press ENTER to play new wav instance!", 200, 180, 20, *rl.LightGray)

		rl.DrawText(fmt.Sprintf("CONCURRENT SOUNDS PLAYING: %d", rl.GetSoundsPlaying()), 220, 280, 20, *rl.Red)

		rl.EndDrawing()
	}

	rl.StopSoundMulti()

	rl.UnloadSound(fxWav)
	rl.UnloadSound(fxOgg)

	rl.CloseAudioDevice()

	rl.CloseWindow()
}
