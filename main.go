package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 1200
	screenHeight = 900
)

type Circle struct {
	Position rl.Vector2
	Velocity rl.Vector2
	Size     float32
	Color    rl.Color
}

func update(Cluster *[]Circle) {
    for i, Circle := range *Cluster {
        Position := rl.Vector2Add(Circle.Position, Circle.Velocity)
        
        if Position.X < Circle.Size || Position.X > float32(screenWidth)-Circle.Size {
            Circle.Velocity.X *= -1
            Position.X = Circle.Position.X + Circle.Velocity.X 
        }
        if Position.Y < Circle.Size || Position.Y > float32(screenHeight)-Circle.Size {
            Circle.Velocity.Y *= -1
            Position.Y = Circle.Position.Y + Circle.Velocity.Y 
        }
        
        (*Cluster)[i].Position = Position 
        (*Cluster)[i].Velocity = Circle.Velocity 
    }
}
func draw(Cluster []Circle) {
	for _, Circle := range Cluster {
		rl.DrawCircleV(Circle.Position, Circle.Size, Circle.Color)
	}
}
func CreateCluster(count int) []Circle {
	Cluster := []Circle{}
	for i := 0; i < count; i++ {
		Cluster = append(Cluster, RandomCircle())
	}
	return Cluster
}

func main() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint | rl.FlagVsyncHint)
	rl.InitWindow(screenWidth, screenHeight, "Hello World - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	Cluster := CreateCluster(1000)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		update(&Cluster)
		draw(Cluster)
		rl.EndDrawing()
	}
}

func RandomPoint() rl.Vector2 {
	return rl.Vector2{
		X: float32(rl.GetRandomValue(-30, screenWidth-30)),
		Y: float32(rl.GetRandomValue(-30, screenHeight-30)),
	}
}
func RandomVelocity() rl.Vector2 {
	return rl.Vector2{
		X: float32(rl.GetRandomValue(-3, 3)),
		Y: float32(rl.GetRandomValue(-3, 3)),
	}
}
func RandomCircle() Circle {
	return Circle{
		Position: RandomPoint(),
		Velocity: RandomVelocity(),
		Size:     float32(rl.GetRandomValue(10, 30)),
		Color:    RandomColor(),
	}
}
func RandomColor() rl.Color {
	return rl.Color{
		R: uint8(rl.GetRandomValue(0, 255)),
		G: uint8(rl.GetRandomValue(0, 255)),
		B: uint8(rl.GetRandomValue(0, 255)),
		A: 120,
	}
}
