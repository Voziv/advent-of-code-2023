package day10

import (
	"fmt"
	"github.com/voziv/advent-of-code-2023/internal/math"
)

type Maze struct {
	width  int
	height int
	data   []string
}

func NewMaze(data []string, width int, height int) *Maze {
	return &Maze{data: data, width: width, height: height}
}

var asciiToUnicodePipeMap = map[string]string{
	AsciiVerticalPipe:   VerticalPipe,
	AsciiHorizontalPipe: HorizontalPipe,
	AsciiNePipe:         NePipe,
	AsciiNwPipe:         NwPipe,
	AsciiSePipe:         SePipe,
	AsciiSwPipe:         SwPipe,
	AsciiGround:         Ground,
	AsciiAir:            Air,
	AsciiStart:          Start,
	AsciiOutOfBounds:    OutOfBounds,
}

const (
	AsciiVerticalPipe   = "|"
	AsciiHorizontalPipe = "-"
	AsciiNePipe         = "L"
	AsciiNwPipe         = "J"
	AsciiSePipe         = "F"
	AsciiSwPipe         = "7"
	AsciiGround         = "."
	AsciiAir            = "*"
	AsciiStart          = "S"
	AsciiOutOfBounds    = "!"
)

const (
	VerticalPipe   = "║"
	HorizontalPipe = "═"
	NePipe         = "╚"
	NwPipe         = "╝"
	SePipe         = "╔"
	SwPipe         = "╗"
	Ground         = "◉"
	Air            = " "
	Start          = "▩"
	OutOfBounds    = "!"
)

var Pipes = []string{
	VerticalPipe,
	HorizontalPipe,
	NePipe,
	NwPipe,
	SePipe,
	SwPipe,
	Start,
}

func (m *Maze) Start() *math.Vector2Int {
	pos := math.NewVector2Int(0, 0)
	for pos.X = 0; pos.X < m.width; pos.X++ {
		for pos.Y = 0; pos.Y < m.height; pos.Y++ {
			if m.data[pos.Y*m.width+pos.X] == Start {
				return pos
			}
		}
	}

	panic("No starting position found")
}

func (m *Maze) Tile(pos *math.Vector2Int) string {
	if pos.X < 0 || pos.X >= m.width || pos.Y < 0 || pos.Y >= m.height {
		return OutOfBounds
	}

	return m.data[pos.Y*m.width+pos.X]
}

func (m *Maze) SetTile(pos *math.Vector2Int, value string) {
	if pos.X < 0 || pos.X >= m.width || pos.Y < 0 || pos.Y >= m.height {
		panic("Cannot set tile outside of bounds")
	}

	m.data[pos.Y*m.width+pos.X] = value
}

func (m *Maze) PositionsConnect(pos1 *math.Vector2Int, pos2 *math.Vector2Int) bool {
	switch m.Tile(pos2) {
	case VerticalPipe:
		return pos1.X == pos2.X && (pos1.Y == pos2.Y-1 || pos1.Y == pos2.Y+1)
	case HorizontalPipe:
		return pos1.Y == pos2.Y && (pos1.X == pos2.X-1 || pos1.X == pos2.X+1)
	case NePipe: // L
		return (pos1.Y == pos2.Y && pos1.X == pos2.X+1) || (pos1.X == pos2.X && pos1.Y == pos2.Y-1)
	case NwPipe: // J
		return (pos1.Y == pos2.Y && pos1.X == pos2.X-1) || (pos1.X == pos2.X && pos1.Y == pos2.Y-1)
	case SePipe:
		return (pos1.Y == pos2.Y && pos1.X == pos2.X+1) || (pos1.X == pos2.X && pos1.Y == pos2.Y+1)
	case SwPipe:
		return (pos1.Y == pos2.Y && pos1.X == pos2.X-1) || (pos1.X == pos2.X && pos1.Y == pos2.Y+1)
	}

	return false
}

func (m *Maze) FindConnections(pos *math.Vector2Int) []*math.Vector2Int {
	var connections []*math.Vector2Int

	positions := []*math.Vector2Int{
		math.NewVector2Int(pos.X, pos.Y-1), // North
		math.NewVector2Int(pos.X+1, pos.Y), // East
		math.NewVector2Int(pos.X, pos.Y+1), // South
		math.NewVector2Int(pos.X-1, pos.Y), // West
	}

	for _, tile := range positions {
		//fmt.Printf("Checking if %+v (%s) connects to %+v (%s)", pos, m.Tile(pos), tile, m.Tile(tile))
		if m.PositionsConnect(pos, tile) {
			//fmt.Printf("Yep!\n")
			connections = append(connections, tile)
		} else {
			//fmt.Printf("Nope!\n")
		}
	}

	return connections
}

func (m *Maze) TravelPipe(cur *math.Vector2Int, pre *math.Vector2Int) (*math.Vector2Int, *math.Vector2Int) {
	switch m.Tile(cur) {
	case VerticalPipe: // |
		if pre.Y < cur.Y {
			return math.NewVector2Int(cur.X, cur.Y+1), cur
		}
		return math.NewVector2Int(cur.X, cur.Y-1), cur
	case HorizontalPipe: // -
		if pre.X < cur.X {
			return math.NewVector2Int(cur.X+1, cur.Y), cur
		}
		return math.NewVector2Int(cur.X-1, cur.Y), cur
	case NePipe: // L
		if pre.X > cur.X {
			return math.NewVector2Int(cur.X, cur.Y-1), cur
		}
		return math.NewVector2Int(cur.X+1, cur.Y), cur
	case NwPipe: // J
		if pre.X < cur.X {
			return math.NewVector2Int(cur.X, cur.Y-1), cur
		}
		return math.NewVector2Int(cur.X-1, cur.Y), cur
	case SePipe: // F
		if pre.X > cur.X {
			return math.NewVector2Int(cur.X, cur.Y+1), cur
		}
		return math.NewVector2Int(cur.X+1, cur.Y), cur
	case SwPipe: // 7
		if pre.X < cur.X {
			return math.NewVector2Int(cur.X, cur.Y+1), cur
		}
		return math.NewVector2Int(cur.X-1, cur.Y), cur
	case Start: // S
		panic("Cannot use start to travel")
	case Ground: // .
		panic("Cannot use ground to travel")
	}

	panic("No connections to travel to")
}

func (m *Maze) GetPipePositions() []*math.Vector2Int {
	var positions []*math.Vector2Int
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			for _, pipe := range Pipes {
				pos := math.NewVector2Int(x, y)
				if m.Tile(pos) == pipe {
					positions = append(positions, pos)
					break
				}
			}
		}
	}

	return positions
}

// Ew. lmao
func (m *Maze) CalculateStartShape(pos *math.Vector2Int) string {
	if m.PositionsConnect(pos, math.NewVector2Int(pos.X+1, pos.Y)) {
		// Has a right connection
		if m.PositionsConnect(pos, math.NewVector2Int(pos.X, pos.Y-1)) {
			// Has a top connection
			return NePipe
		} else if m.PositionsConnect(pos, math.NewVector2Int(pos.X, pos.Y+1)) {
			// Has a bottom connection
			return SePipe
		} else {
			return HorizontalPipe
		}
	} else if m.PositionsConnect(pos, math.NewVector2Int(pos.X-1, pos.Y)) {
		// has a left connection
		if m.PositionsConnect(pos, math.NewVector2Int(pos.X, pos.Y-1)) {
			// Has a top connection
			return NwPipe
		} else if m.PositionsConnect(pos, math.NewVector2Int(pos.X, pos.Y+1)) {
			// Has a bottom connection
			return SwPipe
		} else {
			return HorizontalPipe
		}
	} else {
		// is vertical
		return VerticalPipe
	}
}
func (m *Maze) FindNestsUsingScanLine(replaceOutsideWithAir bool) int {
	count := 0

	pos := math.NewVector2Int(0, 0)
	for pos.Y = 0; pos.Y < m.height; pos.Y++ {
		inside := false
		lastCorner := ""
		for pos.X = 0; pos.X < m.width; pos.X++ {
			tile := m.Tile(pos)
			if tile == Start {
				tile = m.CalculateStartShape(pos)
				fmt.Println("Start shape: " + tile)
			}

			switch tile {
			case Ground:
				if inside {
					count++
				} else {
					if replaceOutsideWithAir {
						m.SetTile(pos, Air)
					}
				}
				break
			case NePipe: // ╚
				lastCorner = tile
				break
			case NwPipe: // ╝
				if lastCorner != NePipe {
					inside = !inside
				}
				lastCorner = ""
				break
			case SePipe: // ╔
				lastCorner = tile
				break
			case SwPipe: // ╗
				if lastCorner != SePipe {
					inside = !inside
				}
				lastCorner = ""
				break
			case VerticalPipe:
				inside = !inside
				break
			case Air:
			case HorizontalPipe:
			case OutOfBounds:
				// Do nothing
				break
			case Start:
				panic("Please replace start tile")
			}
		}
	}
	return count
}

func (m *Maze) Print() {
	fmt.Println("MAP")
	pos := math.NewVector2Int(0, 0)
	for pos.Y = 0; pos.Y < m.height; pos.Y++ {
		for pos.X = 0; pos.X < m.width; pos.X++ {
			fmt.Printf(m.Tile(pos))
		}
		fmt.Println()
	}
	fmt.Println("END OF MAP")
}

func (m *Maze) DistanceToFurthestPointFromStart() int {

	newMap := make([]string, len(m.data))
	for i := 0; i < len(newMap); i++ {
		newMap[i] = Ground
	}
	newMaze := NewMaze(newMap, m.width, m.height)

	pos1 := m.Start()
	pos2 := pos1

	newMaze.SetTile(pos1, m.Tile(pos1))
	newMaze.SetTile(pos2, m.Tile(pos2))

	steps := 0

	// Run first step outside the loop so that
	// we start going in opposite directions
	steps++
	connections := m.FindConnections(pos1)

	prev1 := pos1
	pos1 = connections[0]
	prev2 := pos2
	pos2 = connections[1]

	newMaze.SetTile(pos1, m.Tile(pos1))
	newMaze.SetTile(pos2, m.Tile(pos2))

	for pos1.X != pos2.X || pos1.Y != pos2.Y {
		steps++

		pos1, prev1 = m.TravelPipe(pos1, prev1)
		pos2, prev2 = m.TravelPipe(pos2, prev2)

		newMaze.SetTile(pos1, m.Tile(pos1))
		newMaze.SetTile(pos2, m.Tile(pos2))
	}

	// This way our maze only contains data that includes
	// pipes that are connected to the starting position
	m.data = newMaze.data

	return steps
}
