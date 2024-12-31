package main

func main() {

}

func findCircleNum(isConnected [][]int) int {
	length := len(isConnected)
	components := 0
	visited := make([]bool, length)

	for i := 0; i < length; i++ {
		if !visited[i] {
			components++
			dfs(i, isConnected, visited)
		}
	}

	return components
}

func dfs(node int, isConnected [][]int, visited []bool) {
	visited[node] = true
	for i := 0; i < len(isConnected); i++ {
		if isConnected[node][i] == 1 && !visited[i] {
			dfs(i, isConnected, visited)
		}
	}
}
