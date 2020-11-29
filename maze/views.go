package maze

//def index(request):
//return redirect('0/0/N', True)


func room(x int, y int, direction string) string {
	return ("Boom")
}

	//passages = Passage.objects.all()
	//theMaze = maze.Maze(passages)
	//room = theMaze.get_room(x, y)
	//if room is None:
	//	raise Http404("Room %d,%d does not exist" % (x, y))
	//	destinations = room.get_destinations(direction)
	//	maze_view = text_view.TextView().render(room, direction)
	//	context = {
	//		'x': x,
	//		'y': y,
	//		'direction': direction,
	//		'maze_view': maze_view,
	//		'direction_long': Passage.get_direction(direction),
	//	}
	//	context.update(destinations)
	//return render(request, 'maze/room.html', context)