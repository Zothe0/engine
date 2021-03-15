package game

import (
	mgl "github.com/go-gl/mathgl/mgl32"
)

type Game struct {
	gameState      state
	width          int32
	height         int32
	cameraPosition mgl.Vec3
	view           mgl.Mat4
}
type state int

const (
	Run state = iota
	Stop
)

func InitGame() *Game {
	g := new(Game)
	g.gameState = Stop
	return g
}
