// EbitNew4 project EbitNew4.go
package EbitNew4

//"github.com/hajimehoshi/ebiten/v2"

var idCounter int = 0

func NewId() int {
	idCounter++
	return idCounter
}

func NewEnt[T any](ent Entity) T {
	Entitys = append(Entitys, ent)
	return ent.(T)
}

type Entity interface {
	Getid() int
}

var Entitys []any
