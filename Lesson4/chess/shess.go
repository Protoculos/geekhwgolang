//4 chess
package main

import "fmt"

type Hod struct {
	x int
	y int
}
type Steps struct {
	LeftUp    []int `json:"left_up,omitempty"`
	LeftDown  []int `json:"left_down,omitempty"`
	UpLeft    []int `json:"up_left,omitempty"`
	UpRight   []int `json:"up_right,omitempty"`
	RightUp   []int `json:"right_up,omitempty"`
	RightDown []int `json:"right_down,omitempty"`
	DownLeft  []int `json:"down_left,omitempty"`
	DownRight []int `json:"down_right,omitempty"`
}

func (h *Hod) horsStep() Steps {
	var s Steps
	//LeftUp
	if h.x-2 >= 1 && h.y+1 <= 8 {
		xx := h.x - 2
		yy := h.y + 1
		s.LeftUp = []int{xx, yy}
	}
	//LeftDown
	if h.x-2 >= 1 && h.y-1 >= 1 {
		xx := h.x - 2
		yy := h.y - 1
		s.LeftDown = []int{xx, yy}
	}
	//UpLeft
	if h.x-1 >= 1 && h.y+2 <= 8 {
		xx := h.x - 1
		yy := h.y + 2
		s.UpLeft = []int{xx, yy}
	}

	//UpRight
	if h.x+1 <= 8 && h.y+2 <= 8 {
		xx := h.x + 1
		yy := h.y + 2
		s.UpRight = []int{xx, yy}
	}

	//RightUp
	if h.x+2 <= 8 && h.y+1 <= 8 {
		xx := h.x + 2
		yy := h.y + 1
		s.RightUp = []int{xx, yy}
	}

	//RightDown
	if h.x+2 <= 8 && h.y-1 >= 1 {
		xx := h.x + 2
		yy := h.y - 1
		s.RightDown = []int{xx, yy}
	}

	//DownLeft
	if h.x-1 >= 1 && h.y-2 >= 1 {
		xx := h.x - 1
		yy := h.y - 2
		s.DownLeft = []int{xx, yy}
	}

	//DownRight
	if h.x+1 <= 8 && h.y-2 >= 1 {
		xx := h.x + 1
		yy := h.y - 2
		s.DownRight = []int{xx, yy}
	}

	return s

}

func main() {
	p := Hod{3, 5}
	st := p.horsStep()
	fmt.Println(st)
}
