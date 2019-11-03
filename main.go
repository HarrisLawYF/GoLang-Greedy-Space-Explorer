package main
import (
	"fmt"
	"direction"
	"robot"
	"ground"
	"strconv"
	"bufio"
	"os"
	"strings"
)

type edge = ground.Edge
type machine = robot.Robot

func Get_min(edge1 edge, edge2 edge, edge1_x int, edge1_y int, edge2_x int, edge2_y int)(int,int){
	x:=0
	y:=0
	if(edge1.Weight < edge2.Weight){
		x = edge1_x
		y = edge1_y
	} else{
		x = edge2_x
		y = edge2_y
	}
	return x,y
}

func Move_robot(bot machine, next_x int, next_y int)(machine){
	diff_x := next_x - bot.Get_x()
	diff_y := next_y - bot.Get_y()
	switch{
		case diff_x == 1 && diff_y == 0:
			if(bot.Get_direction() == 0){
				fmt.Println("Turn right, move forward")
				bot = bot.Turn_right()
				bot = bot.Move_forward()
			} else if(bot.Get_direction() == 1 || bot.Get_direction() == -3){
				fmt.Println("Move forward")
				bot = bot.Move_forward()
			} else if(bot.Get_direction() == 2 || bot.Get_direction() == -2){
				fmt.Println("Turn left, move forward")
				bot = bot.Turn_left()
				bot = bot.Move_forward()
			} else{
				fmt.Println("Turn right, turn right, move forward")
				bot = bot.Turn_right()
				bot = bot.Turn_right()
				bot = bot.Move_forward()
			}
		case diff_x == 0 && diff_y == 1:
			if(bot.Get_direction() == 0){
				fmt.Println("Move forward")
				bot = bot.Move_forward()
			} else if(bot.Get_direction() == 1 || bot.Get_direction() == -3){
				fmt.Println("Turn left, move forward")
				bot = bot.Turn_left()
				bot = bot.Move_forward()
			} else if(bot.Get_direction() == 2 || bot.Get_direction() == -2){
				fmt.Println("Turn left, turn left, move forward")
				bot = bot.Turn_left()
				bot = bot.Turn_left()
				bot = bot.Move_forward()
			} else{
				fmt.Println("Turn right, move forward")
				bot = bot.Turn_right()
				bot = bot.Move_forward()
			}
		case diff_x == 0 && diff_y == 0:
			fmt.Println("Do nothing")
		default:
			fmt.Println("Default case")
    }
	fmt.Println("Robot moving to (x,y): ",next_x,next_y)
	fmt.Println()
	return bot
}

func Find_shortest_path(edges [][]edge, current_x int, current_y int, map_bound_size int, bot machine){
	if(current_x == map_bound_size && current_y == map_bound_size){
		fmt.Println("Destination reached: ",current_x,current_y)
	} else if(current_x == map_bound_size){
		next_x := current_x
		next_y := current_y+1
		bot = Move_robot(bot, next_x, next_y)
		Find_shortest_path(edges,next_x,next_y,map_bound_size,bot)
	} else if(current_y == map_bound_size){
		next_x := current_x+1
		next_y := current_y
		bot = Move_robot(bot, next_x, next_y)
		Find_shortest_path(edges,next_x,next_y,map_bound_size,bot)
	} else{
		next_x, next_y := Get_min(edges[current_x+1][current_y], edges[current_x][current_y+1], current_x+1, current_y, current_x, current_y+1)
		bot = Move_robot(bot, next_x, next_y)
		Find_shortest_path(edges,next_x,next_y,map_bound_size,bot)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Map size: ")
    map_size_str, _ := reader.ReadString('\n')
    map_size, err := strconv.Atoi(strings.Replace(map_size_str, "\r\n", "", -1))
	if(err!=nil){
		fmt.Print("Map size has wrong value: ",err)
	}
	
	fmt.Print("Robot X coordinate: ")
    x_str, _ := reader.ReadString('\n')
    x, err := strconv.Atoi(strings.Replace(x_str, "\r\n", "", -1))
	if(err!=nil){
		fmt.Print("Coordinate X has wrong value: ",err)
	}
	if(x < 0 || x > map_size){
		x = 0
	}
	
	fmt.Print("Robot Y coordinate: ")
    y_str, _ := reader.ReadString('\n')
    y, err := strconv.Atoi(strings.Replace(y_str, "\r\n", "", -1))
	if(err!=nil){
		fmt.Print("Coordinate Y has wrong value: ",err)
	}
	if(y < 0 || y > map_size){
		x = 0
	}

	ground := ground.Create(map_size)
	
	navigator := robot.Create(x,y,direction.Enum.E)
	
	fmt.Println("Map generated x0[y0,y1,y2..],x1[y0,y1,y2]... :")
	fmt.Println(ground)
	fmt.Println()
	fmt.Println("Navigator X102 current direction: ",navigator.Get_direction_str())
	fmt.Println("Navigator X102 current coordinates (x,y): ",navigator.Get_x(),navigator.Get_y())
	fmt.Println("--------------------------")
	Find_shortest_path(ground, navigator.Get_x(), navigator.Get_y(), map_size, navigator)
}