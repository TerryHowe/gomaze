package maze

const VIEW = `
                                 
                                 
                                 
                                 
                                 
                                 
                                 
                                 
                                 
                                 
                                 
                                 
                                 
                                 
                                 
                                 
                                 
                                 
                                 
`
const LEFT = `
 \
  \
   +                             
   |                             
   |                             
   |                             
   |                             
   |                             
   |                             
   |                             
   |                             
   |                             
   |                             
   |                             
   |                             
   |                             
   +                             
  /                              
 /                               
`
const RIGHT = `
                                /
                               / 
                              +  
                              |  
                              |  
                              |  
                              |  
                              |  
                              |  
                              |  
                              |  
                              |  
                              |  
                              |  
                              |  
                              |  
                              +  
                               \\ 
                                \\
`
const FORWARD = `
                                 
                                 
   +--------------------------+  
   |                          |  
   |                          |  
   |                          |  
   |                          |  
   |                          |  
   |                          |  
   |                          |  
   |                          |  
   |                          |  
   |                          |  
   |                          |  
   |                          |  
   |                          |  
   +--------------------------+  
                                 
                                 
`
const LEFT_FORWARD_RIGHT = `
                                 
                                 
 --+                             
   |                             
   |                             
   |                             
   |                             
   |                             
   |                             
   |                             
   |                             
   |                             
   |                             
   |                             
   |                             
   |                             
 --+                             
                                 
                                 
`
const RIGHT_FORWARD_LEFT = `
                                 
                                 
                              +--
                              |  
                              |  
                              |  
                              |  
                              |  
                              |  
                              |  
                              |  
                              |  
                              |  
                              |  
                              |  
                              |  
                              +--
                                 
                                 
`
const FORWARD_FORWARD = `
                                 
                                 
                                 
                                 
                                 
      +--------------------+     
      |                    |     
      |                    |     
      |                    |     
      |                    |     
      |                    |     
      |                    |     
      |                    |     
      +--------------------+     
                                 
                                 
                                 
                                 
                                 
`
const FORWARD_LEFT = `
                                 
                                 
   +                             
   |\                            
   | \                           
   |  +                          
   |  |                          
   |  |                          
   |  |                          
   |  |                          
   |  |                          
   |  |                          
   |  |                          
   |  +                          
   | /                           
   |/                            
   +                             
                                 
                                 
`
const FORWARD_RIGHT = `
                                 
                                 
                              +  
                             /|  
                            / |  
                           +  |  
                           |  |  
                           |  |  
                           |  |  
                           |  |  
                           |  |  
                           |  |  
                           |  |  
                           +  |  
                            \ |  
                             \|  
                              +  
                                 
                                 
`
const FORWARD_LEFT_FORWARD_RIGHT = `
                                 
                                 
                                 
                                 
                                 
    --+                          
      |                          
      |                          
      |                          
      |                          
      |                          
      |                          
      |                          
    --+                          
                                 
                                 
                                 
                                 
                                 
`
const FORWARD_RIGHT_FORWARD_LEFT = `
                                 
                                 
                                 
                                 
                                 
                           +--   
                           |     
                           |     
                           |     
                           |     
                           |     
                           |     
                           |     
                           +--   
                                 
                                 
                                 
                                 
                                 
`


///*class TextView(object):
//
//	def render(self, room, direction):
//		t = [VIEW]
//		left_room = room.go_left(direction)
//		forward_room = room.go_forward(direction)
//		right_room = room.go_right(direction)
//		if left_room:
//			if not left_room.go_forward(direction):
//				t.append(LEFT_FORWARD_RIGHT)
//		else:
//			t.append(LEFT)
//		if forward_room:
//			forward_left_room = forward_room.go_left(direction)
//			if forward_left_room:
//				if not forward_left_room.go_forward(direction):
//					t.append(FORWARD_LEFT_FORWARD_RIGHT)
//			else:
//				t.append(FORWARD_LEFT)
//			if not forward_room.go_forward(direction):
//				t.append(FORWARD_FORWARD)
//			forward_right_room = forward_room.go_right(direction)
//			if forward_right_room:
//				if not forward_right_room.go_forward(direction):
//					t.append(FORWARD_RIGHT_FORWARD_LEFT)
//			else:
//				t.append(FORWARD_RIGHT)
//		else:
//			t.append(FORWARD)
//		if right_room:
//			if not right_room.go_forward(direction):
//				t.append(RIGHT_FORWARD_LEFT)
//		else:
//			t.append(RIGHT)
//		#print
//		#print "Direction: " + direction
//		#print ''.join([max(x) for x in zip(*t)])
//		#print
//		return ''.join([max(x) for x in zip(*t)])*/
