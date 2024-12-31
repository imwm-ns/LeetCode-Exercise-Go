package main

func main() {
	// rooms := [][]int{{1}, {2}, {3}, {}}
	rooms := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}

	println(canVisitAllRooms(rooms))
}

func canVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]bool)

	retrieved := []int{0}

	for len(retrieved) > 0 {
		nextRooms := []int{}

		for _, room := range retrieved {
			visited[room] = true

			for _, key := range rooms[room] {
				if !visited[key] {
					nextRooms = append(nextRooms, key)
				}
			}
		}
		retrieved = nextRooms
	}

	return len(rooms) == len(visited)
}
