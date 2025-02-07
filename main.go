package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FindPathReq struct {
	Edges [][]int `json:"edges"`
	Start int     `json:"start"`
	End   int     `json:"end"`
}

func main() {

	r := gin.Default()

	r.POST("/find-paths", FindPaths)
	r.Run(":8080")
}

func FindPaths(c *gin.Context) {
	var req FindPathReq
	c.ShouldBindJSON(&req)
	m := len(req.Edges) // number of edges
	var n = -1          // total number of nodes
	for i := 0; i < m; i++ {
		for j := 0; j < len(req.Edges[i]); j++ {
			if n < req.Edges[i][j] {
				n = req.Edges[i][j]
			}
		}
	}

	var graph = make([][]int, n+1)

	for i := 0; i < m; i++ {
		graph[req.Edges[i][0]] = append(graph[req.Edges[i][0]], req.Edges[i][1])
	}

	// make path
	var ans [][]int
	var path []int
	dfs(graph, path, &ans, req.Start, req.End)
	c.JSON(http.StatusOK, ans)
}

func dfs(graph [][]int, path []int, ans *[][]int, start int, end int) {
	path = append(path, start)
	if start == end {
		copyPath := make([]int, len(path))
		copy(copyPath, path)
		*ans = append(*ans, copyPath)
		return
	}

	for i := 0; i < len(graph[start]); i++ {
		dfs(graph, path, ans, graph[start][i], end)
	}
}

//email id: amisha.mandiwal@streak.tech
