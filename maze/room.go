package maze

type Room struct {
	//COMPASS [...]string = {'N','E','S','W'}
	X int
	Y int
	//Passages make(map[string]Room)
}

func (this *Room) GoForward() bool {
	return true
}

	//def go_forward(self, direction):
	//	return(self.passages.get(direction, None))
	//
	//def get_left_direction(self, direction):
	//	idx = (self.COMPASS.index(direction) - 1) % len(self.COMPASS)
	//	return(self.COMPASS[idx])
	//
	//def go_left(self, direction):
	//	return(self.passages.get(self.get_left_direction(direction), None))
	//
	//def get_right_direction(self, direction):
	//	idx = (self.COMPASS.index(direction) + 1) % len(self.COMPASS)
	//	return(self.COMPASS[idx])
	//
	//def go_right(self, direction):
	//	return(self.passages.get(self.get_right_direction(direction), None))
	//
	//def get_backward_direction(self, direction):
	//	idx = (self.COMPASS.index(direction) + 2) % len(self.COMPASS)
	//	return(self.COMPASS[idx])
	//
	//def go_backward(self, direction):
	//	return(self.passages.get(self.get_backward_direction(direction), None))
	//
	//def get_destinations(self, direction):
	//	left_direction = self.get_left_direction(direction)
	//	right_direction = self.get_right_direction(direction)
	//	forward = self.go_forward(direction)
	//	backward = self.go_backward(direction)
	//	if forward:
	//		forward = ("%d/%d/%s" % (forward.room_x, forward.room_y, direction))
	//	if backward:
	//		backward = ("%d/%d/%s" % (backward.room_x, backward.room_y, direction))
	//	return {
	//		'left': ("%d/%d/%s" % (self.room_x, self.room_y, left_direction)),
	//		'forward': forward,
	//		'right': ("%d/%d/%s" % (self.room_x, self.room_y,  right_direction)),
	//		'backward': backward,
	//	}
